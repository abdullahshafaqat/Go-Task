package separate

import (
	"fmt"
	"sync"
)

func WordCounter(fileData string, wg *sync.WaitGroup) {

	defer wg.Done()
	words := 0

	for i := 0; i < len(fileData); i++ {

		if fileData[i] == ' ' {
			words++
		}
	}
	fmt.Println("Words ", words)
}

func LineCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	lines := 1
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '\n' {
			lines++
		}
	}
	fmt.Println("lines ", lines)
}

func SpaceCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	spaces := 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == ' ' {
			spaces++
		}
	}
	fmt.Println("Spaces ", spaces)
}

func PuncCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	punc := 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case ',', ';', '"', '/', ':', '.', '\'', '?', '`':
			{
				punc++
			}
		}
	}
	fmt.Println("Punctuations ", punc)
}

func SpeciCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	character := 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case '!', '@', '#', '$', '%', '^', '&', '*',
			'(', ')', '[', ']', '{', '}', '+', '=', '-',
			'_', '\\', '|', '<', '>', '?', '~', '`':
			{
				character++
			}
		}
	}
	fmt.Println("Characters ", character)
}
func SentenceCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	senten := 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '.' || fileData[i] == ','  ||fileData[i] == ';'|| fileData[i] == '"' || fileData[i] == '/' || 
		fileData[i] == ':' || fileData[i] == '\\'  || fileData[i] == '?'  || fileData[i] == '`' {
			if i+1 < len(fileData) && fileData[i+1] == ' ' {
				senten++
			}
		}
	}
	fmt.Println("Sentences ", senten)
}

func VowelCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	vowels := 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == 'a' || fileData[i] == 'e' || fileData[i] == 'i' || fileData[i] == 'o' || fileData[i] == 'u' ||
			fileData[i] == 'A' || fileData[i] == 'E' || fileData[i] == 'I' || fileData[i] == 'U' || fileData[i] == 'O' {
			vowels++
		}
	}
	fmt.Println("Vowels ", vowels)
}

func ConsonCounnter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	consonants := 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z',
			'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
			{
				consonants++
			}
		}
	}
	fmt.Println("Consonants ", consonants)
}

func ParaCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	var para = 1
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '\n' {
			if i+1 < len(fileData) && fileData[i+1] == '\n' {
				para++
			}
		}
	}
	fmt.Println("Paragraphs ", para)

}

func DigitCounter(fileData string, wg *sync.WaitGroup) {
	defer wg.Done()
	digit := 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] >= '0' && fileData[i] <= '9' {
			digit++
		}
	}
	fmt.Println("Digits ", digit)
}


