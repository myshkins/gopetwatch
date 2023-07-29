package database

type Reading struct {
	Temperature      float32 `json:"temperature"`
	ReadingTimestamp string  `json:"datetime"`
}
