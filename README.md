## A scuffed playground with Temporal SDK

### Prerequisites

- [temporal-cli](https://docs.temporal.io/cli#install)
- Make
- Python
- Go
- uv

### Start Temporal development server

Start temporal:

```sh
make temporal
```

### Run Python worker/client

Jump into the `python` directory:

```
cd python
```

Prepare virtualenv:

```sh
virtualenv env
source ./env/bin/activate
make sync
```


Start a worker:

```sh
make worker
```

Make a client call:

```sh
make client
```

### Run Go worker/client

Jump into the `go` directory:

```
cd go
```

Start a worker:

```sh
go run ./worker/main.go
```

Make a client call:

```sh
go run ./client/main.go
```
