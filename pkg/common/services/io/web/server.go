package web

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/web/controllers"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type WebServices struct {
	Orchestrator *orchestrator.Orchestrator
}

type LoginPassword struct {
	User string
	Pass string
}

func (ws *WebServices) Run() error {

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	///////token авторизация
	(&controllers.UserController{UseCases: ws.Orchestrator}).Init(v1)

	app.Post("/login", login)
	// Unauthenticated route
	app.Get("/", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// Restricted Routes
	app.Get("/restricted", restricted)

	//////////////

	(&controllers.CustomerController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.GoodController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.OrderController{UseCases: ws.Orchestrator}).Init(v1)
	(&controllers.TableOrderController{UseCases: ws.Orchestrator}).Init(v1)

	//Hello
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app.Listen(":3000")
}

func login(ctx *fiber.Ctx) error {
	//user := ctx.FormValue("User")
	//pass := ctx.FormValue("Pass")

	//user := ctx.Params("user", "none")
	//pass := ctx.Params("pass", "none")

	bodyData := LoginPassword{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}
	user := bodyData.User
	pass := bodyData.Pass

	// Throws Unauthorized error
	if user != "alex" || pass != "hryb" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "Alex Hryb",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
