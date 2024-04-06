# My Go Application

## How to

- **To create a new migration**:
  ```bash
  make migrate-create
  ```

- **To apply all migrations**:
  ```bash
  make migrate-up
  ```

- **To roll back the last migration**:
  ```bash
  make migrate-down
  ```

- **To build the Go application**:
  ```bash
  make build
  ```
  This compiles the application and generates a binary with the name specified in `BINARY_NAME`.

- **To run the Go application**:
  ```bash
  make run
  ```
  This command runs the compiled binary. 

Remember to copy `.env.example` to `.env` and replace the values

## Run PostgreSQL with Docker

```sh
docker run --name my-postgresql \
  -e POSTGRES_PASSWORD=mysecretpassword \
  -e POSTGRES_DB=auth_management\
  -p 15432:5432 \
  -d postgres:13
````