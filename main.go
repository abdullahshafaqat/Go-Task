package main

import (
	"fmt"
	"os"
	"time"
)

type TextState struct {
	Name  string
	Value int
}

func AnalyzeText(filedata string, resultchann chan<- TextState) {
	go func() {
		para := 1
		words := 0
		spaces := 0
		lines := 1
		senten := 0
		punc := 0
		character := 0
		digit := 0
		vowels := 0
		consonants := 0

		for i := 0; i < len(filedata); i++ {
			ch := filedata[i]
			switch ch {
			case '\n':
				lines++
				if i+1 < len(filedata) && filedata[i+1] == '\n' {
					para++
				}
			case ' ':
				words++
				spaces++
			case '.':
				punc++
				if i+1 < len(filedata) && filedata[i+1] == ' ' {
					senten++
				}

			case ',', ';', '"', '/', ':', '\\', '?', '`':
				punc++
			case '@', '$', '#', '%', '!':
				character++
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				{
					digit++
				}
			case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
				'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
				{
					consonants++
				}
			}
		}

		resultchann <- TextState{"Number of paragraphs", para}
		resultchann <- TextState{"Number of words", words}
		resultchann <- TextState{"Number of spaces", spaces}
		resultchann <- TextState{"Number of lines", lines}
		resultchann <- TextState{"Number of sentences", senten}
		resultchann <- TextState{"Number of punctuations", punc}
		resultchann <- TextState{"Number of special characters", character}
		resultchann <- TextState{"Number of digits", digit}
		resultchann <- TextState{"Number of vowels", vowels}
		resultchann <- TextState{"Number of consonants", consonants}

		close(resultchann)
	}()
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(start))
	}()

	fmt.Println("Reading file")
	filename := "Test.txt"
	filedata, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	resultchann := make(chan TextState, 10)
	AnalyzeText(string(filedata), resultchann)


	for res := range resultchann {
		fmt.Printf("%s: %d\n", res.Name, res.Value)
	}

		
}
