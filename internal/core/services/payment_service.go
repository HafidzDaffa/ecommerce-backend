package services

import (
	"bytes"
	"encoding/json"
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepo ports.PaymentRepository
	orderRepo   ports.OrderRepository
	userRepo    ports.UserRepository
	xenditAPIKey string
	xenditBaseURL string
}

func NewPaymentService(
	paymentRepo ports.PaymentRepository,
	orderRepo ports.OrderRepository,
	userRepo ports.UserRepository,
	xenditAPIKey string,
) *PaymentService {
	return &PaymentService{
		paymentRepo:   paymentRepo,
		orderRepo:     orderRepo,
		userRepo:      userRepo,
		xenditAPIKey:  xenditAPIKey,
		xenditBaseURL: "https://api.xendit.co",
	}
}

func (s *PaymentService) CreatePayment(userID uuid.UUID, req *domain.CreatePaymentRequest) (*domain.PaymentResponse, error) {
	// Get order details
	order, err := s.orderRepo.GetByID(req.OrderID)
	if err != nil {
		return nil, errors.New("order not found")
	}
	
	// Verify order belongs to user
	if order.UserID != userID {
		return nil, errors.New("unauthorized: order does not belong to user")
	}

	// Check if order already has payment
	existingPayment, _ := s.paymentRepo.GetByOrderID(order.ID)
	if existingPayment != nil && existingPayment.PaymentStatus == domain.PaymentStatusPending {
		return &domain.PaymentResponse{
			ID:               existingPayment.ID,
			OrderID:          existingPayment.OrderID,
			XenditExternalID: existingPayment.XenditExternalID,
			PaymentStatus:    existingPayment.PaymentStatus,
			Amount:           existingPayment.Amount,
			TotalAmount:      existingPayment.TotalAmount,
			InvoiceURL:       existingPayment.InvoiceURL,
			CheckoutURL:      existingPayment.CheckoutURL,
			ExpiredAt:        existingPayment.ExpiredAt,
			CreatedAt:        existingPayment.CreatedAt,
		}, nil
	}

	// Get user details
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate external ID
	externalID := fmt.Sprintf("ORDER-%s-%d", order.ID.String(), time.Now().Unix())

	// Prepare Xendit invoice request
	xenditReq := domain.XenditInvoiceRequest{
		ExternalID:      externalID,
		Amount:          order.TotalAmount,
		PayerEmail:      user.Email,
		Description:     fmt.Sprintf("Payment for Order #%s", order.ID.String()[:8]),
		InvoiceDuration: 86400, // 24 hours
		Currency:        "IDR",
		SuccessRedirect: req.SuccessRedirect,
		FailureRedirect: req.FailureRedirect,
	}

	// Filter payment methods based on request
	if req.PaymentMethod != "" {
		xenditReq.PaymentMethods = []string{req.PaymentMethod}
	}

	// Prepare items for invoice
	if order.Items != nil && len(order.Items) > 0 {
		xenditReq.Items = make([]domain.XenditInvoiceItem, 0)
		for _, item := range order.Items {
			xenditReq.Items = append(xenditReq.Items, domain.XenditInvoiceItem{
				Name:     item.Product.ProductName,
				Quantity: item.Quantity,
				Price:    item.Price,
				Category: "Product",
			})
		}
		// Add shipping as an item
		if order.ShippingCost > 0 {
			xenditReq.Items = append(xenditReq.Items, domain.XenditInvoiceItem{
				Name:     "Shipping Cost",
				Quantity: 1,
				Price:    order.ShippingCost,
				Category: "Shipping",
			})
		}
	}

	// Set customer details
	fullName := ""
	if user.FullName != nil {
		fullName = *user.FullName
	}
	phoneNumber := ""
	if user.PhoneNumber != nil {
		phoneNumber = *user.PhoneNumber
	}
	xenditReq.Customer = domain.XenditCustomer{
		GivenNames:   fullName,
		Email:        user.Email,
		MobileNumber: phoneNumber,
	}

	// Create invoice in Xendit
	xenditResp, err := s.createXenditInvoice(&xenditReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create xendit invoice: %w", err)
	}

	// Create payment transaction in database
	payment := &domain.PaymentTransaction{
		ID:               uuid.New(),
		OrderID:          order.ID,
		UserID:           userID,
		XenditInvoiceID:  &xenditResp.ID,
		XenditExternalID: externalID,
		PaymentStatus:    domain.PaymentStatusPending,
		Amount:           order.TotalAmount,
		PaidAmount:       0,
		AdminFee:         0,
		TotalAmount:      order.TotalAmount,
		InvoiceURL:       &xenditResp.InvoiceURL,
		ExpiredAt:        &xenditResp.ExpiryDate,
		Description:      &xenditReq.Description,
	}

	// Marshal xendit response to JSON
	xenditRespJSON, _ := json.Marshal(xenditResp)
	payment.XenditResponse = xenditRespJSON

	err = s.paymentRepo.Create(payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment transaction: %w", err)
	}

	return &domain.PaymentResponse{
		ID:               payment.ID,
		OrderID:          payment.OrderID,
		XenditExternalID: payment.XenditExternalID,
		PaymentStatus:    payment.PaymentStatus,
		Amount:           payment.Amount,
		TotalAmount:      payment.TotalAmount,
		InvoiceURL:       payment.InvoiceURL,
		CheckoutURL:      payment.CheckoutURL,
		ExpiredAt:        payment.ExpiredAt,
		CreatedAt:        payment.CreatedAt,
	}, nil
}

func (s *PaymentService) createXenditInvoice(req *domain.XenditInvoiceRequest) (*domain.XenditInvoiceResponse, error) {
	url := fmt.Sprintf("%s/v2/invoices", s.xenditBaseURL)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(s.xenditAPIKey, "")
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("xendit API error: %s - %s", resp.Status, string(body))
	}

	var xenditResp domain.XenditInvoiceResponse
	err = json.Unmarshal(body, &xenditResp)
	if err != nil {
		return nil, err
	}

	return &xenditResp, nil
}

func (s *PaymentService) GetPaymentByID(userID uuid.UUID, paymentID uuid.UUID) (*domain.PaymentTransaction, error) {
	payment, err := s.paymentRepo.GetByID(paymentID)
	if err != nil {
		return nil, err
	}

	if payment.UserID != userID {
		return nil, errors.New("unauthorized access to payment")
	}

	return payment, nil
}

func (s *PaymentService) GetPaymentByOrderID(userID uuid.UUID, orderID uuid.UUID) (*domain.PaymentTransaction, error) {
	// Verify order belongs to user
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}
	
	// Verify order belongs to user
	if order.UserID != userID {
		return nil, errors.New("unauthorized: order does not belong to user")
	}

	payment, err := s.paymentRepo.GetByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) GetUserPayments(userID uuid.UUID, page, perPage int) (*domain.PaymentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	payments, total, err := s.paymentRepo.GetUserPayments(userID, page, perPage)
	if err != nil {
		return nil, err
	}

	totalPages := (total + perPage - 1) / perPage

	return &domain.PaymentListResponse{
		Payments:   payments,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *PaymentService) GetAllPayments(page, perPage int) (*domain.PaymentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	payments, total, err := s.paymentRepo.GetAllPayments(page, perPage)
	if err != nil {
		return nil, err
	}

	totalPages := (total + perPage - 1) / perPage

	return &domain.PaymentListResponse{
		Payments:   payments,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *PaymentService) HandleXenditCallback(payload *domain.XenditCallbackPayload) error {
	// Get payment by external ID
	payment, err := s.paymentRepo.GetByExternalID(payload.ExternalID)
	if err != nil {
		return fmt.Errorf("payment not found for external_id: %s", payload.ExternalID)
	}

	// Update payment status based on Xendit status
	var status domain.PaymentStatus
	switch payload.Status {
	case "PAID", "SETTLED":
		status = domain.PaymentStatusPaid
	case "EXPIRED":
		status = domain.PaymentStatusExpired
	case "FAILED":
		status = domain.PaymentStatusFailed
	default:
		status = domain.PaymentStatusPending
	}

	// Update payment
	payment.PaymentStatus = status
	payment.PaidAmount = payload.PaidAmount
	if payload.Status == "PAID" || payload.Status == "SETTLED" {
		payment.PaidAt = &payload.PaidAt
	}
	payment.PaymentChannel = &payload.PaymentChannel
	payment.PaymentMethod = &payload.PaymentMethod

	// Marshal callback payload to JSON
	callbackJSON, _ := json.Marshal(payload)
	payment.XenditResponse = callbackJSON

	err = s.paymentRepo.Update(payment)
	if err != nil {
		return err
	}

	// Update order status if payment is successful
	if status == domain.PaymentStatusPaid {
		// Update order status to "Paid" (assuming status_id 2 is for paid orders)
		statusID := 2 // You might want to make this configurable
		err = s.orderRepo.UpdateStatus(payment.OrderID, statusID, nil)
		if err != nil {
			return fmt.Errorf("failed to update order status: %w", err)
		}
	}

	return nil
}

func (s *PaymentService) CheckPaymentStatus(userID uuid.UUID, paymentID uuid.UUID) (*domain.PaymentTransaction, error) {
	payment, err := s.GetPaymentByID(userID, paymentID)
	if err != nil {
		return nil, err
	}

	// If payment is still pending, check with Xendit
	if payment.PaymentStatus == domain.PaymentStatusPending && payment.XenditInvoiceID != nil {
		xenditInvoice, err := s.getXenditInvoice(*payment.XenditInvoiceID)
		if err != nil {
			return payment, nil // Return existing payment if Xendit API fails
		}

		// Update payment status based on Xendit response
		var status domain.PaymentStatus
		switch xenditInvoice.Status {
		case "PAID", "SETTLED":
			status = domain.PaymentStatusPaid
		case "EXPIRED":
			status = domain.PaymentStatusExpired
		default:
			status = domain.PaymentStatusPending
		}

		if payment.PaymentStatus != status {
			payment.PaymentStatus = status
			if status == domain.PaymentStatusPaid {
				payment.PaidAmount = xenditInvoice.Amount
				paidAt := time.Now()
				payment.PaidAt = &paidAt
			}
			_ = s.paymentRepo.Update(payment)
		}
	}

	return payment, nil
}

func (s *PaymentService) getXenditInvoice(invoiceID string) (*domain.XenditInvoiceResponse, error) {
	url := fmt.Sprintf("%s/v2/invoices/%s", s.xenditBaseURL, invoiceID)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(s.xenditAPIKey, "")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("xendit API error: %s - %s", resp.Status, string(body))
	}

	var xenditResp domain.XenditInvoiceResponse
	err = json.Unmarshal(body, &xenditResp)
	if err != nil {
		return nil, err
	}

	return &xenditResp, nil
}

func (s *PaymentService) CancelPayment(userID uuid.UUID, paymentID uuid.UUID) error {
	payment, err := s.GetPaymentByID(userID, paymentID)
	if err != nil {
		return err
	}

	// Can only cancel pending payments
	if payment.PaymentStatus != domain.PaymentStatusPending {
		return errors.New("can only cancel pending payments")
	}

	// Update payment status to cancelled
	payment.PaymentStatus = domain.PaymentStatusCancelled
	err = s.paymentRepo.Update(payment)
	if err != nil {
		return err
	}

	// Optionally expire the invoice in Xendit
	// Note: Xendit doesn't have a cancel endpoint, invoices will expire automatically

	return nil
}
