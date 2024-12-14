package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Моковые данные
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

type Product struct {
	ID          int     `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
type Products []*Product

// Добавление к Products метода FromJSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r) // Сначала кодируем

	return e.Decode(p) // Потом декодируем
}

// Добавление к Products метода ToJSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w) // Сначала кодируем

	return e.Encode(p) // Потом декодируем
}

// UpdateProduct обновляет информацию о продукте с заданным ID.
// Принимает ID продукта и указатель на объект продукта, который содержит обновленные данные.
// Возвращает ошибку, если продукт не найден или произошла другая ошибка.
func UpdateProduct(id int, p *Product) error {
	// Поиск продукта по ID в списке продуктов
	var _, f_ID, err = findProduct(id)

	// Если произошла ошибка при поиске, возвращаем её
	if err != nil {
		return err
	}

	// Устанавливаем ID для обновленного продукта
	p.ID = id
	// Обновляем продукт в списке продуктов
	productList[f_ID] = p

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	// Итерация и поиск продукта по ид
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	// Если не найден продукт, то nil и ошибка
	return nil, -1, ErrProductNotFound
}

// Добавление продукта
func AddProduct(p *Product) {
	p.ID = getNexID()
	productList = append(productList, p)
}

// Получение продуктов
func GetProducts() Products {
	return productList
}

func getNexID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
