package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//set a flag tp specify the message to display
	msg := flag.String("text", "Hello there!", "message to display")
	flag.Parse()

	//display message
	fmt.Println(*msg)
	fmt.Println(os.Args[1])

}
