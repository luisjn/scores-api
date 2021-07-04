package entity

type Score struct {
	Id        int64  `json:"id"`
	Points    int64  `json:"points"`
	Player    string `json:"player"`
	CreatedAt int64  `json:"createdAt"`
}
