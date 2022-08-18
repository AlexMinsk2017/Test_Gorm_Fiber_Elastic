package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CustomerController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *CustomerController) Init(router fiber.Router) {
	router.Get("customers/get_by_id", cc.GetByID)
	router.Get("customers/update_elastic", cc.LoadElastic)
	router.Post("customers/create", cc.Create)
	router.Post("customers/delete", cc.DeleteMark)
}

func (cc *CustomerController) GetByID(ctx *fiber.Ctx) error {
	id_str := ctx.Query("Id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}
	customer, err := cc.UseCases.CustomerOrchestrator.GetByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.JSON(customer)
}
func (cc *CustomerController) Create(ctx *fiber.Ctx) error {
	bodyData := web.Customer{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		return err
	}
	customer, err := cc.UseCases.CustomerOrchestrator.Create(ctx.Context(), &bodyData)
	if err != nil {
		return err
	}
	return ctx.JSON(customer)
}
func (cc *CustomerController) DeleteMark(ctx *fiber.Ctx) error {
	id_str := ctx.Query("Id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}
	err = cc.UseCases.CustomerOrchestrator.DeleteMark(ctx.Context(), uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (cc *CustomerController) LoadElastic(ctx *fiber.Ctx) error {
	err := cc.UseCases.CustomerOrchestrator.UpdateElastic(ctx.Context())
	if err != nil {
		return err
	}
	return nil
}
