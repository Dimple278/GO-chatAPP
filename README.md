# CLI Chat Application

## Overview

This is a Command Line Interface (CLI) chat application built using Golang. It allows users to register, log in, and chat with other logged-in users. The chat messages and user information are stored in a PostgreSQL database, providing persistent storage for registered users and chat history. The application also includes commands to view chat history, list online users, and log out.

## Features

- **User Registration**: New users can register with an email, username, and password.
- **User Login**: Users can log in with their username and password to access the chat.
- **Real-time Chat**: Logged-in users can send and receive messages in real time with other online users.
- **Chat History**: Users can view their past chat history.
- **List Users**: View all currently logged-in users.
- **User Logout**: Users can log out, disconnecting them from the chat server.
- **Persistent Storage**: User information and chat history are saved in a PostgreSQL database.

## Technologies and Libraries

- **Golang**: The core language used for building the CLI and server functionality.
- **PostgreSQL**: Used for storing user data and chat history.
- **Cobra**: A library for building CLI commands and argument parsing.
  - [github.com/spf13/cobra](https://github.com/spf13/cobra)
- **Socket.io**: For real-time communication between the client and server.
  - Server: [github.com/googollee/go-socket.io](https://github.com/googollee/go-socket.io)
  - Client: [github.com/hesh915/go-socket.io-client](https://github.com/hesh915/go-socket.io-client)

## Setup and Installation

### Prerequisites

1. **Golang**: Ensure you have Go installed on your system.
   - [Download Golang](https://golang.org/dl/)
2. **PostgreSQL**: Install PostgreSQL for storing user data and chat history.
   - [Download PostgreSQL](https://www.postgresql.org/download/)
3. **Go Modules**: Enable Go modules by running `go mod init`.

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-repo/cli-chat-app.git
   cd cli-chat-app
   ```

2. **Install Dependencies**:

   Install the required Go libraries:

   ```bash
   go mod tidy
   ```

3. **Set Up PostgreSQL**:

   Create a PostgreSQL database:

   ```sql
   CREATE DATABASE chatapp;
   ```

4. **Update Configuration**:

   Create a `.env` file in the root directory with the following details:

   ```env
   DATABASE_URL=postgres://<DB_USERNAME>:<DB_PASSWORD>@localhost:5432/<DB_NAME>
   ```

### Running Migrations

Before starting the application, ensure the database schema is up to date by running migrations:

1. **Install Migrate Tool**:

   - Install `golang-migrate` by running:
     ```bash
     go get -u github.com/golang-migrate/migrate/v4
     ```

2. **Run Migrations**:

   - To apply all pending migrations, use:
     ```bash
     migrate -path ./migrations -database ${DATABASE_URL} up
     ```

3. **Rollback Migrations**:

   - If needed, you can rollback the last migration by running:
     ```bash
     migrate -path ./migrations -database ${DATABASE_URL} down
     ```

4. **Run the Application**:

   ```bash
   go run main.go
   ```

## Commands

### `register`

Register a new user.

```bash
go run main.go register
```

It will prompt you for:

- Email
- Username
- Password

### `login`

Log in an existing user.

```bash
go run main.go login
```

It will prompt you for:

- Username
- Password

Once logged in, you can start chatting with all the other logged-in users.

### `chat`

After logging in, the user can start chatting with other logged-in users by entering messages directly. Commands can be used to interact with the chat server.

### Chat Commands (Options)

After logging in, users can use the following options prefixed by `/`:

- **`/history`**: View your chat history.

  ```bash
  /history
  ```

- **`/listusers`**: View all currently logged-in users.

  ```bash
  /listusers
  ```

- **`/logout`**: Log out the current user.

  ```bash
  /logout
  ```

## Additional Notes

- The `Socket.io` server runs on `http://localhost:8000`.
- Ensure the database is running before starting the application.
- The chat messages are broadcasted in real time to all online users.
- The `cobra` library is used to handle the CLI commands and their flags.
