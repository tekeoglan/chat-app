# Chat App

This is a full-stack implementation of a chat application featuring real-time chat, user authentication, and more.

## Overview

The project is a monorepo with a Go backend and a Next.js frontend. It uses WebSockets for real-time communication and is fully containerized with Docker.

## Features

*   User authentication (registration and login)
*   Real-time messaging in channels
*   WebSocket for live updates
*   Friend management (e.g., add friend)
*   RESTful API for all major features

## Tech Stack

*   **Backend**: Go
    *   **Framework**: Gin (assumed)
    *   **Database**: MongoDB
    *   **Cache**: Redis
    *   **Real-time**: Gorilla WebSocket (assumed)
*   **Frontend**:
    *   **Framework**: Next.js 13+ (with App Router)
    *   **Language**: TypeScript
    *   **Styling**: Tailwind CSS
*   **Containerization**: Docker & Docker Compose

## Project Structure

The repository is divided into two main parts: `client` and `server`.

```
.
├── client/         # Next.js frontend application
│   ├── src/
│   │   ├── app/        # Pages and layouts
│   │   ├── components/ # React components
│   │   ├── lib/        # API calls and WebSocket logic
│   │   └── ...
│   ├── docker/       # Docker configuration for the client
│   └── ...
├── server/         # Go backend application
│   ├── api/          # Controllers, routes, middleware
│   ├── cmd/          # Main application entrypoint
│   ├── model/        # Data structures and models
│   ├── repository/   # Database logic
│   ├── service/      # Business logic
│   ├── ws/           # WebSocket handling
│   ├── docker/       # Docker configuration for the server
│   └── ...
└── README.md
```

### Server

The `server` directory contains the Go backend. It follows a standard layered architecture:

*   `api/`: Handles HTTP requests, routing, and middleware.
*   `service/`: Contains the core business logic of the application.
*   `repository/`: Abstract layer for data access (MongoDB and Redis).
*   `model/`: Defines the data structures used throughout the application.
*   `ws/`: Manages WebSocket connections, rooms, and message broadcasting.
*   `bootstrap/`: Handles application startup and initialization of services like the database and environment variables.

### Client

The `client` directory contains the Next.js frontend.

*   `src/app/`: Uses the Next.js App Router for routing and layouts.
*   `src/components/`: Contains reusable UI components.
*   `src/lib/`: Manages state, API requests, and WebSocket communication.

## Getting Started

### Prerequisites

*   Docker and Docker Compose
*   Go (if running the server locally)
*   Node.js and npm (if running the client locally)

### Running with Docker (Recommended)

This is the easiest way to get both the client and server running.

    ```bash
    docker-compose -f docker-compose.yaml up -d --build
    ```

The client will be available at `http://localhost:3000` and the server at `http://localhost:8080` (or as configured in your `.env` files).

### Running Locally

#### Server

1.  Navigate to the `server` directory: `cd server`
2.  Install dependencies: `go mod tidy`
3.  Set up your `.env` file in `conf/` for database connections.
4.  Run the server: `go run cmd/main.go`

#### Client

1.  Navigate to the `client` directory: `cd client`
2.  Install dependencies: `npm install`
3.  Run the development server: `npm run dev`

The application will be accessible at `http://localhost:3000`.
