package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type GoodController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *GoodController) Init(router fiber.Router) {
	router.Get("goods/get_by_id", cc.GetByID)
	router.Post("goods/create", cc.Create)
}

func (cc *GoodController) GetByID(ctx *fiber.Ctx) error {
	id_str := ctx.Query("Id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}
	Good, err := cc.UseCases.GoodOrchestrator.GetByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.JSON(Good)
}
func (cc *GoodController) Create(ctx *fiber.Ctx) error {
	bodyData := web.Good{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}
	Good, err := cc.UseCases.GoodOrchestrator.Create(ctx.Context(), &bodyData)
	if err != nil {
		return err
	}
	return ctx.JSON(Good)
}
