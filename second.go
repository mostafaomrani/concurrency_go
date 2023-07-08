package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms..."
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every two minute ..."
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
