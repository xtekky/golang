package main

import (
    "fmt"
	"time"
	"strconv"
)

func main() {
	start := time.Now()

	n := 1
	length := 9999999999

	for n < length {
		n += 1
	}
	//fmt.Println(n) // 8 (1*2*2*2)

    elapsed := time.Since(start)
    fmt.Printf("  [ > ] TTC: %s | Calculations: %s\n\n\n\n\n",  elapsed, strconv.Itoa(n))

}
