## Temporal Python SDK playground

### Pre-requiesites

- [temporal-cli](https://docs.temporal.io/cli#install)
- Make
- Python 3.11
- uv

### Run

Prepare virtualenv:

```sh
virtualenv env
source ./env/bin/activate
make sync
```

Start temporal:

```sh
make temporal
```

Start a worker:

```sh
make worker
```

Start a client:

```sh
make client
```
