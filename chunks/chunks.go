package chunks

import (
	"GOTASKS/combine"
	"fmt"
)

func Chunks(fileData string) {
	chunkSize := len(fileData) / 4

	var names = [10]string{"Words", "lines", "Spaces", "Punctuations", "Special Characters",
		"Sentences", "Vowels", "Consonants", "Digits", "Paragraphs"}
	chunk := []string{
		fileData[:chunkSize],
		fileData[chunkSize : chunkSize*2],
		fileData[chunkSize*2 : chunkSize*3],
		fileData[chunkSize*3:],
	}

	for i := 0; i < 4; i++ {
		resultchann := make(chan []int)
		go combine.AnalyzeText(chunk[i], resultchann)
		fmt.Println("Chunk No", i)
		value := <-resultchann
		for j := 0; j < 10; j++ {
			fmt.Println(names[j], value[j])
		}
	}

}
