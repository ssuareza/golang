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
   build/bank --name=bbva --file=data/bbva.xlsx --date=01/03/2024
   CSV file created successfully on /tmp/bbva.csv
   ```

More examples:

```sh
build/bank --name=ing --file=data/ing.xlsx --date=01/05/2024
```

```sh
build/bank --name=wise --date=01/03/2024
```

## Wise

In the case of Wise the application gets the data from the Wise API, no need to use a spreadsheet.

The only requirement is to create a configuration file on "$HOME/.config/bank/wise.yml":

```yaml
api_endpoint: https://api.wise.com
api_key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
profile_id: 12345678
```
