
# Instruction

The tantan-demo implements a simple RESTful service in Golang for a simplified Tantan backend:
adding users and swiping other people in order to find a match.

*It is only a demo.*

# Installation

## Dependencies

- Go 1.8+
- PostgreSQL
- govendor
- curl (only for test, can by replaced by any other HTTP client)

## Build & Run

- Unzip tantan-demo.zip to directory `$GOPATH/src/tantan-demo/`.
- `cd $GOPATH/src/tantan-demo/`.
- Create PostgreSQL tables and indexes by `demo.sql`.
- Revise file `config/config.json` according to your enviroment.
- Build by `go build`.
- Start the server by `./tantan-demo -c config/config.json`.
- By default, the log is outputted to the console.

# Test & Examples

## User

### List all users

Request command:

`curl -XGET http://localhost:8000/api/v1/users`

Response body:

```json
[
    {
        "id": "7",
        "name": "Alice",
        "type": "user"
    },
    {
        "id": "8",
        "name": "Bob",
        "type": "user"
    },
    {
        "id": "9",
        "name": "Carol",
        "type": "user"
    },
    {
        "id": "10",
        "name": "Dale",
        "type": "user"
    }
]
```

### Create a user

Request command:

`curl -XPOST -H 'Content-Type: application/json' -d '{"name":"Eric"}' http://localhost:8000/api/v1/users
`

Response body:

```json
{
    "id": "11",
    "name": "Eric",
    "type": "user"
}
```

## Relationship

### List a user's all relationships

Request command:

`curl -XGET http://localhost:8000/api/v1/users/1/relationships`

Response body:

```json
[
    {
        "user_id": "2",
        "state": "liked",
        "type": "relationship"
    },
    {
        "user_id": "3",
        "state": "liked",
        "type": "relationship"
    },
    {
        "user_id": "10",
        "state": "disliked",
        "type": "relationship"
    }
]

```

### Create/update relationship state to another user

Request command:

`curl -XPUT -H 'Content-Type: application/json' -d '{"state":"liked"}' http://localhost:8000/api/v1/users/1/relationships/2
`

Response body:

```json
{
    "user_id": "2",
    "state": "matched",
    "type": "relationship"
}
```

