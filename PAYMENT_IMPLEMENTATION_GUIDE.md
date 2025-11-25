# Payment Integration with Xendit - Implementation Guide

**Date**: November 24, 2024
**Payment Gateway**: Xendit (Indonesia)
**Version**: 1.0

## Overview

This guide covers the complete implementation of payment functionality integrated with Xendit payment gateway for the E-Commerce Backend API.

## 📋 Table of Contents

1. [Features Implemented](#features-implemented)
2. [Database Schema](#database-schema)
3. [Setup Instructions](#setup-instructions)
4. [API Endpoints](#api-endpoints)
5. [Xendit Integration](#xendit-integration)
6. [Testing Guide](#testing-guide)
7. [Webhook Configuration](#webhook-configuration)
8. [Payment Flow](#payment-flow)
9. [Security Considerations](#security-considerations)

## ✅ Features Implemented

### Core Features
- ✅ Create payment for orders via Xendit API
- ✅ Multiple payment methods support (Bank Transfer, E-Wallet, Retail)
- ✅ Real-time payment status tracking
- ✅ Automatic order status update on successful payment
- ✅ Payment cancellation for pending payments
- ✅ Xendit webhook callback handling
- ✅ Payment history with pagination
- ✅ Admin payment management

### Payment Methods Supported
- Bank Transfer (All major Indonesian banks)
- E-Wallets (OVO, Dana, LinkAja, GoPay)
- Retail Outlets (Alfamart, Indomaret)
- Credit Cards
- QRIS

## 🗄️ Database Schema

### Payment Transactions Table

```sql
CREATE TABLE payment_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Xendit transaction details
    xendit_invoice_id VARCHAR(255) UNIQUE,
    xendit_external_id VARCHAR(255) UNIQUE NOT NULL,
    
    -- Payment information
    payment_method VARCHAR(50),
    payment_channel VARCHAR(100),
    payment_status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    
    -- Amount details
    amount DECIMAL(15,2) NOT NULL,
    paid_amount DECIMAL(15,2) DEFAULT 0,
    admin_fee DECIMAL(15,2) DEFAULT 0,
    total_amount DECIMAL(15,2) NOT NULL,
    
    -- Payment URLs
    invoice_url TEXT,
    checkout_url TEXT,
    
    -- Payment dates
    paid_at TIMESTAMP,
    expired_at TIMESTAMP,
    
    -- Additional information
    description TEXT,
    payment_details JSONB,
    xendit_response JSONB,
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Payment Status Values
- `PENDING` - Payment created, waiting for customer to pay
- `PAID` - Payment successfully completed
- `EXPIRED` - Payment expired (default 24 hours)
- `FAILED` - Payment failed
- `CANCELLED` - Payment cancelled by user

## 🛠️ Setup Instructions

### 1. Install Dependencies

```bash
# Install golang-migrate tool
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Or using brew (macOS)
brew install golang-migrate

# Or using curl (Linux)
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

### 2. Configure Environment Variables

Add to your `.env` file:

```bash
# Xendit Payment Configuration
XENDIT_API_KEY=your-xendit-secret-api-key-here
XENDIT_CALLBACK_TOKEN=your-xendit-callback-verification-token-here
```

**To get Xendit API Key:**
1. Register at https://dashboard.xendit.co
2. Navigate to Settings > API Keys
3. Copy your **Secret Key** (starts with `xnd_`)
4. Copy your **Callback Verification Token**

### 3. Run Database Migration

```bash
# Run migration to create payment_transactions table
make migrate-up

# Or manually:
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce_db?sslmode=disable" up

# To rollback (if needed):
make migrate-down
```

### 4. Verify Installation

```bash
# Check if migration was successful
psql -d ecommerce_db -c "\d payment_transactions"

# Start the server
make run

# Check Swagger documentation
open http://localhost:8080/swagger/index.html
```

## 🔗 API Endpoints

### User Payment Endpoints

#### 1. Create Payment
**POST** `/api/v1/payments`

Create a payment transaction for an order.

**Request:**
```json
{
  "order_id": "uuid",
  "payment_method": "BANK_TRANSFER",
  "success_redirect_url": "https://yourapp.com/payment/success",
  "failure_redirect_url": "https://yourapp.com/payment/failed"
}
```

**Response (201):**
```json
{
  "message": "Payment created successfully",
  "payment": {
    "id": "uuid",
    "order_id": "uuid",
    "xendit_external_id": "ORDER-uuid-timestamp",
    "payment_status": "PENDING",
    "amount": 100000,
    "total_amount": 100000,
    "invoice_url": "https://checkout.xendit.co/web/xxxx",
    "expired_at": "2024-11-25T10:00:00Z",
    "created_at": "2024-11-24T10:00:00Z"
  }
}
```

#### 2. Get User Payments
**GET** `/api/v1/payments?page=1&per_page=10`

Retrieve user's payment history with pagination.

#### 3. Get Payment by ID
**GET** `/api/v1/payments/:id`

Get detailed information about a specific payment.

#### 4. Get Payment by Order ID
**GET** `/api/v1/payments/order?order_id=uuid`

Get payment information for a specific order.

#### 5. Check Payment Status
**GET** `/api/v1/payments/:id/status`

Check the current payment status from Xendit.

#### 6. Cancel Payment
**POST** `/api/v1/payments/:id/cancel`

Cancel a pending payment transaction.

### Admin Endpoints

#### 7. Get All Payments (Admin)
**GET** `/api/v1/admin/payments?page=1&per_page=10`

Get all payments with pagination (admin only).

### Webhook Endpoint

#### 8. Xendit Callback
**POST** `/api/v1/payments/xendit/callback`

Webhook endpoint for Xendit payment notifications (public, no auth required).

## 🔄 Xendit Integration

### Invoice Creation Flow

1. User creates an order
2. User initiates payment for the order
3. Backend creates Xendit invoice with:
   - Order details
   - Customer information
   - Payment methods
   - Invoice duration (24 hours)
4. Xendit returns invoice URL
5. Customer is redirected to Xendit checkout page
6. Customer completes payment
7. Xendit sends webhook callback
8. Backend updates payment and order status

### Supported Payment Methods

```go
// Example payment methods you can specify:
payment_methods := []string{
    "BANK_TRANSFER",    // All major banks
    "CREDIT_CARD",      // Visa, Mastercard, JCB, Amex
    "OVO",              // OVO e-wallet
    "DANA",             // DANA e-wallet
    "LINKAJA",          // LinkAja e-wallet
    "GOPAY",            // GoPay e-wallet
    "SHOPEEPAY",        // ShopeePay e-wallet
    "ALFAMART",         // Alfamart retail
    "INDOMARET",        // Indomaret retail
    "QRIS",             // QRIS standard
}
```

### Invoice Details Structure

```go
XenditInvoiceRequest{
    ExternalID:      "ORDER-uuid-timestamp",
    Amount:          100000,
    PayerEmail:      "customer@example.com",
    Description:     "Payment for Order #12345",
    InvoiceDuration: 86400,  // 24 hours in seconds
    Currency:        "IDR",
    Items: []XenditInvoiceItem{
        {
            Name:     "Product Name",
            Quantity: 2,
            Price:    40000,
            Category: "Product",
        },
        {
            Name:     "Shipping Cost",
            Quantity: 1,
            Price:    20000,
            Category: "Shipping",
        },
    },
    Customer: XenditCustomer{
        GivenNames:   "John Doe",
        Email:        "customer@example.com",
        MobileNumber: "+628123456789",
    },
    PaymentMethods:  []string{"BANK_TRANSFER"},
    SuccessRedirect: "https://yourapp.com/payment/success",
    FailureRedirect: "https://yourapp.com/payment/failed",
}
```

## 🧪 Testing Guide

### 1. Testing with Postman

Import the updated Postman collection: `E-Commerce_Complete_API.postman_collection.json`

**Test Flow:**
1. Register/Login to get JWT token
2. Create a product
3. Add product to cart
4. Create an order
5. Create payment for the order
6. Copy `invoice_url` from response
7. Open invoice URL in browser
8. Use Xendit test cards/accounts

### 2. Xendit Test Mode

In test mode, use these test credentials:

**Bank Transfer:**
- Use any bank code
- No actual money transfer needed
- Status changes automatically in dashboard

**E-Wallets (OVO):**
- Phone: `+6281212345678`
- OTP: `123456`

**Credit Card:**
- Card Number: `4000000000000002` (Success)
- Card Number: `4000000000000010` (Declined)
- CVV: `123`
- Expiry: Any future date

### 3. Testing Webhook Locally

Use ngrok to expose your local server:

```bash
# Install ngrok
brew install ngrok

# Start your server
make run

# In another terminal, expose port 8080
ngrok http 8080

# Use the ngrok URL in Xendit webhook settings
# Example: https://abc123.ngrok.io/api/v1/payments/xendit/callback
```

### 4. Manual Webhook Testing

Use Postman to simulate webhook callback:

**POST** `/api/v1/payments/xendit/callback`

```json
{
  "id": "xendit-invoice-id",
  "external_id": "ORDER-uuid-timestamp",
  "user_id": "xendit-user-id",
  "status": "PAID",
  "merchant_name": "Your Store",
  "amount": 100000,
  "paid_amount": 100000,
  "bank_code": "BCA",
  "paid_at": "2024-11-24T10:00:00.000Z",
  "payer_email": "customer@example.com",
  "description": "Payment for Order",
  "adjusted_received_amount": 100000,
  "fees_paid_amount": 0,
  "updated": "2024-11-24T10:00:00.000Z",
  "created": "2024-11-24T09:00:00.000Z",
  "currency": "IDR",
  "payment_channel": "BCA",
  "payment_method": "BANK_TRANSFER"
}
```

## 🪝 Webhook Configuration

### Setting up Xendit Webhook

1. Login to Xendit Dashboard
2. Go to Settings > Callbacks/Webhooks
3. Add new webhook URL:
   - **Invoice Paid**: `https://yourdomain.com/api/v1/payments/xendit/callback`
   - **Invoice Expired**: Same URL
4. Save the callback verification token
5. Add token to `.env` as `XENDIT_CALLBACK_TOKEN`

### Webhook Security

The webhook endpoint should verify:
1. Callback token from Xendit header
2. Request signature (if enabled)
3. External ID matches your system

**TODO**: Add callback token verification in handler:

```go
// In payment_handler.go XenditCallback function
callbackToken := c.Get("x-callback-token")
if callbackToken != cfg.Xendit.CallbackToken {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Invalid callback token",
    })
}
```

## 💳 Payment Flow Diagram

```
┌─────────────┐
│   Customer  │
└──────┬──────┘
       │ 1. Create Order
       ▼
┌─────────────────┐
│   Your Backend  │
└────────┬────────┘
         │ 2. Create Payment
         │    (POST /api/v1/payments)
         ▼
    ┌────────────┐
    │   Xendit   │
    │   Create   │
    │  Invoice   │
    └─────┬──────┘
          │ 3. Return invoice_url
          ▼
    ┌──────────────┐
    │   Customer   │
    │   Redirected │
    │   to Xendit  │
    │   Checkout   │
    └──────┬───────┘
           │ 4. Complete Payment
           ▼
      ┌──────────┐
      │  Xendit  │
      │  Sends   │
      │ Callback │
      └────┬─────┘
           │ 5. Webhook notification
           ▼
    ┌─────────────────┐
    │   Your Backend  │
    │  Update Payment │
    │   & Order Status│
    └─────────────────┘
```

## 🔒 Security Considerations

### 1. API Key Security
- ✅ Store Xendit API key in environment variables
- ✅ Never commit `.env` file to version control
- ✅ Use different keys for development and production
- ✅ Rotate keys periodically

### 2. Webhook Security
- ⚠️ **TODO**: Implement callback token verification
- ⚠️ **TODO**: Implement signature verification
- ✅ Validate external_id matches your system
- ✅ Check payment status before updating order

### 3. Payment Validation
- ✅ Verify order belongs to user
- ✅ Prevent duplicate payments for same order
- ✅ Check order amount matches payment amount
- ✅ Validate payment status transitions

### 4. Error Handling
- ✅ Handle Xendit API errors gracefully
- ✅ Log all payment transactions
- ✅ Implement retry mechanism for failed callbacks
- ✅ Alert on suspicious payment activities

## 📊 Monitoring & Logging

### What to Monitor

1. **Payment Success Rate**
   - Track successful vs failed payments
   - Monitor by payment method

2. **Payment Processing Time**
   - Time from creation to completion
   - Webhook callback delay

3. **Failed Payments**
   - Reason for failures
   - Customer drop-off rate

4. **Webhook Health**
   - Callback success rate
   - Response time

### Logging Best Practices

```go
// Log payment creation
log.Printf("Payment created: payment_id=%s, order_id=%s, amount=%.2f", 
    payment.ID, payment.OrderID, payment.Amount)

// Log webhook received
log.Printf("Webhook received: external_id=%s, status=%s, amount=%.2f",
    payload.ExternalID, payload.Status, payload.Amount)

// Log errors
log.Printf("Payment creation failed: order_id=%s, error=%v",
    orderID, err)
```

## 🚀 Production Checklist

- [ ] Use production Xendit API key
- [ ] Configure production webhook URL (HTTPS required)
- [ ] Implement callback token verification
- [ ] Implement signature verification
- [ ] Set up monitoring and alerts
- [ ] Test all payment methods
- [ ] Configure retry mechanism
- [ ] Set up payment reconciliation
- [ ] Document payment dispute handling
- [ ] Train support team on payment issues

## 📖 Additional Resources

- [Xendit API Documentation](https://developers.xendit.co/api-reference/)
- [Xendit Invoice API](https://developers.xendit.co/api-reference/#create-invoice)
- [Xendit Webhook Guide](https://developers.xendit.co/api-reference/#invoice-callback)
- [Xendit Test Credentials](https://developers.xendit.co/api-reference/#test-scenarios)
- [Payment Best Practices](https://developers.xendit.co/guides/best-practices)

## 🐛 Troubleshooting

### Payment not created
- Check Xendit API key is correct
- Verify order exists and belongs to user
- Check order doesn't already have pending payment
- Review server logs for Xendit API errors

### Webhook not received
- Verify webhook URL is accessible (use ngrok for local testing)
- Check Xendit webhook configuration
- Review server logs for webhook endpoint errors
- Test webhook manually using Postman

### Payment status not updating
- Check webhook is being received
- Verify external_id matches your payment record
- Check order status update logic
- Review database transaction logs

## 📝 Notes

- Default invoice expiration: 24 hours
- Minimum payment amount: Rp 10,000
- Maximum payment amount: Rp 100,000,000
- Xendit fees: 0-4% depending on payment method
- Settlement time: T+1 to T+7 days depending on method

## 👥 Support

For issues or questions:
- Xendit Support: support@xendit.co
- Xendit Community: https://community.xendit.co
- Backend Team: backend-team@yourcompany.com

---

**Last Updated**: November 24, 2024
**Version**: 1.0
**Author**: E-Commerce Backend Team
