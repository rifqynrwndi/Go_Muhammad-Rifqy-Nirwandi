package controllers

import (
	"net/http"
	"word-api/models"

	"github.com/labstack/echo/v4"
)

func GetWords(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data":    models.Words,
		"message": "all words",
	})
}

func AddWord(c echo.Context) error {
	wordInput := new(struct {
		Word string `json:"word"`
	})

	if err := c.Bind(wordInput); err != nil || wordInput.Word == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to add word data",
		})
	}

	wordData := models.Word{
		Word:        wordInput.Word,
		Length:      len(wordInput.Word),
		NumOfVocals: models.CountVocals(wordInput.Word),
		Palindrome:  models.IsPalindrome(wordInput.Word),
	}

	models.Words = append(models.Words, wordData)

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "word added",
		"data":    wordData,
	})
}
