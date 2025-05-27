package middlewares

import (
	"time"

	"github.com/abdullahshafaqat/GOTASKS/combine"
	"github.com/abdullahshafaqat/GOTASKS/models"
)

func TextAnalysis(Str string) (models.Results, time.Duration) {
	startTime := time.Now()
	resultchann := make(chan combine.ResultStruct)
	go combine.AnalyzeText(Str, resultchann)
	value := <-resultchann
	timeTaken := time.Since(startTime)
	result := models.Results{
		Words:      value.Words,
		Lines:      value.Lines,
		Spaces:     value.Spaces,
		Punc:       value.Punc,
		Characters: value.Characters,
		Sentences:  value.Sentences,
		Vowels:     value.Vowels,
		Consonants: value.Consonants,
		Digits:     value.Digits,
		Para:       value.Para,
	}

	return result, timeTaken
}
