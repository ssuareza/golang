#!/usr/bin/env sh

# autowatch and run tests
find . | entr -r go test $(go list -buildvcs=false  ./...)  &

# autowatch and run main process
find . | entr -r go run cmd/api/main.go

