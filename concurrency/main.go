package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	keyboard := bufio.NewReader(os.Stdin)
	for line, _, err := keyboard.ReadLine(); err == nil; {
		go handle(string(line))
		line, _, err = keyboard.ReadLine()
	}
}

func handle(input string) {
	time.Sleep(time.Second * 5)
	fmt.Println("Handling: " + input)
}
