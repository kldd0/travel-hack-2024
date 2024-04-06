package entity

type Review struct {
	TourId        int      `json:"tour_id" db:"tour_id"`               // id тура
	Liked         []string `json:"liked" db:"liked"`                   // рейтинг
	Author        string   `json:"author" db:"username"`               // автор отзыва
	Positive      string   `json:"positive" db:"positive"`             //достоинства
	Negative      string   `json:"negative" db:"negative"`             //недостатки
	LocalResident bool     `json:"local_resident" db:"local_resident"` // местный?
	Type          string   `json:"type" db:"type"`                     //тип отдыха (в одиночку. с детьми...)
	Frequency     string   `json:"frequency" db:"frequency"`           //частота посещения
	Video         string   `json:"video" db:"video"`                   //ссылка на видео
}
