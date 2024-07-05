# Backend Test en Go

## Description

This project is a backend test in Go. It aims to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Installation

1. Clone the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`

## Tasks
- Implement database migrations using the `database_actions` directory.
- Implement the CRUD functionality for the breeds resource regarding the `breeds.csv` file.
- Implement the search functionality to filter breeds based on pet characteristics (weight and species).

