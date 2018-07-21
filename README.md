# di-example

REST API to manage a list of car stored in mongodb.

It shows how to use [sarulabs/di](https://github.com/sarulabs/di) in a web application.

More explanations can be found in this post:

- [How to write a REST API in Go with DI](https://www.sarulabs.com/post/3/2018-08-02/how-to-write-a-rest-api-in-go-with-di.html).

## Routes

| Method  | URL         | Role        | JSON Body example                    |
| ------- | ----------- | ----------- | ------------------------------------ |
| GET     | /cars       | List cars   |                                      |
| POST    | /cars       | Insert car  | {"brand": "audi", "color": "black"}  |
| GET     | /cars/{id}  | Get car     |                                      |
| PUT     | /cars/{id}  | Update car  | {"brand": "audi", "color": "black"}  |
| DELETE  | /cars/{id}  | Delete car  |                                      |

## Dependencies and vgo

The dependencies are handled with `vgo`. You need to install vgo first:

```sh
go get -u golang.org/x/vgo
```

Then you can use `vgo run main.go` to start the application.

## Environment

You need to set some environment variables before running this application:

- **SERVER_PORT**: the port for the rest api (eg: `8080`)
- **MONGO_URL**: the mongodb address (eg: `127.0.0.1:27017`)

To use the default values you can use the `default_env.sh` script:

```sh
source default_env.sh
vgo run main.go
```

## Run with docker

It is also possible to run this application with docker compose:

```sh
docker-compose -f docker/compose.yml up --build
```
