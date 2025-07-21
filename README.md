# OTP Service Project

This repository contains the implementation of an OTP (One-Time Password) service project following Clean Architecture
principles. The project is Dockerized and provides RESTful APIs for sending and verifying OTP codes.

---

## Table of Contents

- [Project Overview](#project-overview)
- [Why Redis for OTP Storage?](#why-redis-for-otp-storage)
- [User Storage with MySQL](#user-storage-with-mysql)
- [Rate Limiting](#rate-limiting)
- [Project Structure](#project-structure)
- [How to Run](#how-to-run)
- [API Documentation (Swagger)](#api-documentation-swagger)

---

## Project Overview

This service is designed to manage OTP codes for user authentication. It follows the Clean Architecture pattern as
introduced by Uncle Bob, ensuring separation of concerns and maintainability.

The project is fully containerized with Docker, allowing easy setup and deployment.

---

## Why Redis for OTP Storage?

Redis is used for storing OTP codes because:

- **Speed:** Redis is an in-memory data store, which offers extremely fast read/write operations, critical for
  time-sensitive OTP validation.
- **Expiration:** Redis provides built-in key expiration, making it easy to automatically invalidate OTP codes after a
  short lifetime.
- **Concurrency:** It handles high concurrency well, suitable for multiple OTP requests simultaneously.
- **Scalability:** Redis can scale easily and can be used as a distributed cache.

---

## User Storage with MySQL

User information is stored in MySQL. This choice is somewhat arbitrary and not critical to the OTP functionality. For
this project, using an in-memory store could have sufficed, but MySQL was chosen to illustrate persistence with a
relational database.

---

## Rate Limiting

**The project requirements regarding rate limiting were not explicitly defined.**

Based on experience with similar projects, rate limiting has been implemented at the API Gateway layer to control
request frequency.

Alternatively, rate limiting could be implemented using Redis based on the user's mobile number, which would provide
fine-grained control directly within the service.

---

## Project Structure

The project follows **Clean Architecture** principles (Uncle Bob's approach), dividing the codebase into clear layers:

```
+---------------------+
|    Delivery Layer    |  <-- HTTP / REST API handlers
+---------------------+
|     Service Layer    |  <-- Business logic / application rules
+---------------------+
|    Repository Layer  |  <-- Data access (MySQL, Redis)
+---------------------+
|       Domain         |  <-- Core entities and interfaces
+---------------------+
```

Additionally, this project uses the **go-standard-layout** for organizing Go code, which is a widely adopted and
idiomatic project layout in the Go community.

This structure helps maintain separation of concerns and allows easy testing and scalability.

---

## How to Run

To get the project up and running, follow these steps:

1. **Start the application (with Docker Compose):**
   ```bash
   make development-up
   ```

2. **Run database migrations:**
   ```bash
   make migrate-up
   ```

These commands will spin up all required services (app, MySQL, Redis, etc.) and prepare the database schema.

For more commands and options, please refer to the `Makefile` in the root of the project.

---

## API Documentation (Swagger)

After the services are running, you can access the Swagger UI documentation for the API at:

```
http://localhost/v1/swagger/index.html
```

Additionally, to generate or update Swagger docs, run the following command:

```bash
swag init --generalInfo cmd/main.go --parseDependency
```

This interface provides details on all available endpoints, request/response formats, and allows you to test the API
interactively.

---

Thank you for reviewing the project!  