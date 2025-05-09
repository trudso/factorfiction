package main

import (
	"fmt"

	"github.com/trudso/factorfiction/engine"
)

func main() {
	reply, err := engine.GenerateQuestion() 
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
