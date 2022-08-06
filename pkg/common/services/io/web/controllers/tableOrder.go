package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TableOrderController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *TableOrderController) Init(router fiber.Router) {
	router.Get("table_orders/get_by_id", cc.GetByID)
	router.Post("table_orders/create", cc.Create)
}

func (cc *TableOrderController) GetByID(ctx *fiber.Ctx) error {
	id_str := ctx.Query("Id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}
	TableOrder, err := cc.UseCases.TableOrderOrchestrator.GetByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.JSON(TableOrder)
}
func (cc *TableOrderController) Create(ctx *fiber.Ctx) error {
	bodyData := web.TableOrder{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}
	TableOrder, err := cc.UseCases.TableOrderOrchestrator.Create(ctx.Context(), &bodyData)
	if err != nil {
		return err
	}
	return ctx.JSON(TableOrder)
}
