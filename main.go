package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
    "strconv"
    "github.com/adshao/go-binance/v2"
)

type Credentials struct {
    apiKey      string `json:"apiKey"`
    secretKey   string `json:"secretKey"`
}

func check_error(err error) {
    if err != nil {
        panic(err)
    }
}

func get_keys() (apiKey string, secretKey string) {
    jsonFile, err := os.Open("credentials.json")
    check_error(err)
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    var credentials Credentials
    json.Unmarshal(byteValue, &credentials)

    return credentials.apiKey, credentials.secretKey
}

func get_price(client *binance.Client, symbol string) float64 {
    res, err := client.NewListBookTickersService().Symbol(symbol).Do(context.Background())

    check_error(err)

    bidPrice, _ := strconv.ParseFloat(res[0].BidPrice, 64)
    bidQuantity, _ := strconv.ParseFloat(res[0].BidQuantity, 64)
    askPrice, _ := strconv.ParseFloat(res[0].AskPrice, 64)
    askQuantity, _ := strconv.ParseFloat(res[0].AskQuantity, 64)

    return (bidPrice * bidQuantity + askPrice * askQuantity) / (bidQuantity + askQuantity)
}

func get_percentage_change(client *binance.Client, symbol string, interval string, lastPrice float64)float64 {
    klines, err := client.NewKlinesService().Symbol(symbol).Interval(interval).Do(context.Background())

    check_error(err)

    openPrice, _ := strconv.ParseFloat(klines[len(klines) - 1].Open, 64)

    return (lastPrice - openPrice) / (openPrice) * 100
}

func main() {
    apiKey, secretKey := get_keys()
    client := binance.NewClient(apiKey, secretKey)

    price := get_price(client, "BTCUSDT")
    percentage := get_percentage_change(client, "BTCUSDT", "1h", price)

    if percentage >= 2.0 {
        fmt.Printf("{\"state\":\"Good\",\"text\":\" ₿ %.2f $\"}\n", price)
    } else if percentage <= -2.0 {
        fmt.Printf("{\"state\":\"Critical\",\"text\":\" ₿ %.2f $\"}\n", price)
    } else {
        fmt.Printf("{\"state\":\"Idle\",\"text\":\" ₿ %.2f $\"}\n", price)
    }
}
