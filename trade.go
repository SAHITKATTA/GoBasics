package main

import "fmt"

// Trade is a trade in stockss
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// Value return the Trade value
func (t *Trade) Value() float64 { // pointer so that to get pointer reciever
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)
	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)
	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
	fmt.Println("-----------")
	fmt.Println(t1.Value())
}
