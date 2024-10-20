package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/google/uuid"
    "word-api/models"
)

type DictionaryAPIResponse struct {
    Meanings []struct {
        PartOfSpeech string `json:"partOfSpeech"`
        Definitions  []struct {
            Definition string   `json:"definition"`
            Synonyms   []string `json:"synonyms"`
            Antonyms   []string `json:"antonyms"`
        } `json:"definitions"`
    } `json:"meanings"`
}

var db *sql.DB

func InitializeDB(database *sql.DB) {
    db = database
}

func GetAllWords(c echo.Context) error {
    rows, err := db.Query("SELECT id, word, partOfSpeech, definition FROM words")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch words"})
    }
    defer rows.Close()

    var words []models.Word
    for rows.Next() {
        var word models.Word
        if err := rows.Scan(&word.ID, &word.Word, &word.PartOfSpeech, &word.Definition); err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan word"})
        }
        words = append(words, word)
    }

    return c.JSON(http.StatusOK, words)
}

func GetWordByID(c echo.Context) error {
    id := c.Param("id")

    var word models.Word
    err := db.QueryRow("SELECT id, word, partOfSpeech, definition FROM words WHERE id = ?", id).Scan(&word.ID, &word.Word, &word.PartOfSpeech, &word.Definition)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Word not found"})
    }

    return c.JSON(http.StatusOK, word)
}

func AddWord(c echo.Context) error {
    var newWord models.Word
    if err := c.Bind(&newWord); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
    }

    response, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + newWord.Word)
    if err != nil || response.StatusCode != 200 {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch word details"})
    }
    defer response.Body.Close()

    var dictionaryResponse []DictionaryAPIResponse
    json.NewDecoder(response.Body).Decode(&dictionaryResponse)

    if len(dictionaryResponse) > 0 && len(dictionaryResponse[0].Meanings) > 0 {
        newWord.PartOfSpeech = dictionaryResponse[0].Meanings[0].PartOfSpeech
        newWord.Definition = dictionaryResponse[0].Meanings[0].Definitions[0].Definition
    }

    newWord.ID = uuid.New().String()

    _, err = db.Exec("INSERT INTO words (id, word, partOfSpeech, definition) VALUES (?, ?, ?, ?)", newWord.ID, newWord.Word, newWord.PartOfSpeech, newWord.Definition)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add word"})
    }

    return c.JSON(http.StatusCreated, newWord)
}

func DeleteWord(c echo.Context) error {
    id := c.Param("id")

    _, err := db.Exec("DELETE FROM words WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Word not found"})
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully"})
}
