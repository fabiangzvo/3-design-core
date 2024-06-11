# 3-design core

This is a backend application built with Go, using `gin-gonic` as the web framework, `gorm` as the ORM for database interactions, and `golint` for ensuring code quality. The application is containerized using Docker and managed with `docker-compose`.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Enviroment Variables](#enviroment-variables)
- [Installation](#installation)
- [Run Locally](#run-locally)
- [Running Tests](#running-tests)
- [Code Quality](#code-quality)
- [Tecnologies used](#tecnologies-used)
## Prerequisites

Ensure you have the following installed on your local machine:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
## enviroment variables

Set up the environment variables (create a `.env` file in the project root):

| Variable | example     | Description                |
| :-------- | :------- | :------------------------- |
| `SERVER_PORT` | `8080` | **Required**. port through which the server will be exposed |
| `DB_NAME` | `database-name` | **Required**. database name to connect |
| `DB_USER` | `user-example` | **Required**. database user to connect |
| `DB_PASSWORD` | `dummy-password` | **Required**. database password to connect |
| `DB_PORT` | `5432` | **Required**. database port to connect |
| `DB_HOST` | `localhost` | **Required**. database host |

## Run Locally

Clone the project

```bash
  git clone https://github.com/fabiangzvo/3-design-core.git
```

Go to the project directory

```bash
  cd ./3-design-core
```

run docker container

```bash
   docker compose -f ./docker-compose-dev.yaml up --build
```

The application should now be running and accessible at `http://localhost:8080`.

## Running Tests

To run tests, run the following command

```bash
  go test ./...
```


## Code Quality

To ensure the code adheres to best practices, we use golint. to configure follow these steps:

- to install run the following command:

    ```bash
        go install golang.org/x/lint/golint@latest
    ```
- In vscode, go to ``File > Preferences > Settings`` and search by ``go.lintTool`` and select ``golint`` option.

Run the following command to lint all code:
```bash
    golint ./...
```
## Tecnologies used

the following technologies were fundamental to the creation of this project:

- [Go](https://golang.org/)
- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Docker](https://www.docker.com/)
- [Docker-Compose](https://docs.docker.com/compose/)
- [Golint](https://github.com/golang/lint)
- [Godotenv](https://github.com/joho/godotenv)
- [Logrus](https://github.com/sirupsen/logrus)

