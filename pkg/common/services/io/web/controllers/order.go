package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type OrderController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *OrderController) Init(router fiber.Router) {
	router.Get("orders/get_by_id", cc.GetByID)
	router.Post("orders/create", cc.Create)
}

func (cc *OrderController) GetByID(ctx *fiber.Ctx) error {
	id_str := ctx.Query("Id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}
	Order, err := cc.UseCases.OrderOrchestrator.GetByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.JSON(Order)
}
func (cc *OrderController) Create(ctx *fiber.Ctx) error {
	bodyData := web.Order{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}
	Order, err := cc.UseCases.OrderOrchestrator.Create(ctx.Context(), &bodyData)
	if err != nil {
		return err
	}
	return ctx.JSON(Order)
}
