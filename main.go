package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func ReadData()string{
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	
	return strings.TrimSpace(data)
}


func main(){
	inputChannel := make(chan string)
	done := make(chan bool)

	go func(){
		for{
			fmt.Print("Enter input (type 'next' to continue or 'exit' to quit): ")
			data := ReadData()
			if data == "exit"{
				done <- true
				return
			}
			inputChannel <- data
		}

	}()

	for{
		select{
		case input:= <- inputChannel:
			fmt.Println("You have entered",input)
		case <-done:
			fmt.Println("Exiting program")
			return
		}
	}

}