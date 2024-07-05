# Japhy Backend Test in Golang

## Technical Stack
- Go
- Docker
- MySQL

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Tasks
The frontend team needs to have a CRUD api to manage the breeds of pets. The breeds are stored in a CSV file located at `./breeds.csv`. 
the aim of this test is to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources
you are free to take initiatives and make improvements to the codebase.
Have fun and good luck!

### you need to implement the following tasks:
- database migrations using the `database_actions` directory.
- CRUD functionality for the breeds resource regarding the `breeds.csv` file.
- search functionality to filter breeds based on pet characteristics (weight and species).


## Installation

1. Clone the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`
5. Once the application is up and running, you can access the REST API at http://localhost:50010. Use tools like Postman or curl to interact with the API.
6. `curl -v http://localhost:50010/health` to ensure your application is running.
