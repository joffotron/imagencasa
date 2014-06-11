package main

import (
	"fmt"
	server "imagencasa/server"
)

func main() {
	fmt.Println("Imagencasa started!")

	go server.Start()

	select {} //Sleep
}

