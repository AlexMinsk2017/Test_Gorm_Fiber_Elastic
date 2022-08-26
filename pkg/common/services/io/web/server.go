package web

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/web/controllers"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
)

type WebServices struct {
	Orchestrator *orchestrator.Orchestrator
}

func (ws *WebServices) Run() error {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	(&controllers.CustomerController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.GoodController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.OrderController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.TableOrderController{UseCases: ws.Orchestrator}).Init(v1)

	//Hello
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	///////token авторизация

	//////////////
	return app.Listen(":3000")
}
