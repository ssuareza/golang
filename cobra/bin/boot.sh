#!/usr/bin/env sh

# autowatch and run tests
find . | entr -r go test $(go list -buildvcs=false  ./...)
