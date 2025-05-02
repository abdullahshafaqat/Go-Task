package main

import (
	"fmt"
	"os"
	"time"
)

<<<<<<< HEAD
		func wordCounter(filedata string) {
				words := 0

				for i := 0; i < len(filedata); i++ {

					if filedata[i] == ' ' {
						words++
					}
				}
				fmt.Println("Number of words :", words)
=======
func wordCounter(filedata string) {
	words := 0

	for i := 0; i < len(filedata); i++ {

		if filedata[i] == ' ' || filedata[i] == '.' || filedata[i] == ',' {
			words++
		}
	}
	fmt.Println("Number of words :", words)
}

func lineCounter(filedata string) {
	lines := 1
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == '\n' {
			lines++
		}
	}
	fmt.Println("Numbers of lines", lines)
}

func spaceCounter(filedata string) {
	spaces := 0
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == ' ' {
			spaces++
		}
	}
	fmt.Println("Numbers of spaces :", spaces)
}

func puncCounter(filedata string) {
	punc := 0
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == ',' || filedata[i] == '.' {
			punc++
		}
	}
	fmt.Println("Numbers of punctuations", punc)
}

func speciCounter(filedata string) {
	character := 0
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == '@' || filedata[i] == '$' || filedata[i] == '#' || filedata[i] == '%' || filedata[i] == '!' {
			character++
		}
	}
	fmt.Println("Numbers of special characters :", character)
}
func sentenceCounter(filedata string) {
	senten := 0
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == '.' {
			senten++
		}
	}
	fmt.Println("Numbers of sentences :", senten)
}

func vowelCounter(filedata string) {
	vowels := 0
	for i := 0; i < len(filedata); i++ {
		if filedata[i] == 'a' || filedata[i] == 'e' || filedata[i] == 'i' || filedata[i] == 'o' || filedata[i] == 'u' {
			vowels++
		}
	}
	fmt.Println("Numbers of vowels :", vowels)
}

func consonCounnter(filedata string) {

	consonance := 0
	for i := 0; i < len(filedata); i++ {
		switch filedata[i] {
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z',
			'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
			{
				consonance++
>>>>>>> c610f5beaedd1bb6f3f6658ee5a6129026fa92cf
			}

		func lineCounter(filedata string) {
				lines := 1
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == '\n' {
						lines++
					}
				}
				fmt.Println("Numbers of lines", lines)
			}

		func spaceCounter(filedata string) {
				spaces := 0
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == ' ' {
						spaces++
					}
				}
				fmt.Println("Numbers of spaces :", spaces)
			}

		func puncCounter(filedata string) {
				punc := 0
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == ',' || filedata[i] == '.' {
						punc++
					}
				}
				fmt.Println("Numbers of punctuations", punc)
			}

		func speciCounter(filedata string) {
				character := 0
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == '@' || filedata[i] == '$' || filedata[i] == '#' || filedata[i] == '%' || filedata[i] == '!' {
						character++
					}
				}
				fmt.Println("Numbers of special characters :", character)
			}
		func sentenceCounter(filedata string) {
				senten := 0
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == '.' {
						if i+1 < len(filedata) && filedata[i+1] == ' ' {
							senten++
						}
					}
				}
				fmt.Println("Numbers of sentences :", senten)
			}

		func vowelCounter(filedata string) {
				vowels := 0
				for i := 0; i < len(filedata); i++ {
					if filedata[i] == 'a' || filedata[i] == 'e' || filedata[i] == 'i' || filedata[i] == 'o' || filedata[i] == 'u' {
						vowels++
					}
				}
				fmt.Println("Numbers of vowels :", vowels)
			}

		func consonCounnter(filedata string) {

				consonants := 0
				for i := 0; i < len(filedata); i++ {
					switch filedata[i] {
					case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z',
						'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
						{
							consonants++
						}
					}
				}
				fmt.Println("Numbers of consonants :", consonants)
			}


		func paraCounter(filedata string){
					var para = 1
					for i := 0; i < len(filedata); i++ {
						if filedata[i] == '\n' {
							if i+1 < len(filedata) && filedata[i+1] == '\n' {
								para++
							}
						}
					}
					fmt.Println("Numbers of Paragraphs:", para)
				
			}


		func digitCounter(filedata string) {
					digit := 0
					for i := 0; i < len(filedata); i++ {
						if filedata[i] >= '0' && filedata[i] <= '9' {
							digit++
						}
					}
					fmt.Println("Numbers of digits :", digit)
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
					Str := string(filedata)
					wordCounter(Str)
					lineCounter(Str)
					spaceCounter(Str)
					puncCounter(Str)
					speciCounter(Str)
					sentenceCounter(Str)
					vowelCounter(Str)
					consonCounnter(Str)
					digitCounter(Str)
					paraCounter(Str)

			}
