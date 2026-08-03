package model

import (
	"time"

	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

type Category string

const (
	CategoryUnknown  Category = "UNKNOWN"
	CategoryEngine   Category = "ENGINE"
	CategoryFuel     Category = "FUEL"
	CategoryPorthole Category = "PORTHOLE"
	CategoryWing     Category = "WING"
)

type Part struct {
	//  Уникальный идентификатор детали
	Uuid string `bson:"_id,omitempty"`
	//  Название детали
	Name string `bson:"name,omitempty"`
	//  	Описание детали
	Description string `bson:"description"`
	//  Цена за единицу
	Price float64 `bson:"price"`
	//  Количество на складе
	StockQuantity int64 `bson:"stock_quantity"`
	//  Категория
	Category Category `bson:"category"`
	//  Размеры детали
	Dimensions *Dimensions `bson:"dimensions"`
	//  Информация о производителе
	Manufacturer *Manufacturer `bson:"manufacturer"`
	//  Теги для быстрого поиска
	Tags []string `bson:"tags"`
	//  Гибкие метаданные
	Metadata map[string]*inventoryv1.Value `bson:"metadata"`
	//  Дата создания
	CreatedAt *time.Time `bson:"created_at"`
	//  	Дата обновления
	UpdatedAt *time.Time `bson:"updated_at,omitempty"`
}

type PartsFilter struct {
	//  Список UUID. Пусто — не фильтруем по UUID
	Uuids []string
	//  Список имён. Пусто — не фильтруем по имени
	Names []string
	//  Список категорий. Пусто — не фильтруем по категории
	Categories []Category
	//  Список стран производителей. Пусто — не фильтруем по стране
	ManufacturerCountries []string
	//  Список тегов. Пусто — не фильтруем по тегам
	Tags []string
}

type Dimensions struct {
	//  Длина в см
	Length float64
	//  Ширина  в см
	Width float64
	//  о в см
	Height float64
	//  Вес в кг
	Weight float64
}

type Manufacturer struct {
	//  Название
	Name string
	//  Название
	Country string
	//  Название
	Website string
}
