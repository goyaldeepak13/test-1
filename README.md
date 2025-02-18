# Customer-Product Service

This service provides APIs to create customers and products, as well as fetch them by their IDs.

## API Endpoints

### 1. Create Customer
**Endpoint**: `POST /customers`  
**Request and Response**:
![Create Customer Request](https://github.com/user-attachments/assets/96582d77-d8df-4bda-ba83-7a0171b4e16e)  
 
---

### 2. Fetch Customer
**Endpoint**: `GET /customers/:id`  
**Request and Response**:
![Fetch Customer Request](https://github.com/user-attachments/assets/1d5fc62b-d322-4bfb-96cf-1d8a40fe62b7)  
 
---

### 3. Create Product
**Endpoint**: `POST /products`  
**Request and Response**:
![Create Product Request](https://github.com/user-attachments/assets/f445ddb1-121d-417a-a330-b70c9d1cf4e8)  
 

---

### 4. Fetch Product
**Endpoint**: `GET /products/:id`  
**Request and Response**:
![Fetch Product Request](https://github.com/user-attachments/assets/d09d0214-38b0-456e-a2c0-7e69bf497ce9)  
 

---

# Order Management Service

This service allows the creation of orders after validating customers and products, and retrieving orders by ID.

## Features

- **Create a New Order**:
  - Validates the customer using the `/customers` endpoint.
  - Validates the product using the `/products` endpoint.
  - Stores the order in the database.
- **Retrieve Order by ID**:
  - Fetch order details by providing the order ID.

---

## API Endpoints

### 1. Create Order
**Endpoint**: `POST /orders`  
**Request**:  
Validate customer and product, then create an order with the selected product and total bill.  
**Response**:  
Returns the created order.

---

### 2. Fetch Order
**Endpoint**: `GET /orders/:id`  
**Request**:  
Fetch an order by its ID.  
**Response**:  
Returns the order details if found, or an error message if not.

---
 

 




 
