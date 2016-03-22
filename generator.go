package main

import (
	"math/rand"
	"time"
)

var start = [...]string{
	"Б", "В", "Г", "Д", "Ђ", "Ж", "З", "Ј", "К", "Л", "Љ", "М", "Н", "Њ", "П", "Р",
	"С", "Т", "Ћ", "Ф", "Х", "Ц", "Ч", "Џ", "Ш", "Б", "В", "Г", "Д", "Ђ", "Ж", "З",
	"Ј", "К", "Л", "Љ", "М", "Н", "Њ", "П", "Р", "С", "Т", "Ћ", "Ф", "Х", "Ц", "Ч",
	"Џ", "Ш", "Бл", "Бр", "Вл", "Вр", "Гл", "Гр", "Дл", "Др", "Жл", "Зл", "Зр",
	"Кр", "Кл", "Мр", "Мл", "Пј", "Пл", "Пљ", "Пњ", "Пр", "Св", "Сл", "См", "Сп",
	"Ст", "Тл", "Тр", "Фл", "Фљ", "Фњ", "Фр", "Хл", "Хр",
}
var middle = [...]string{
	"а", "е", "и", "о", "у", "р",
}
var end = [...]string{
	"б", "в", "г", "д", "ђ", "ж", "з", "ј", "к", "л", "љ", "м", "н", "њ", "п", "р",
	"с", "т", "ћ", "ф", "х", "ц", "ч", "џ", "ш",
}
var bad = map[string]bool{
	"л": true,
	"р": true,
	"ј": true,
	"љ": true,
	"њ": true,
	"Ђ": true,
	"Ж": true,
	"Ј": true,
	"Л": true,
	"Љ": true,
	"Н": true,
	"Њ": true,
	"Р": true,
	"Ћ": true,
	"Ч": true,
	"Џ": true,
	"Ш": true,
}

const sufix = "о Полумента"

func generate() string {
	rand.Seed(time.Now().UnixNano())

	first := start[rand.Int31n(int32(len(start)-1))]
	second := middle[rand.Int31n(int32(len(middle)-1))]
	for second == "р" {
		if _, ok := bad[first]; ok {
			second = middle[rand.Int31n(int32(len(middle)-1))]
		} else {
			break
		}
	}
	third := end[rand.Int31n(int32(len(end)-1))]
	return first + second + third + sufix
}
