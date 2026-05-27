# API Contract v1

Base URL: /v1
Auth: Bearer access token (except login/refresh).

## Auth

### POST /auth/login
Request:
{
  "email": "admin@example.com",
  "password": "string"
}

Response 200:
{
  "accessToken": "jwt",
  "refreshToken": "jwt",
  "user": {
    "id": "uuid",
    "email": "admin@example.com",
    "fullName": "Admin",
    "role": "admin"
  }
}

### POST /auth/refresh
Request:
{
  "refreshToken": "jwt"
}

Response 200:
{
  "accessToken": "jwt"
}

## Customers

### GET /customers
Query params:
- page, pageSize
- search

Response 200:
{
  "items": [{"id":"uuid","code":"C001","name":"ABC","phone":"..."}],
  "total": 1
}

### POST /customers
Request:
{
  "code": "C001",
  "name": "Customer Name",
  "phone": "090...",
  "email": "customer@example.com",
  "address": "Address",
  "note": "Optional"
}

Response 201: customer object

## Products

### GET /products
Query params:
- page, pageSize
- search
- isActive

Response 200:
{
  "items": [{"id":"uuid","sku":"SKU001","name":"Product A","salePrice":100000}],
  "total": 1
}

### POST /products
Request:
{
  "sku": "SKU001",
  "name": "Product A",
  "unit": "item",
  "costPrice": 50000,
  "salePrice": 100000,
  "isActive": true
}

Response 201: product object

### PUT /products/:id
Request: same fields as create (partial update allowed)

Response 200: product object

## Orders

### GET /orders
Query params:
- page, pageSize
- fromDate, toDate
- status
- search (order_no or customer name)

Response 200:
{
  "items": [
    {
      "id": "uuid",
      "orderNo": "SO-20260527-001",
      "orderDate": "2026-05-27",
      "status": "confirmed",
      "customer": {"id":"uuid","name":"ABC"},
      "totalAmount": 1200000,
      "paidAmount": 800000
    }
  ],
  "total": 1
}

### GET /orders/:id
Response 200:
{
  "id": "uuid",
  "orderNo": "SO-20260527-001",
  "customerId": "uuid",
  "orderDate": "2026-05-27",
  "status": "draft",
  "subtotal": 1000000,
  "discountAmount": 0,
  "taxAmount": 100000,
  "totalAmount": 1100000,
  "paidAmount": 0,
  "note": "Optional",
  "items": [
    {
      "productId": "uuid",
      "quantity": 2,
      "unitPrice": 500000,
      "discountAmount": 0,
      "lineTotal": 1000000
    }
  ],
  "payments": []
}

### POST /orders
Request:
{
  "customerId": "uuid",
  "orderDate": "2026-05-27",
  "note": "Optional",
  "items": [
    {
      "productId": "uuid",
      "quantity": 2,
      "unitPrice": 500000,
      "discountAmount": 0
    }
  ]
}

Response 201:
{
  "id": "uuid",
  "orderNo": "SO-20260527-001",
  "status": "draft"
}

### PUT /orders/:id
Request: same structure as POST /orders, partial update allowed for draft only.

Response 200: updated order object

### POST /orders/:id/confirm
Response 200:
{
  "id": "uuid",
  "status": "confirmed"
}

### POST /orders/:id/cancel
Response 200:
{
  "id": "uuid",
  "status": "canceled"
}

## Payments

### POST /orders/:id/payments
Request:
{
  "method": "bank",
  "amount": 300000,
  "paidAt": "2026-05-27T10:00:00Z",
  "referenceNo": "BANK123",
  "note": "Optional"
}

Response 201: payment object

## Reports

### GET /reports/revenue
Query params:
- fromDate, toDate

Response 200:
{
  "totalOrders": 30,
  "grossRevenue": 50000000,
  "discountAmount": 1000000,
  "netRevenue": 49000000,
  "paidAmount": 43000000,
  "outstandingAmount": 6000000
}

### GET /reports/top-products
Query params:
- fromDate, toDate
- limit

Response 200:
{
  "items": [
    {
      "productId": "uuid",
      "sku": "SKU001",
      "name": "Product A",
      "totalQty": 100,
      "totalRevenue": 25000000
    }
  ]
}

## Error Format
All endpoints should return a consistent error shape:
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input",
    "details": []
  }
}
