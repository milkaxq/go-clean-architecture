# Test

[![Go Version](https://img.shields.io/badge/go-1.22-blue.svg)](https://golang.org/dl/)

## Overview

**Test** is a Go application built using Clean Architecture principles. The goal of this architecture is to keep the business logic at the core of the application, independent of any external frameworks or technologies. This makes the application more maintainable, testable, and flexible for future changes.

## Project Structure

The project is organized into several layers, each with its own responsibility. Below is an overview of the main directories:

```plaintext
project/
│
├── cmd/
│   └── app/
│       └── main.go                // Entry point of the application
│
├── internal/
│   ├── domain/                    // Domain Layer: Core business logic and entities
│   │   ├── entities/              // Domain Entities (e.g., User)
│   │   ├── repositories/          // Repository Interfaces for persistence
│   │   └── services/              // Domain Services containing business rules
│   │
│   ├── application/               // Application Layer: Use cases and application logic
│   │   ├── usecases/              // Application Use Cases (e.g., CreateUser)
│   │   ├── interfaces/            // Application Interfaces (e.g., Logger)
│   │   ├── dtos/                  // Data Transfer Objects (DTOs)
│   │   └── responses/             // Response models returned by the application
│   │
│   ├── infrastructure/            // Infrastructure Layer: Frameworks and external tools
│   │   ├── logging/               // Logging Implementations (e.g., Zap, Logrus)
│   │   ├── repository/            // Database implementations of repositories
│   │   ├── server/                // HTTP Server setup and routing
│   │   ├── persistence/           // Persistence (e.g., database connection)
│   │   └── configuration/         // Configuration management (e.g., Viper)
│   │
│   └── interfaces/                // Interface Adapters (Controllers)
│       ├── http/                  // HTTP Controllers
│       └── grpc/                  // gRPC Controllers (if applicable)
│
├── pkg/                           // Shared Utility Libraries (optional)
├── configs/                       // Configuration Files (e.g., config.yaml)
├── scripts/                       // Scripts for CI/CD or local setup
├── .gitignore                     // Git ignore file
├── go.mod                         // Go module file
└── go.sum                         // Go dependencies file
```