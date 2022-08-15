package main

import (
	"fmt"

	"github.com/dev-kazuya/study-golang/module/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
