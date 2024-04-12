# Microservice Project README

This repository contains a microservice project designed to execute Go code within a Docker environment.

## Components

- **Frontend Service**: Responsible for providing the user interface.
- **Gateway Service**: Acts as a broker for incoming requests and forwards them to the appropriate microservices.
- **Authentication Service**: Handles user authentication and authorization.
- **Problem Service**: Executes Go code within a sandboxed environment.
- **GoExec Service**: Executes Go code submitted by users.

## Prerequisites

- [Docker](https://docs.docker.com/compose/install/): Docker is required to run the application. If you don't have Docker installed, you can follow the instructions [here](https://docs.docker.com/compose/install/) to install Docker Compose, which is included with Docker Desktop for Windows and macOS.

## Installation

1. Copy the `docker-compose.yml` in this repository to your local machine.
2. Run `docker-compose up --build -d`.
3. Navigate to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your web browser to access the Swagger UI.





