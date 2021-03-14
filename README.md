# PomodoroGO Server

PomodoroGo Server is backend application for pomodoro technique. It aims to store your working time with membership and give statistics in details for efficient. Also, will be developed client application to achieve this aims. After finished the client application, publish this project with domain (pomodorogo.app). This project is
released under the terms of the GNU. PomodoroGo Server develops with Go programming language and PostgreSQL for the database.

## API Endpoint

|Method  |Path                 |Description         |   |   |
|---     |---                  |---                 |---|---|
|POST    |/v1/auth/signup      |                    |   |   |
|POST    |/v1/auth/signin      |                    |   |   |
|POST    |/api/v1/tags         |                    |   |   |
|GET     |/api/v1/tags/{id}    |                    |   |   |
|PUT     |/api/v1/tags         |                    |   |   |
|DELETE  |/api/v1/tags/{id}    |                    |   |   |
|POST    |/api/v1/statistics   |                    |   |   |


## Environment

### Server 
* $SERVER_ADDRESS
* $ACCESS_TOKEN_PRIVATE_KEY_PATH
* $ACCESS_TOKEN_PUBLIC_KEY_PATH
* $REFRESH_TOKEN_PRIVATE_KEY_PATH
* $REFRESH_TOKEN_PUBLIC_KEY_PATH
* $JWT_EXPIRATION

### Database  
* $DB_HOST
* $DB_NAME
* $DB_PASSWORD
* $DB_PORT

## Development 

```shell
git clone https://github.com/ahmetcancicek/pomodorogo-server.git
cd pomodorogo-server
go run ./cmd/server/main.go
```

## Docker

```shell
docker-compose up --build
```

## Contributions

PomodoroGo is open source, and need to contribute. Especially for the client application, If you want, suggest any idea and implement new features.

