package entity

import (
	"strings"
	"time"
)

type Tour struct {
	Id                int         `db:"id"`
	Title             string      `db:"title"`
	Location          string      `db:"location"`
	Category          string      `db:"category"`
	Tags              []string    `db:"tags"`
	Desc              string      `db:"desc"`
	NightsCount       int         `db:"nights_count"`
	Program           []string    `db:"program"`
	Included          []string    `db:"included"`
	NotIncluded       []string    `db:"not_included"`
	DifficultyLevel   string      `db:"difficulty_level"`
	ComfortLevel      string      `db:"comfort_level"`
	Dates             []time.Time `db:"tour_dates"`
	ImportantInfo     string      `db:"important_info"`
	HeadMedia         []string    `db:"head_media"`
	ProgramMedia      []string    `db:"program_media"`
	AccomodationMedia []string    `db:"acc_media"`
	Faq               string      `db:"faq"`
	Rating            int         `db:"rating"`
	IsLiked           bool        `db:"is_liked"`
	MapSrc            string      `db:"map"`
}

// выплевывает мини инфу о мини карточке
type SimplifiedTourView struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Location    string      `json:"location"`
	Category    string      `json:"category"`
	NightsCount int         `json:"nights_count"`
	Tags        []string    `json:"tags"`
	Rating      int         `json:"rating"`
	IsLiked     bool        `json:"is_liked"`
	Media       []MediaType `json:"media"`
}

type DTOTour struct {
	Id              int          `json:"id"`
	Title           string       `json:"title"`
	Location        string       `json:"location"`
	Category        string       `json:"category"`
	Tags            []string     `json:"tags"`
	Desc            string       `json:"description"`
	NightsCount     int          `json:"nights_count"`
	Program         []string     `json:"program"`
	Included        []string     `json:"included"`
	NotIncluded     []string     `json:"not_included"`
	DifficultyLevel string       `json:"difficulty_level"`
	ComfortLevel    string       `json:"comfort_level"`
	TourDates       []TourDate   `json:"tour_dates"`
	ImportantInfo   string       `json:"important_info"`
	Media           MediaSectors `json:"media"`
	Faq             string       `json:"faq"`
	Rating          int          `json:"rating"`
	IsLiked         bool         `json:"is_liked,omitempty"`
	MapSrc          string       `json:"map" db:"map"`
	Reviews         []DTOReview  `json:"reviews"`

	/* liked by user */
	/* embed count of companions and their age group */
}

type MediaSectors struct {
	Head         []MediaType `json:"head"`
	Program      []MediaType `json:"program"`
	Accomodation []MediaType `json:"acc"`
}

type MediaType struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

type TourDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (t Tour) ToDTOModel() DTOTour {
	var dates []TourDate
	for i := 0; i < len(t.Dates)-1; i += 2 {
		dates = append(dates, TourDate{t.Dates[i].Local().Format(time.DateOnly), t.Dates[i+1].Local().Format(time.DateOnly)})
	}

	mediaSectors := MediaSectors{
		Head:         FromSrcToMediaTypes(t.HeadMedia),
		Program:      FromSrcToMediaTypes(t.ProgramMedia),
		Accomodation: FromSrcToMediaTypes(t.AccomodationMedia),
	}

	return DTOTour{
		Id:              t.Id,
		Title:           t.Title,
		Location:        t.Location,
		Category:        t.Category,
		Tags:            t.Tags,
		Desc:            t.Desc,
		NightsCount:     t.NightsCount,
		Program:         t.Program,
		Included:        t.Included,
		NotIncluded:     t.NotIncluded,
		DifficultyLevel: t.DifficultyLevel,
		ComfortLevel:    t.ComfortLevel,
		TourDates:       dates,
		ImportantInfo:   t.ImportantInfo,
		Media:           mediaSectors,
		Faq:             t.Faq,
		Rating:          t.Rating,
		IsLiked:         t.IsLiked,
		MapSrc:          t.MapSrc,
	}
}

func (t Tour) ToSimplifiedTour() SimplifiedTourView {
	media := FromSrcToMediaTypes(t.HeadMedia)

	return SimplifiedTourView{
		Id:          t.Id,
		Title:       t.Title,
		Location:    t.Location,
		Category:    t.Category,
		NightsCount: t.NightsCount,
		Tags:        t.Tags,
		Rating:      t.Rating,
		Media:       media,
		IsLiked:     t.IsLiked,
	}
}

func (t DTOTour) ToSimplifiedTour() SimplifiedTourView {
	return SimplifiedTourView{
		Id:          t.Id,
		Title:       t.Title,
		Location:    t.Location,
		Category:    t.Category,
		NightsCount: t.NightsCount,
		Tags:        t.Tags,
		Rating:      t.Rating,
		Media:       t.Media.Head,
		IsLiked:     t.IsLiked,
	}
}

func FromSrcToMediaTypes(src []string) []MediaType {
	var media []MediaType
	for _, v := range src {
		strSlice := strings.Split(v, ".")
		ext := strSlice[len(strSlice)-1]
		if ext == "jpeg" {
			media = append(media, MediaType{"image", v})
		} else {
			media = append(media, MediaType{"video", v})
		}
	}

	return media
}
