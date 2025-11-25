package domain

import (
	"time"

	"github.com/google/uuid"
)

// PaymentStatus represents the status of a payment transaction
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusPaid      PaymentStatus = "PAID"
	PaymentStatusExpired   PaymentStatus = "EXPIRED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
)

// PaymentTransaction represents a payment transaction in the database
type PaymentTransaction struct {
	ID               uuid.UUID      `db:"id" json:"id"`
	OrderID          uuid.UUID      `db:"order_id" json:"order_id"`
	UserID           uuid.UUID      `db:"user_id" json:"user_id"`
	XenditInvoiceID  *string        `db:"xendit_invoice_id" json:"xendit_invoice_id,omitempty"`
	XenditExternalID string         `db:"xendit_external_id" json:"xendit_external_id"`
	PaymentMethod    *string        `db:"payment_method" json:"payment_method,omitempty"`
	PaymentChannel   *string        `db:"payment_channel" json:"payment_channel,omitempty"`
	PaymentStatus    PaymentStatus  `db:"payment_status" json:"payment_status"`
	Amount           float64        `db:"amount" json:"amount"`
	PaidAmount       float64        `db:"paid_amount" json:"paid_amount"`
	AdminFee         float64        `db:"admin_fee" json:"admin_fee"`
	TotalAmount      float64        `db:"total_amount" json:"total_amount"`
	InvoiceURL       *string        `db:"invoice_url" json:"invoice_url,omitempty"`
	CheckoutURL      *string        `db:"checkout_url" json:"checkout_url,omitempty"`
	PaidAt           *time.Time     `db:"paid_at" json:"paid_at,omitempty"`
	ExpiredAt        *time.Time     `db:"expired_at" json:"expired_at,omitempty"`
	Description      *string        `db:"description" json:"description,omitempty"`
	PaymentDetails   interface{}    `db:"payment_details" json:"payment_details,omitempty"`
	XenditResponse   interface{}    `db:"xendit_response" json:"xendit_response,omitempty"`
	CreatedAt        time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt        *time.Time     `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt        *time.Time     `db:"deleted_at" json:"deleted_at,omitempty"`
	Order            *Order         `db:"-" json:"order,omitempty"`
}

// CreatePaymentRequest represents the request to create a payment
type CreatePaymentRequest struct {
	OrderID          uuid.UUID `json:"order_id" validate:"required"`
	PaymentMethod    string    `json:"payment_method" validate:"required"`
	SuccessRedirect  string    `json:"success_redirect_url,omitempty"`
	FailureRedirect  string    `json:"failure_redirect_url,omitempty"`
}

// PaymentResponse represents the payment response
type PaymentResponse struct {
	ID               uuid.UUID     `json:"id"`
	OrderID          uuid.UUID     `json:"order_id"`
	XenditExternalID string        `json:"xendit_external_id"`
	PaymentStatus    PaymentStatus `json:"payment_status"`
	Amount           float64       `json:"amount"`
	TotalAmount      float64       `json:"total_amount"`
	InvoiceURL       *string       `json:"invoice_url,omitempty"`
	CheckoutURL      *string       `json:"checkout_url,omitempty"`
	ExpiredAt        *time.Time    `json:"expired_at,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
}

// XenditInvoiceRequest represents the request to Xendit API to create invoice
type XenditInvoiceRequest struct {
	ExternalID      string                   `json:"external_id"`
	Amount          float64                  `json:"amount"`
	PayerEmail      string                   `json:"payer_email,omitempty"`
	Description     string                   `json:"description"`
	InvoiceDuration int                      `json:"invoice_duration"` // in seconds
	Currency        string                   `json:"currency"`
	Items           []XenditInvoiceItem      `json:"items,omitempty"`
	Customer        XenditCustomer           `json:"customer,omitempty"`
	SuccessRedirect string                   `json:"success_redirect_url,omitempty"`
	FailureRedirect string                   `json:"failure_redirect_url,omitempty"`
	PaymentMethods  []string                 `json:"payment_methods,omitempty"`
}

// XenditInvoiceItem represents an item in the Xendit invoice
type XenditInvoiceItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
}

// XenditCustomer represents customer information for Xendit
type XenditCustomer struct {
	GivenNames   string `json:"given_names,omitempty"`
	Email        string `json:"email,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
}

// XenditInvoiceResponse represents the response from Xendit API
type XenditInvoiceResponse struct {
	ID              string                 `json:"id"`
	ExternalID      string                 `json:"external_id"`
	UserID          string                 `json:"user_id"`
	Status          string                 `json:"status"`
	MerchantName    string                 `json:"merchant_name"`
	Amount          float64                `json:"amount"`
	PayerEmail      string                 `json:"payer_email"`
	Description     string                 `json:"description"`
	ExpiryDate      time.Time              `json:"expiry_date"`
	InvoiceURL      string                 `json:"invoice_url"`
	AvailableBanks  []XenditAvailableBank  `json:"available_banks"`
	AvailableRetail []XenditAvailableRetail `json:"available_retail_outlets"`
	AvailableEwallet []XenditAvailableEwallet `json:"available_ewallets"`
	ShouldExclude   []string               `json:"should_exclude_credit_card"`
	ShouldSendEmail bool                   `json:"should_send_email"`
	Created         time.Time              `json:"created"`
	Updated         time.Time              `json:"updated"`
	Currency        string                 `json:"currency"`
}

// XenditAvailableBank represents available bank for payment
type XenditAvailableBank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int    `json:"identity_amount"`
}

// XenditAvailableRetail represents available retail outlet
type XenditAvailableRetail struct {
	RetailOutletName string `json:"retail_outlet_name"`
}

// XenditAvailableEwallet represents available e-wallet
type XenditAvailableEwallet struct {
	EwalletType string `json:"ewallet_type"`
}

// XenditCallbackPayload represents the callback payload from Xendit
type XenditCallbackPayload struct {
	ID              string    `json:"id"`
	ExternalID      string    `json:"external_id"`
	UserID          string    `json:"user_id"`
	Status          string    `json:"status"`
	MerchantName    string    `json:"merchant_name"`
	Amount          float64   `json:"amount"`
	PaidAmount      float64   `json:"paid_amount"`
	BankCode        string    `json:"bank_code"`
	PaidAt          time.Time `json:"paid_at"`
	PayerEmail      string    `json:"payer_email"`
	Description     string    `json:"description"`
	AdjustedReceived float64   `json:"adjusted_received_amount"`
	FeesReceived    float64   `json:"fees_paid_amount"`
	Updated         time.Time `json:"updated"`
	Created         time.Time `json:"created"`
	Currency        string    `json:"currency"`
	PaymentChannel  string    `json:"payment_channel"`
	PaymentMethod   string    `json:"payment_method"`
}

// PaymentListResponse represents the response for listing payments
type PaymentListResponse struct {
	Payments   []PaymentTransaction `json:"payments"`
	Total      int                  `json:"total"`
	Page       int                  `json:"page"`
	PerPage    int                  `json:"per_page"`
	TotalPages int                  `json:"total_pages"`
}
