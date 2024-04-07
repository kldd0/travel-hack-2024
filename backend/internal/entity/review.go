package entity

type Review struct {
	Id              int      `json:"id" db:"id"`
	TourId          int      `json:"tour_id" db:"tour_id"`               // id тура
	Liked           []string `json:"liked" db:"liked"`                   // рейтинг
	Username        string   `json:"username" db:"username"`             // автор отзыва
	PositiveComment string   `json:"positive_comment" db:"positive"`     // достоинства
	NegativeComment string   `json:"negative_comment" db:"negative"`     // недостатки
	LocalResident   bool     `json:"local_resident" db:"local_resident"` // местный?
	Type            string   `json:"type" db:"type"`                     // тип отдыха (в одиночку. с детьми...)
	Frequency       string   `json:"frequency" db:"frequency"`           // частота посещения
	Video           string   `json:"video" db:"video"`                   // ссылка на видео
}

type DTOReview struct {
	Liked           []string  `json:"liked"`
	Username        string    `json:"username"`
	PositiveComment string    `json:"positive_comment"`
	NegativeComment string    `json:"negative_comment"`
	LocalResident   bool      `json:"local_resident"`
	Type            string    `json:"type"`
	Frequency       string    `json:"frequency"`
	Video           MediaType `json:"video"`
}

func (r Review) ToDTOModel() DTOReview {
	return DTOReview{
		Liked:           r.Liked,
		Username:        r.Username,
		PositiveComment: r.PositiveComment,
		NegativeComment: r.NegativeComment,
		LocalResident:   r.LocalResident,
		Type:            r.Type,
		Frequency:       r.Frequency,
		Video: MediaType{
			"video", r.Video,
		},
	}
}
