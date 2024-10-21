package models

type Package struct {
	ID               uint    `json:"id" gorm:"primaryKey"`
	Name             string  `json:"name"`
	Sender           string  `json:"sender"`
	Receiver         string  `json:"receiver"`
	SenderLocation   string  `json:"sender_location"`
	ReceiverLocation string  `json:"receiver_location"`
	Fee              int     `json:"fee"`
	Weight           float64 `json:"weight"`
}
