package main

import (
	"easygenerate/generate"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		panic(err)
	}
	pog := generate.NewJavaPOGenerator(string(file))
	dtog := generate.NewJavaDTOGenerator(string(file))
	mog := generate.NewMySQLGenerator(string(file))
	fmt.Println(pog.Build())
	fmt.Println(dtog.Build())
	fmt.Println(mog.Build())
}
