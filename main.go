package main

import (
	"fmt"
	"os"

)

func readFile()string{
	fmt.Println("Reading file")
	filename := "Test.txt"
	filedata ,err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(filedata)
}

func wordCounter(filedata string){
	count := 0
	for _, v := range filedata {
		if v == ' '{
			count++
		}
	}
	fmt.Println("Number of words :", count)

}


func main() {
	filedata := readFile()
	wordCounter(filedata)
}

