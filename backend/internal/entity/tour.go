package entity

type Tour struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"username"`
}
