package controller

import (
	"encoding/json"
	"net/http"

	"github.com/danurwijayanto/api-go-ecomm/entity"
	"github.com/danurwijayanto/api-go-ecomm/services"
)

type productController struct{}

var (
	productServices services.ProductServices = services.NewProductServices()
)

// ProductController :
type ProductController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Add(response http.ResponseWriter, request *http.Request)
}

// NewProductController :
func NewProductController() ProductController {
	return &productController{}
}

func (*productController) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	data, err := productServices.GetAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting data}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}

func (*productController) Add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var data entity.Product
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting data}`))
		return
	}
	productServices.Add(&data)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
