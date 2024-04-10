package models

type Content struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Cast        Cast   `json:"cast"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
	Category    string `json:"category"`
	Trailer     string `json:"trailer"`
}
