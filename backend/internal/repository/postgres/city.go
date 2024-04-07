package postgres

import (
	"context"
	"strings"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

var (
	cities = []entity.City{
		{Id: 1, Name: "Владимир"},
		{Id: 2, Name: "Кемерово"},
		{Id: 3, Name: "Хабаровск"},
		{Id: 4, Name: "Курск"},
		{Id: 5, Name: "Волжский"},
		{Id: 6, Name: "Орёл"},
		{Id: 7, Name: "Мурманск"},
		{Id: 8, Name: "Ульяновск"},
		{Id: 9, Name: "Пермь"},
		{Id: 10, Name: "Домодедово"},
		{Id: 11, Name: "Сочи"},
		{Id: 12, Name: "Петрозаводск"},
		{Id: 13, Name: "Северодвинск"},
		{Id: 14, Name: "Дзержинск"},
		{Id: 15, Name: "Казань"},
		{Id: 16, Name: "Волжский"},
		{Id: 17, Name: "Воронеж"},
		{Id: 18, Name: "Абакан"},
		{Id: 19, Name: "Раменское"},
		{Id: 20, Name: "Иркутск"},
		{Id: 21, Name: "Домодедово"},
		{Id: 22, Name: "Пермь"},
		{Id: 23, Name: "Пушкино"},
		{Id: 24, Name: "Кемерово"},
		{Id: 25, Name: "Липецк"},
		{Id: 26, Name: "Энгельс"},
		{Id: 27, Name: "Брянск"},
		{Id: 28, Name: "Стерлитамак"},
		{Id: 29, Name: "Оренбург"},
		{Id: 30, Name: "Калуга"},
		{Id: 31, Name: "Воронеж"},
		{Id: 32, Name: "Владикавказ"},
		{Id: 33, Name: "Уфа"},
		{Id: 34, Name: "Абакан"},
		{Id: 35, Name: "Воронеж"},
		{Id: 36, Name: "Волгоград"},
		{Id: 37, Name: "Каменск-Уральский"},
		{Id: 38, Name: "Рыбинск"},
		{Id: 39, Name: "Великий Новгород"},
		{Id: 40, Name: "Архангельск"},
		{Id: 41, Name: "Серпухов"},
		{Id: 42, Name: "Иркутск"},
		{Id: 43, Name: "Солнечногорск"},
		{Id: 44, Name: "Домодедово"},
		{Id: 45, Name: "Петрозаводск"},
		{Id: 46, Name: "Уфа"},
		{Id: 47, Name: "Красноярск"},
		{Id: 48, Name: "Абакан"},
		{Id: 49, Name: "Ульяновск"},
		{Id: 50, Name: "Королёв"},
		{Id: 51, Name: "Курган"},
		{Id: 52, Name: "Рыбинск"},
		{Id: 53, Name: "Сыктывкар"},
		{Id: 54, Name: "Северодвинск"},
		{Id: 55, Name: "Астрахань"},
		{Id: 56, Name: "Владикавказ"},
		{Id: 57, Name: "Санкт-Петербург"},
		{Id: 58, Name: "Химки"},
		{Id: 59, Name: "Краснодар"},
		{Id: 60, Name: "Ставрополь"},
		{Id: 61, Name: "Химки"},
		{Id: 62, Name: "Рязань"},
		{Id: 63, Name: "Краснодар"},
		{Id: 64, Name: "Москва"},
		{Id: 65, Name: "Смоленск"},
		{Id: 66, Name: "Северодвинск"},
		{Id: 67, Name: "Бийск"},
		{Id: 68, Name: "Самара"},
		{Id: 69, Name: "Благовещенск"},
		{Id: 70, Name: "Солнечногорск"},
		{Id: 71, Name: "Пенза"},
		{Id: 72, Name: "Иркутск"},
		{Id: 73, Name: "Краснодар"},
		{Id: 74, Name: "Салават"},
		{Id: 75, Name: "Новосибирск"},
		{Id: 76, Name: "Новокузнецк"},
		{Id: 77, Name: "Махачкала"},
		{Id: 78, Name: "Калуга"},
		{Id: 79, Name: "Братск"},
		{Id: 80, Name: "Стерлитамак"},
		{Id: 81, Name: "Астрахань"},
		{Id: 82, Name: "Орёл"},
		{Id: 83, Name: "Петрозаводск"},
		{Id: 84, Name: "Курган"},
		{Id: 85, Name: "Ставрополь"},
		{Id: 86, Name: "Белгород"},
		{Id: 87, Name: "Барнаул"},
		{Id: 88, Name: "Ярославль"},
		{Id: 89, Name: "Уфа"},
		{Id: 90, Name: "Каменск-Уральский"},
		{Id: 91, Name: "Рыбинск"},
		{Id: 92, Name: "Пенза"},
		{Id: 93, Name: "Ижевск"},
		{Id: 94, Name: "Троицк"},
		{Id: 95, Name: "Абакан"},
		{Id: 96, Name: "Мурманск"},
		{Id: 97, Name: "Архангельск"},
		{Id: 98, Name: "Чебоксары"},
		{Id: 99, Name: "Волгодонск"},
		{Id: 100, Name: "Мурманск"},
	}
)

type CityRepository struct {
	*postgres.Postgres
}

func NewCityRepository(pg *postgres.Postgres) *CityRepository {
	return &CityRepository{pg}
}

func (r *CityRepository) GetMany(ctx context.Context, prefix string, limit int) ([]entity.City, error) {
	var result []entity.City
	for _, city := range cities {
		if strings.HasPrefix(city.Name, prefix) {
			result = append(result, city)
		}

		if len(result) == limit {
			break
		}
	}

	return result, nil
}
