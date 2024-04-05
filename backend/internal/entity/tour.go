package entity

import "time"

type Tour struct {
	Id              int        `json:"id"`
	Title           string     `json:"title"`
	Location        string     `json:"location"`
	Category        string     `json:"category"`
	Type            []Tag      `json:"type"`
	Desc            string     `json:"description"`
	NightCount      int        `json:"night_count"`
	Program         []string   `json:"program"`
	Included        []string   `json:"included"`
	NotIncluded     []string   `json:"not_included"`
	DifficultyLevel string     `json:"difficulty_level"`
	ComfortLevel    string     `json:"comfort_level"`
	NearestDate     time.Time  `json:"nearest_date"`
	GroupDates      []TourDate `json:"dates"`
	ImportantInfo   string     `json:"important_info"`
	Media           []Image    `json:"media"`
	Faq             string     `json:"faq"`

	/* embed count of companions */
}

type Tag struct {
	Title string `json:"title"`
}

type Image struct {
	Caption string `json:"caption"`
	URL     string `json:"url"`
}

type TourDate struct {
	Var   string    `json:"Var"`
	Start time.Time `json:"Start"`
	End   time.Time `json:"End"`
}
