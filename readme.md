# Golang Bank API

A simple banking API built with Golang that supports account creation, transactions, and balance inquiries. The project is containerized using Docker.

## Features
- User account creation
- Deposit and withdrawal functionality
- Transfer funds between accounts
- View account balances
- Secure authentication and authorization

## Technologies Used
- Golang
- Gorilla Mux (Router)
- PostgreSQL (Database)
- GORM (ORM)
- JWT Authentication
- Docker (Containerization)

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.19.5 or later)
- [Docker](https://www.docker.com/get-started) (version 24.0.5)
- PostgreSQL (if running locally)

## Cloning the Repository

```sh
git clone https://github.com/d3kanesa/Golang-Bank-Api.git
cd Golang-Bank-Api
```

## Running within Go

```sh
go mod init github.com/d3kanesa/Golang-Bank-Api
go mod tidy
go run cmd/api/main.go
```

## Running with Docker

```sh
docker build -t golang-bank-api .
docker run -p 8000:8000 golang-bank-api
```

## API Endpoints

| Method | Endpoint               | Description                 |
|--------|------------------------|-----------------------------|
| GET    | `/account/coins`       | Get account balance        |
| GET    | `/account/transactions`| View transaction history   |
| POST   | `/account/deposit`     | Deposit funds              |
| POST   | `/account/withdraw`    | Withdraw funds             |
| POST   | `/account/transfer`    | Transfer funds to another account |

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a Pull Request

## License
This project is licensed under the MIT License.

