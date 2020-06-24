package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // must have a trailing comma in multi line
	}
	fmt.Println(len(stocks))
	fmt.Println(stocks["MSFT"])
	fmt.Println(stocks["TSLA"]) // 0
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	//set
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)
	delete(stocks, "AMZN")
	fmt.Println(stocks)
	for key := range stocks {
		fmt.Println(key, ":", stocks[key])
	}

}
