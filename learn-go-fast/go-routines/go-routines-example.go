package main

import (
	"fmt"
	"time"
)

func someTask(data int) {
   time.Sleep(2 * time.Second)
   fmt.Printf("Task %d executed\n", data)
}

func main() {

   for i := 0; i < 5; i++ {
      someTask(i)
   }

}