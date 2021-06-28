# Go Microservice Boilerplate
A Go Microservice boilerplate that includes REST Server, PosGreSQL, Migrations, Logging, Docker and Kubernetes Deployment File

## What's included:
- Go REST Server using go-json-rest
- Models + Migrations using xorm
- JSON logging with logrus
- Multi-stage Docker
- Kubernetes Deployment file

## Setup
copy the environment variables from `.example.env` to `.env` and make necessary changes and run the following command:
```
go run main.go
```

## Create a micro-service from this boilerplate
1. copy this boilerplate
2. make the repository endpoint in go.mod
3. Fix imports in all files 
```
example:
From : github.com/sahithvibudhi/full-stack/boilerplates/go-pgsql-microservice/helpers
To: github.com/<username>/<reposity>
[or]
To: github/com/<username>/<repository>/<micro-service-name>
```

## REST Endpoints
This boiler plate uses [go-json-rest](https://github.com/ant0ine/go-json-rest) for a REST HTTP Server

To add more endpoints, modify `handlers/setup.go`

Endpoints in the projects are listed in `api.http` file in `RFC2616` format

## DB Models and Migrations
This boiler plate uses [xorm](https://github.com/go-xorm/xorm) for DB connection, migrations and Queries

To add a DB Model, Copy `models/user.go` and make changes, register the `go struct` in `migrations` in `models/setup.go` to automatically create the table upon starting the server.

## Logging
This service uses [logrus](https://github.com/sirupsen/logrus)

Use it like:
```
import (
	"github.com/sahithvibudhi/full-stack/boilerplates/go-pgsql-microservice/helpers"
)

helpers.Log.Info("This is info log")
```
