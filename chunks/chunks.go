package chunks

import (

	"github.com/abdullahshafaqat/GOTASKS/combine"
)
type names struct {
      chunknumber int
	  Words int
	  Lines int
	  Spaces int 
	  Punctuation int
	  Specialcharacters int
	  Sentences int
	  Vowels int
	  Consonants int
	  Digits int
	  Paragraphs int

}

func Chunks(fileData string) []names{
	chunkSize := len(fileData) / 4
	chunk := []string{
		fileData[:chunkSize],
		fileData[chunkSize : chunkSize*2],
		fileData[chunkSize*2 : chunkSize*3],
		fileData[chunkSize*3:],
	}
 var chunkresult [] names
	for i := 0; i < 4; i++ {
		resultchann := make(chan []int)
		go combine.AnalyzeText(chunk[i], resultchann)
		value := <-resultchann
        chunkresult = append(chunkresult, names{
chunknumber: i+1,
Words: value[0],
Lines: value[1],
Spaces: value[2],
Punctuation: value[3],
Specialcharacters: value[4],
Sentences: value[5],
Vowels: value[6],
Consonants: value[7],
Digits: value[8],
Paragraphs: value[9],

		})
		
		
	}
return chunkresult
}
