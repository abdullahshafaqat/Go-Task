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
		resultchann := make(chan []int)
		go combine.AnalyzeText(Str, resultchann)
		value := <-resultchann
		_, err = db.Exec(`
			INSERT INTO results (
				filename, words, lines, spaces, punctuation, characters,
				sentences, vowels, consonants, digits, paragraphs, combine_time
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			file.Filename,
			value[0],
			value[1],
			value[2],
			value[3],
			value[4],
			value[5],
			value[6],
			value[7],
			value[8],
			value[9],
			timeTaken.String(),
		)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error saving to database: %v", err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Filename":       file.Filename,
			"Words":          value[0],
			"Lines":          value[1],
			"Spaces":         value[2],
			"Punctuation":    value[3],
			"Characters":     value[4],
			"Sentences":      value[5],
			"Vowels":         value[6],
			"Consonants":     value[7],
			"Digits":         value[8],
			"Paragraphs":     value[9],
			"Execution time": timeTaken.String(),
		})
	}
}
