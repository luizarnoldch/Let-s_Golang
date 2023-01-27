package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    fmt.Println("Enter a number...")
    reader := bufio.NewReader(os.Stdin)
    
    input, err := reader.ReadString('\n')
    //Let's golang

    if err != nil {
        log.Fatal("Error while reading input")
    }

    fmt.Print("Your input was: ", input)
    //Your input was: Let's golang
}