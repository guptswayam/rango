package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Rango!")

	fmt.Println("Rango is working...")

	// time.Sleep(time.Second * 5)

	DeleteFiles()

	fmt.Println("Closing Rango!")

	var isFinish string = ""
	for isFinish != "yes" {
		fmt.Println("Type yes to close")
		fmt.Scanln(&isFinish)
	}

}
