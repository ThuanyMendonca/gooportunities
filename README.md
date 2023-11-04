# Go Oportunities

- Introduction to Golang and building modern APIs
- Development environment setup for creating the API
- Using Go-Gin as a router for route management
- Implementing SQLite as the database for the API
- Using GoORM for communication with the database
- Integrating Swagger for API documentation and testing
- Creating a modern package structure for organizing the project
- Implementing a complete job opportunities API with endpoints for searching, creating, editing, and deleting opportunities
- Implementing automated tests to ensure API quality

## About
This repository is based on Arthur Andrade's [video tutorial](https://youtu.be/wyEYpX5U4Vg) on golang

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Installation
Clone the repository: git clone https://github.com/username/repo-name.git
Install the dependencies: go mod download
Build the application: go build
Run the application: ./main

## Make
The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

make run: Run the application without generating API documentation.
make run-with-docs: Generate the API documentation using Swag, then run the application.
make build: Build the application and create an executable file named gopportunities.
make test: Run tests for all packages in the project.
make docs: Generate the API documentation using Swag.
make clean: Remove the gopportunities executable and delete the ./docs directory.

