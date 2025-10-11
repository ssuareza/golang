# gcs

This is a test I did to test download times from GCS.

1. `cmd/gcs/main.go`: download file from bucket.
2. `cmd/chunks/main.go`: download file splitting in chunks.
3. `cmd/grpc/main.go`: download using GRPC.

## Usage

1.- Authenticate with Google:

```sh
gcloud auth application-default login
```

2.- And then run any of the scripts:

```sh
go run cmd/gcs/main.go
```

```sh
go run cmd/chunks/main.go
```

```sh
go run cmd/grpc/main.go
```
