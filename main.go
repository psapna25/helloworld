package main

import (
    "fmt"
    "time"
)
func helloworld() string {
	return "Hello World!!"
}

func main() {
	fmt.Println(helloworld())
	time.Sleep(600 * time.Second)
}
