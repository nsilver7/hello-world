package main

import (
	"fmt"
	"time"

	"github.com/nsilver7/stringutil"
)

func main() {
	fmt.Println("nonce: ", time.Now().Unix()*10000)
	fmt.Printf("hello, world\n")
	fmt.Printf(stringutil.Reverse("hello, world")+"\n")
}
