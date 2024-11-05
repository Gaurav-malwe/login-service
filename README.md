# Golang Login Service

This project is a simple login service built with Go that integrates with AWS Cognito for user authentication. It uses MongoDB for storing user-related data. The development environment is simplified using Docker and LocalStack to simulate AWS services.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Docker Setup](#docker-setup)
- [LocalStack Initialization](#localstack-initialization)

## Features

- User sign-up and sign-in using AWS Cognito.
- User data storage with MongoDB.
- Local development using Docker and LocalStack.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.21.6 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/Gaurav-malwe/login-service
    cd login-service
    ```

2. Build and run the service:

    ```bash
    docker-compose up
    ```

    This will start the LocalStack service and your login service.

## API Endpoints

This document provides details about the API endpoints available in the Golang Login Service.

### Sign Up

- **Endpoint**: `POST /signup`
- **Description**: This endpoint allows a new user to sign up for the service.
- **Request Body**:

  ```json
  {
    "email": "gsmalwe02@gmail.com",
    "password": "Password@123",
    "username": "gsmalwe02@gmail.com",
    "mobile": "8983878968",
    "admin": true
  }
  ```

- **Response**:
  - **Success**:
    - Status Code: `200`
    - Body:

    ```json
      {
        "message": "User created successfully"
      }
    ```

- **Endpoint**: `POST /confirm`
- **Description**: This endpoint allows a new user to confirm user registration.
- **Request Body**:

  ```json
  { 
    "username": "gaurav",
    "confirmation_code": "123456"
  }
  ```

- **Response**:
  - **Success**:
    - Status Code: `200`
    - Body:

    ```json
      {
        "message": "User created successfully"
      }
    ```

### Sign In

- **Endpoint**: `POST /signin`
- **Description**: This endpoint allows a user to sign in to the service.
- **Request Body**:

    ```json
    {
        "username": "exampleuser",
        "password": "examplepassword"
    }
    ```

- **Response**:
  - **Success**:
    - Status Code: `200 OK`
    - Body:

    ```json
    {
    "message": "User signed in successfully",
    "token": "your_token"
    }
    ```

  - **Error**:
    - Status Code: `401 Unauthorized`
    - Body:

    ```json
    {
    "error": "Invalid username or password"
    }
    ```

## Notes

- Ensure to replace `exampleuser` and `examplepassword` with actual values when making requests.
- The token returned on successful sign-in can be used for authenticating subsequent requests.

This API can be tested using tools like Postman or cURL. Please refer to the main README for setup instructions and running the service.

## Docker Setup

This project includes a Dockerfile and docker-compose.yml to facilitate building and running the application in a containerized environment.

### Running the Service

1. Make sure Docker is running on your machine.
2. Use the following command to start the application:

    ```bash
    docker-compose up
    ```

## LocalStack Initialization

The project uses LocalStack to simulate AWS services. The initialization script located in the localstack directory sets up the necessary resources (e.g., Cognito User Pool).
Use the CLIENT_ID and USER_POOL_ID generated from the script while initializing service.
