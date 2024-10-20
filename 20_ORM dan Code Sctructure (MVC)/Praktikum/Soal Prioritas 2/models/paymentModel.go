package models

type Payment struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
}
