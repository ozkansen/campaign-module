# Campaign Auto Pricing Module

## How to running the application

### With Docker

```sh
docker build -t campaign_application:latest .
docker run campaign_application:latest
```

### Without Docker

```sh
go mod tidy
go test ./...
go build -o campaign_application ./cmd

./campaign_application
```

## Check Output

```text
date:  2000-01-01 00:00:00 +0000 UTC
Product created; code P1, price 100, stock 100
Campaign created; name C1, product P1, duration 1, limit 20, target sales count 100
campaign service create: campaign already exists
Product P1 info; price 80, stock:100
Order created; product P1, quantity 10
Product P1 info; price 84, stock:90
Campaign C1 info; Status Active, Target Sales 100, Total Sales 10, Turnover 800, Average Item Price 80
Order created; product P1, quantity 10
Campaign C1 info; Status Ended, Target Sales 100, Total Sales 20, Turnover 1640, Average Item Price 82
Campaign C1 info; Status Ended, Target Sales 100, Total Sales 20, Turnover 1640, Average Item Price 82
Campaign C1 info; Status Ended, Target Sales 100, Total Sales 20, Turnover 1640, Average Item Price 82
Product P1 info; price 100, stock:80
Order created; product P1, quantity 10
Order created; product P1, quantity 10
Product P1 info; price 100, stock:60
Campaign C1 info; Status Ended, Target Sales 100, Total Sales 40, Turnover 3640, Average Item Price 91
```
