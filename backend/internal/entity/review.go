package entity

type Review struct {
	Id          int     `json:"id" db:"id"`            // id тура
	Author      string  `json:"author" db:"username"`  // автор отзыва
	Description string  `json:"desc" db:"description"` // сам отзыв
	Rating      float32 `json:"rating" db:"rating"`    // рейтинг
}
