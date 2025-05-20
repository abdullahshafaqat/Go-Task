package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/abdullahshafaqat/GOTASKS/combine"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AnalyzeText(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "error reading file: %v", err)
			return
		}
		Data, _ := file.Open()
		defer Data.Close()
		fileData, _ := io.ReadAll(Data)
		Str := string(fileData)
		timeTaken := time.Since(startTime)
		resultchann := make(chan combine.ResultStruct)
		go combine.AnalyzeText(Str, resultchann)
		value := <-resultchann
		_, err = db.Exec(`
			INSERT INTO results (
				filename, words, lines, spaces, punctuation, characters,
				sentences, vowels, consonants, digits, paragraphs, combine_time
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			file.Filename,
			value.Words,
			value.Lines,
			value.Spaces,
			value.Punc,
			value.Characters,
			value.Sentences,
			value.Vowels,
			value.Consonants,
			value.Digits,
			value.Para,
			timeTaken.String(),
		)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error saving to database: %v", err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Filename":       file.Filename,
			"Words":          value.Words,
			"Lines":          value.Lines,
			"Spaces":         value.Spaces,
			"Punctuation":    value.Punc,
			"Characters":     value.Characters,
			"Sentences":      value.Sentences,
			"Vowels":         value.Vowels,
			"Consonants":     value.Consonants,
			"Digits":         value.Digits,
			"Paragraphs":     value.Para,
			"Execution time": timeTaken.String(),
		})
	}
}
