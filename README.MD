
# toGO-List

toGO-List is a todo list HTTP API written in Go, developed mainly for self-learning different practices in writing in Go programming languages.

Technology stack used: **Golang, Gin. Gorm, Swagger, PostgreSQL**
## Installation

To install Togo List, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/aWatLove/togo-list.git
```

2. Navigate to the project directory:

```bash
cd togo-list
```

3. Install the required dependencies:

```bash
go get -d ./...
```
4. Create a `.env` file in the root directory of the project and add database connection details. Here is an example of what the `.env` file should look like:
```
DB_USERNAME=username
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=togo-list
DB_SSLMODE=disable
```


5. Build the project:

```bash
go build .\cmd\main.go
```

## Usage

To run the toGO-List API, use the following command:

```bash
./main.exe
```

The API will be available at `http://localhost:8080`.

## API Documentation

The toGO-List API documentation can be found in the `docs/` directory. You can access the Swagger UI by visiting `http://localhost:8080/swagger/index.html` in your web browser.
