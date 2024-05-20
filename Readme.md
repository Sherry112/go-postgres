# Stock Management API

This repository contains a sample project for a Stock Management API using Go and PostgreSQL. The API allows you to perform CRUD operations on stock data.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Running Tests](#running-tests)
- [License](#license)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/stock-management-api.git
   cd stock-management-api

2. **Install dependencies:**

   ```bash
    go mod tidy

3. **Set up your environment variables:**

    Create a .env file in the root directory of your project and add the following:
   ```bash
    POSTGRES_URL=your_postgres_connection_string

4. **Run the application:**

   ```bash
    go build

5. **Start the server:**

   ```bash
    go run main.go

6. **Access the API using a tool like Postman or curl.**



## API Endpoints

### Get All Stocks

- **URL:** `/stocks`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** `[{ "STOCKID": 1, "NAME": "Stock1", "PRICE": 100, "COMPANY": "Company1" }, ...]`

### Get Stock by ID

- **URL:** `/stocks/{id}`
- **Method:** `GET`
- **URL Params:**
  - `id=[integer]`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{ "STOCKID": 1, "NAME": "Stock1", "PRICE": 100, "COMPANY": "Company1" }`

### Create Stock

- **URL:** `/stocks`
- **Method:** `POST`
- **Data Params:**
  - `{ "NAME": "Stock1", "PRICE": 100, "COMPANY": "Company1" }`
- **Success Response:**
  - **Code:** 201
  - **Content:** `{ "id": 1, "message": "stock created successfully" }`

### Update Stock

- **URL:** `/stocks/{id}`
- **Method:** `PUT`
- **URL Params:**
  - `id=[integer]`
- **Data Params:**
  - `{ "NAME": "Stock1", "PRICE": 100, "COMPANY": "Company1" }`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{ "id": 1, "message": "Stock updated successfully. Total rows/records affected 1" }`

### Delete Stock

- **URL:** `/stocks/{id}`
- **Method:** `DELETE`
- **URL Params:**
  - `id=[integer]`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{ "id": 1, "message": "Stock deleted successfully. Total rows/records affected 1" }`

## Environment Variables

- `POSTGRES_URL`: The connection string for your PostgreSQL database.


## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher.
- [GoDotEnv](https://github.com/joho/godotenv) - Loads environment variables from `.env`.
- [PostgreSQL Driver](https://github.com/lib/pq) - PostgreSQL database driver for Go.
