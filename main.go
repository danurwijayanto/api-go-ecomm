package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
)

var (
	productController controller.ProductController = controller.NewProductController()
	httpRouter        router.Router                = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running")
	})

	httpRouter.POST("/add", productController.Add)
	httpRouter.GET("/get-all-product", productController.GetAll)

	httpRouter.SERVE(port)
}
