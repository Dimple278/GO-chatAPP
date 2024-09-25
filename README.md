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
   git clone https://github.com/Dimple278/GO-chatAPP
   cd GO-chatApp
   ```

2. **Install Dependencies**:

   Install the required Go libraries:

   ```bash
   go mod tidy
   ```

3. **Set Up PostgreSQL**:

   Create a PostgreSQL database and run the necessary migrations to set up the users and chat history tables.

   ```sql
   CREATE DATABASE chat_app;
   \c chat_app;

   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       email VARCHAR(255) UNIQUE NOT NULL,
       username VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL
   );

   CREATE TABLE messages (
       id SERIAL PRIMARY KEY,
       user_id INT REFERENCES users(id),
       message TEXT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

4. **Update Configuration**:

   Create a `.env` file in the root directory with the following details:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=chat_app
   ```

5. **Run the Application**:

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

Once logged in, you can start sending chat messages.

### `listusers`

View all currently logged-in users.

```bash
go run main.go listusers
```

### `history`

View the chat history.

```bash
go run main.go history
```

### `logout`

Log out the current user.

```bash
go run main.go logout
```

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

### Messages Table

```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Additional Notes

- The `Socket.io` server runs on `http://localhost:8000`.
- Ensure the database is running before starting the application.
- The chat messages are broadcasted in real time to all online users.
- The `cobra` library is used to handle the CLI commands and their flags.

## Contributing

Feel free to submit issues or pull requests if you want to contribute to this project.

## License

This project is licensed under the MIT License.
