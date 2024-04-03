package entity

type Review struct {
	Id          int     `db:"id"`          //id тура
	Author      string  `db:"username"`    // автор отзыва
	Description string  `db:"description"` // сам отзыв
	Rating      float32 `db:"rating"`      //рейтинг
}
