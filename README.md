# i3-status-rs-crypto-price
A block to add in your i3bar with i3status-rs.
It use the binance-v2 API.

# Install dependencies
```sh
go get github.com/adshao/go-binance/v2
```

# How to use
1. Clone the repositoty
2. Add your credentials in a file credentials.json in the cloned repository
```json
{
    "apiKey": "your binance api key"
    "secretKey": "your binance secret key"
}
```
3. Build the project
```sh
# In the cloned repository
go build main.go
```
4. Add in your status.toml
```toml
[[block]]
block = "custom"
command = "<path/to/main>"
json = true
interval = 1
```
