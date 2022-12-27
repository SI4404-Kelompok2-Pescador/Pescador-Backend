# Pescador-Backend

Pescador-Backend is a backend service for Pescador, a web application for managing and sharing fishing spots.
Using DDD (Domain Driven Design) and Clean Architecture, this project aims to provide a clean and maintainable codebase.
It is written in Go and uses PostgreSQL as the database.
This project is part of the final project for the Web Application Development course at the Telkom University.

## Project Structure


```
. root
├───cmd
├───config
├───domain
│   └───entity
└───internal
    ├───controllers
    │   ├───admin
    │   ├───product
    │   ├───store
    │   └───user
    ├───dto
    ├───middleware
    └───routes

```

## API Documentation
API documentation can be found [here](https://documenter.getpostman.com/view/16260600/2s8Z6x1smZ)

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
$ go mod tidy
```

- Create a `.env` file in the root directory of the project and copy the contents of `.env.example` into it then change with your own configuration. 
```bash
$ cp .env.example .env
```


### Run the application

```bash
$ docker-compose up -d
$ go run cmd/main.go
```

Running the application will create a database named `pescador` in your local PostgreSQL instance.

## Database
database url: `localhost:8080`

## Contributing
TL;DR: Please read the [Contributing Guide](CONTRIBUTING.md) before contributing.

## Regards
- Faruqi Hafiz 