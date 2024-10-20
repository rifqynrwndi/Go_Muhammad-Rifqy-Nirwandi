package models

type Word struct {
    ID          string `json:"id"`
    Word        string `json:"word"`
    PartOfSpeech string `json:"partOfSpeech"`
    Definition  string `json:"definition"`
}
