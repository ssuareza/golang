# wise

A cli to use Wise API to obtain data of your account.

## Usage

1. Build application:

```sh
make build
```

2. Create configuration file on "$HOME/.config/wise/wise.yml":

```yaml
api_endpoint: https://api.wise.com
api_key: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
profile_id: 12345678
```

3. And run:

```sh
build/wise
```

```text
> Summary
- Balance USD: 10000.20
- Balance EUR: 8000.15
- Rate USD to EUR: 0.9216

> Card transactions
- 2024-05: 300.64 EUR (T-shirt 128.17 EUR)
- 2024-04: 1005 EUR (New bike 830.14 EUR)
- 2024-03: 344 EUR (Market 127.06 EUR)
- 2024-02: 217.75 EUR (Doctor 166 EUR)
- 2024-01: 150.99 EUR (Market 117.74 EUR)
- 2023-12: 166.51 EUR (Amazon 156.37 EUR)
```

You can also filter by "label":

```sh
build/wise --label bike
```
