# Todo Application

A simple Todo application built with Go (Golang) and DynamoDB. The application supports creating, updating, deleting, and fetching todos for authenticated users.

## Features

- User Authentication (Registration and Login)
- CRUD operations on todos:
  - Create a new todo
  - Update an existing todo
  - Delete a todo
  - Fetch todos by ID
- Uses DynamoDB as the database

---

## Prerequisites

Ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/dl/) (1.20 or higher)


---

## Getting Started

### 1. Clone the Repository
```bash
git clone git@github.com:akshay-kgen/todo-app.git
cd todo-app
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory and update it as per the env.sample:


### 3. Install Dependencies

Run the following command to install the Go dependencies:
```bash
go mod download
```

### 4. Run the Application Locally

#### Using Docker Compose
1. Start the services (app and DynamoDB):
   ```bash
   docker-compose up -d
   ```

2. Verify that the application is running at `http://localhost:3000`.


## License

This project is licensed under the [MIT License](LICENSE).

