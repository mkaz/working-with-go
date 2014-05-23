package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "../extras/rabbits.txt"

	f, _ := os.Open(filepath)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
