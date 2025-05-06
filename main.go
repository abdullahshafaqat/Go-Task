package main

import (
	"GOTASKS/chunks"
	"GOTASKS/combine"
	"GOTASKS/separate"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	fmt.Println("Reading file")
	filename := "Test.txt"
	fileData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	Str := string(fileData)
	go separate.WordCounter(Str, &wg)
	go separate.LineCounter(Str, &wg)
	go separate.SpaceCounter(Str, &wg)
	go separate.PuncCounter(Str, &wg)
	go separate.SpeciCounter(Str, &wg)
	go separate.SentenceCounter(Str, &wg)
	go separate.VowelCounter(Str, &wg)
	go separate.ConsonCounnter(Str, &wg)
	go separate.DigitCounter(Str, &wg)
	go separate.ParaCounter(Str, &wg)
	wg.Wait()
	timeTaken := time.Since(startTime)
	fmt.Println("Time Taken After Using Goroutine: ", timeTaken)

	startTime2 := time.Now()
	resultchann := make(chan []int)
	go combine.AnalyzeText(Str, resultchann)
	var names = [10]string{"Words", "lines", "Spaces", "Punctuations", "Special Characters", "Sentences", "Vowels", "Consonants", "Digits", "Paragraphs"}
	value := <-resultchann
	for i := 0; i < 10; i++ {
		fmt.Println(names[i], value[i])
	}

	timeTaken2 := time.Since(startTime2)
	fmt.Println("Time Taken After Using Channels & Goroutines: ", timeTaken2)

	timeStart3 := time.Now()
	chunks.Chunks(Str)
	timeTaken3 := time.Since(timeStart3)

	fmt.Println("Time taken after making chunks :", timeTaken3)

}
