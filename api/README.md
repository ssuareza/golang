# api

A basic API using go-chi.

It includes:

- A logger middleware for requests.
- A basic auth middleware.

## Usage

```bash
make build
build/api
```

## Endpoints

### GET /healthz

Returns the health of the API.

### GET /auth

To test the auth middleware.
