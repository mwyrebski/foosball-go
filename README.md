# Foosball (in Go)

Simple web service with API for tracking game stats for three-set Foosball games.

This is a very simple version in Go using only built-in libraries. Data is not persisted
and will be wiped out after server restart.

## Domain

Basic assumptions:
* each `Game` consists of maximum number of 3 `Set`s
* each `Set` consists of maximum number of 10 `Goal`s
* team wins a `Set` when shooting 10 goals
* team wins a `Game` by winning any of two `Set`s
* `Game` can be in three states:
  - `NotStarted` - when a `Game` is newly created and until first goal is shot
  - `InProgress` - after first shot and until `Game` is `Finished`
  - `Finished` - when any of teams win whole `Game`

## Running

Tests can be run with:

```
go test
```

To start the server:

```
go run .
```

or with published executable:

```
go build
foosball-go
```
