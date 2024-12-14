package handlers

import (
	"fmt"
	"log"
	"main-mode/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP обрабатывает входящие HTTP-запросы и перенаправляет их
// на соответствующие методы обработки в зависимости от HTTP-метода.
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Обработка GET-запросов
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Обработка POST-запросов
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// Обработка PUT-запросов
	if r.Method == http.MethodPut {
		// Регулярное выражение для извлечения ID из URL
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		// Проверка корректности извлеченного ID
		if len(g) != 1 {
			p.l.Println("len(g) != 1")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("len(g[0]) != 2")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		// Преобразование ID из строки в целое число
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Логирование полученного ID
		fmt.Printf("Got id: %v\n", id)
		fmt.Printf("g: %v\n", g)

		// Обновление продукта с указанным ID
		p.updateProducts(id, rw, r)
		return
	}

	// Обработка всех остальных не реализованных методов
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// updateProducts обновляет информацию о продукте с заданным ID.
// Он принимает ID продукта, а также объекты для записи HTTP-ответа и запроса.
func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	// Логирование начала обработки PUT-запроса для продуктов
	p.l.Println("Handle PUT Products")

	// Создание нового экземпляра продукта
	prod := &data.Product{}

	// Попытка декодирования JSON-данных из тела запроса в объект продукта
	err := prod.FromJSON(r.Body)
	if err != nil {
		// Возврат ошибки, если данные не могут быть декодированы
		http.Error(rw, "Невозможно отправить данные", http.StatusBadRequest)
		return
	}

	// Логирование обновленного продукта и его ID
	p.l.Printf("Prod updated: %#v  idProd: %#v", prod, id)

	// Попытка обновления продукта в хранилище данных
	var errUpdProd = data.UpdateProduct(id, prod)
	if errUpdProd == data.ErrProductNotFound {
		// Возврат ошибки, если продукт с заданным ID не найден
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if errUpdProd != nil {
		// Возврат ошибки, если произошла другая ошибка при обновлении продукта
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// Получение продуктов
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Невозможно получить данные", http.StatusInternalServerError) // Возврат ошибки
	}
}

// Добавление продуктов
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Невозможно отправить данные", http.StatusBadRequest) // Возврат ошибки
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
