# filesplit

A cli to split and save a file into Memcached.

## Usage

### Start

```bash
make start
```

### Set file

```bash
docker compose exec filesplit go run cmd/filesplit/main.go set file.txt
```

### Get file

```bash
docker compose exec filesplit go run cmd/filesplit/main.go get file.txt
```

### Delete file

```bash
docker compose exec filesplit go run cmd/filesplit/main.go del file.txt
```

## Dependencies

- Docker

## Useful commands

Create a 15MB file:

```bash
dd if=/dev/zero of=file.txt bs=15MB count=1
```
