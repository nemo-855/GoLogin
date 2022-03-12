package main

import (
	"fmt"
)

func main()  {
	fmt.Print(greeting("japan") + "world")
}

func greeting(country string) string {
	var greeting string
	switch country {
		case "Japan": greeting = "こんにちは"
		case "America": greeting = "Hello"
		default: greeting = "知らねえよ"
	}
	return greeting
}