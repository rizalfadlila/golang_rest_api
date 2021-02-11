# Example Rest API

## Setup Local Environment

### Prerequisites

+ Git
+ Go (minimum 1.13 or above, make sure you set your local machine environment with `GO111MODULE=on`)
+ Docker compose
+ Docker
+ [Mockgen](https://github.com/golang/mock)
+ [Swag](https://github.com/swaggo/swag)

### Installation

+ Clone this repository.

+ Run `docker-compose` to build the dependency docker instances:
```bash
$ docker-compose up --build
```

+ Make the environment files. Adjust your local configuration.
```bash
$ cp config.yaml.tpl .env
```

+ Run these commands to get all the app dependency's
```bash
$ swag i
$ go get -v ./...
```

+ Run the app. The app will run inside the local machine with exposed port configured in the env (by default: [localhost:3030](http://localhost:3030))
```bash
$ go run main.go
```

### Development

#### Documentation
Access [localhost:3030/swagger/index.html](http://localhost:3030/swagger/index.html) to read the API docs.

#### Migration
Reference to this [doc](https://github.com/xakep666/mongo-migrate#use-case-1-migrations-in-files) to create a new migration file.

#### Mock an interface
Use `mockgen` to create a new mock. Tips: place mock file inside [mocks](mocks) directory
```bash
$ mockgen -package=mock -source=/path/to/interface/file -destination=/path/to/generated/mock/file
```

#### Add dependency
Use `go mod` as dependency tool.
```bash
$ go get github.com/gin-gonic/gin
```

#### Unit test
Use `go test` to run unit test.
```bash
$ go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
```
