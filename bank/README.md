# bank

A cli to create a CSV file from the bank transactions exported data.

Every bank uses its format to export data. This application gets that data and creates a CSV with the description and amount of the transactions.

## Usage

1. Build application:

```sh
make build
```

2. And run:

```sh
build/bank --name=bbva --file=data/bbva.xlsx --month=03
CSV file created successfully on /tmp/bbva.csv
```

```sh
cat /tmp/bbva.csv

Parking;0,66
Parking;0,95
Parking;-0,95
Parking;-1,75
Peaje;-3,05
Parking;-2,4
Nomina;1000
Bizum;-200
Parking;-0,55
Escola;-129,71
Luz;-77,44
Online;-24,99
Bizum;25
Bizum;56
Bizum;-11,3
Telefono;-54,9
Club;-84,4
```
