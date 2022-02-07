package main

import (
	"fmt"
	"io/ioutil"
)

const amountThreads = 100

func main() {
	c := make(chan int, amountThreads)
	for i := 0; i < amountThreads; i++ {
		go func() {
			namespaceBytes, err := ioutil.ReadFile("file-with-random-content.txt")
			if err != nil {
				fmt.Println("Error!!!" + err.Error())
			}
			fmt.Println(len(string(namespaceBytes)))
			c <- 1
		}()
	}
	
	for i := 0; i < amountThreads; i++ {
		<- c
	}
	close(c)
}