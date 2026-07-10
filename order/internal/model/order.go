package model

import (
	"time"

	inventoryv1 "github.com/MoMentalochka/HomeWork/shared/pkg/proto/inventory/v1"
)

type Part struct {
	//  Уникальный идентификатор детали
	Uuid string
	//  Название детали
	Name string
	//  	Описание детали
	Description string
	//  Цена за единицу
	Price float64
	//  Количество на складе
	StockQuantity int64
	//  Категория
	Category string
	//  Размеры детали
	Dimensions *Dimensions
	//  Информация о производителе
	Manufacturer *Manufacturer
	//  Теги для быстрого поиска
	Tags []string
	//  Гибкие метаданные
	Metadata map[string]*inventoryv1.Value
	//  Дата создания
	CreatedAt *time.Time
	//  	Дата обновления
	UpdatedAt *time.Time
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

type PartsFilter struct {
	//  Список UUID. Пусто — не фильтруем по UUID
	Uuids []string
	//  Список имён. Пусто — не фильтруем по имени
	Names []string
	//  Список категорий. Пусто — не фильтруем по категории
	Categories []string
	//  Список стран производителей. Пусто — не фильтруем по стране
	ManufacturerCountries []string
	//  Список тегов. Пусто — не фильтруем по тегам
	Tags []string
}

type CreateOrderRequest struct {
	// User ID.
	UserUUID string
	// Parts ID array.
	PartUuids []string
}

type OrderDto struct {
	// Order ID.
	OrderUUID string
	// User ID.
	UserUUID string
	// Parts Id array.
	PartUuids []string
	// Total price.
	TotalPrice float64
	// User ID.
	TransactionUUID string
	PaymentMethod   string
	Status          string
}

func (*OrderDto) getOrderByIdRes() {

}

type GetOrderByIdRes interface {
	getOrderByIdRes()
}
