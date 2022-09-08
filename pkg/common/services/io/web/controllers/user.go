package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *UserController) Init(router fiber.Router) {
	//router.Get("users/get_by_id", cc.GetByID)
	//router.Post("users/create", cc.Create)
	//router.Post("users/delete", cc.DeleteMark)
	router.Post("users/login", cc.Login)
}

func (cc *UserController) Login(ctx *fiber.Ctx) error {
	bodyData := web.User{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}

	user := bodyData.User
	pass := bodyData.Pass

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
