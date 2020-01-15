# Cisco Api

An api which communicate with sqlite database

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Before starting make sure you have below packages installed:

- go version go1.13.1 linux/amd64
Optional:
- Docker version 18.09.7
- GNU Make 4.1
  
### Installing

First, make sure your github profil have a ssh key to allow `clone`

```
git clone git@github.com:Royalsspirit/cisco-api.git
```

This repository is made to run in containers but the classic way is still available. Just follow corresponding section

#### Makefile and docker

Make sure that requirements are satisfied.

To run the api, just type:
- `make up`

it will expose api server at port `8080`.

#### Just with go

- Just run: `go run cmd/api/main.go`. 
  
If you want to use other db than the db available in repository you could add a environnement variable in the previous command:
- `DB=PAHT_TO_DBFILE go run cmd/api/main.go`

#### Endpoints

There are 4 available endpoints:

- `GET /character` to get list of people with their vehicles and species
- `PUT /character/{id}` to update a people properties (send one of object returned by the previous endpoint)
- `DELETE /character/{id}` to delete a people identified by his id
- `POST /character` to create a people
  
## Running the tests

`- make unit-tests`

## Built With

- [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Watches your .go files in a directory and invokes go build if a file changed. Nothing more.
- [sqlite3](https://github.com/mattn/go-sqlite3) - sqlite3 driver conforming to the built-in database/sql interface
- [Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers
- [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library

## Contributing

TODO

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Acknowledgments

TODO
