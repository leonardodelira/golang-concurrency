package main

import (
	"go_concurrency/controller"
	"go_concurrency/http"
	"go_concurrency/service"
)

var (
	carDetailsService service.CarDetailsService = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsControler = controller.NewCarDetailsControler(carDetailsService)
	httpRouter http.Router = http.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails);
	httpRouter.SERVE(port)
}