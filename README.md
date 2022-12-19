# Pescador-Backend

Pescador-Backend is a backend service for Pescador, a web application for managing and sharing fishing spots.

## Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Installation

- Clone the repository

```bash
$ git clone https://github.com/SI4404-Kelompok2-Pescador/Pescador-Backend.git
```

- Install dependencies

```bash
$ go mod download
```

- Create a `.env` file in the root directory of the project and copy the contents of `.env.example` into it

- Run the application

```bash
$ docker-compose up
$ go run main.go
```

Running the application will create a database named `pescador` in your local PostgreSQL instance.

## Contributing
TL;DR: Please read the [Contributing Guide](CONTRIBUTING.md) before contributing.

## Regards
- Faruqi Hafiz 