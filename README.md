## Why ? 

We come from nodejs with nestJs framework. 

We are looking for solution to reproduce response and comportement of nestJs in a golang api project. 

*** So architecture is not remocommanded by gin ** 

## Todo

- Error api on fields
- Swagger
- Unit and functionals tests
- ReadyZ : Check if database is ready 

## Presentation

This project describe an ideal architecture of api :  

- Security
- Authentication by JWT
- Model, Controller, Service
- Roles and permissions
- Tests
- Dockerfile
- CI with sonnar

## Start 

VS Code determine path of package from route go project. So you need to open your vs code project at root of project go. 

Start by install libs : 

```bash 
go get . 
```

Create a `pulic.pem` file (for valide JWT token) in `/commons/guard/jwt`

Copie `.env.model` => `.env` and complete information

And start server by : 

```bash 
set -o allexport; source .env; set +o allexport
go run .
```

## Build docker image : 
```bash
$ docker build -t api .
$ docker run -it --rm --name api api
```

## Description

[TODO]


## Thank's

- https://github.com/vsouza/go-gin-boilerplate
- https://github.com/yhagio/go_api_boilerplate
- https://github.com/Massad/gin-boilerplate
- https://github.com/daystram/go-gin-gorm-boilerplate
- https://github.com/imbarwinata/go-gin-gorm-boilerplate/
- https://github.com/dedidot/gorm-gin




