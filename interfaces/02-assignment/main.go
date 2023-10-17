package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	filename := os.Args[1]
	// 01 Way
	// res, err := os.ReadFile(filename)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(res))

	f, e := os.Open(filename)

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
