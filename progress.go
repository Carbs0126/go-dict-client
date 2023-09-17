package main

import (
	"fmt"
	"time"
)

func PrintProgress(c <-chan interface{}) {

	progressRunes := []rune("/-\\|")
	index := 0
	requestReturn := false
	for {
		select {
		case <-c:
			requestReturn = true
		default:
			fmt.Printf("\r%s", string(progressRunes[index]))
			index++
			index = index % 4
			time.Sleep(100 * time.Millisecond)
		}
		if requestReturn {
			return
		}
	}
}
