package entity

type VerseResponse struct {
	Group      string   `json:"group"`
	Name       string   `json:"name"`
	Verse      []string `json:"verse"`
	Pagination `json:"pagination"`
}

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}
