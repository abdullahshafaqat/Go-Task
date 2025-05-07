package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/abdullahshafaqat/GOTASKS/chunks"
	"github.com/abdullahshafaqat/GOTASKS/combine"
	"github.com/gin-gonic/gin"
)

func main() {
	startTime := time.Now()
	fmt.Println("Reading file")
	router := gin.Default()
	router.POST("/analyzer", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "error reading file: %v", err)
			return
		}
		Data, _ := file.Open()
		defer Data.Close()
		fileData, _ := io.ReadAll(Data)
		Str := string(fileData)

		// var wg sync.WaitGroup
		// wg.Add(10)
		// go separate.WordCounter(Str, &wg)
		// go separate.LineCounter(Str, &wg)
		// go separate.SpaceCounter(Str, &wg)
		// go separate.PuncCounter(Str, &wg)
		// go separate.SpeciCounter(Str, &wg)
		// go separate.SentenceCounter(Str, &wg)
		// go separate.VowelCounter(Str, &wg)
		// go separate.ConsonCounnter(Str, &wg)
		// go separate.DigitCounter(Str, &wg)
		// go separate.ParaCounter(Str, &wg)
		// wg.Wait()
		timeTaken := time.Since(startTime)
		resultchann := make(chan []int)
		go combine.AnalyzeText(Str, resultchann)
		value := <-resultchann
		// c.JSON(200,gin.H{
		//    "Filename": file.Filename,
		//    "Using combine function Words": value[0],
		//    "Using combine function Lines": value[1],
		//    "Using combine function Spaces": value[2],
		//    "Using combine function Punctuation": value[3],
		//    "Using combine function Characters": value[4],
		//    "Using combine function Sentences" : value[5],
		//    "Using combine function Vowels": value[6],
		//    "Using combine function Consonants" : value[7],
		//    "Using combine function Digits": value[8],
		//    "Using combine function Paragraphs" : value[9],
		// })
		// timeTaken := time.Since(startTime)
		// fmt.Println("Time Taken After Using Channels & Goroutines: ", timeTaken)
		// for i := 0; i < 10; i++ {
		// 	fmt.Println(names[i], value[i])
		// }
		timeStart3 := time.Now()

		result := chunks.Chunks(Str)
		timeTaken3 := time.Since(timeStart3)
		c.JSON(200, gin.H{
			"Filename":                           file.Filename,
			"Chunks Result":                             result,
			"Using combine function Words":       value[0],
			"Using combine function Lines":       value[1],
			"Using combine function Spaces":      value[2],
			"Using combine function Punctuation": value[3],
			"Using combine function Characters":  value[4],
			"Using combine function Sentences":   value[5],
			"Using combine function Vowels":      value[6],
			"Using combine function Consonants":  value[7],
			"Using combine function Digits":      value[8],
			"Using combine function Paragraphs":  value[9],
			"Time taken for combine function":    timeTaken.String(),
			"Time Taken for chunks":              timeTaken3.String(),
		})

	})
	router.Run(":8080")

}
