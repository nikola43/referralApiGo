package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/pdexrefapi/models"
	"github.com/nikola43/pdexrefapi/services"
	"github.com/nikola43/pdexrefapi/utils"
)

func CreateUser(ctx *fiber.Ctx) error {
	req := new(models.CreateUserRequest)

	err := utils.ParseAndValidate(ctx, req)
	if err != nil {
		fmt.Println("error parsin", err)
		return utils.ErrorResponse(fiber.StatusBadRequest, err, ctx)
	}

	isValidAddress := utils.IsValidAddress(req.Address)
	if !isValidAddress {
		return utils.ErrorResponse(fiber.StatusBadRequest, fmt.Errorf("invalid address"), ctx)
	}

	tx, err := services.CreateUser(req)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, err, ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(tx)
}

func GetUser(ctx *fiber.Ctx) error {
	address := ctx.Query("address")
	if address == "" {
		return utils.ErrorResponse(fiber.StatusBadRequest, fmt.Errorf("address is required"), ctx)
	}

	tx, err := services.GetUser(address)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, err, ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(tx)
}

func AddReferral(ctx *fiber.Ctx) error {
	req := new(models.AddReferralRequest)

	err := utils.ParseAndValidate(ctx, req)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusBadRequest, err, ctx)
	}

	tx, err := services.AddReferral(req)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusBadRequest, err, ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(tx)
}

func GetUserWithReferrals(ctx *fiber.Ctx) error {
	address := ctx.Query("address")
	if address == "" {
		return utils.ErrorResponse(fiber.StatusBadRequest, fmt.Errorf("address is required"), ctx)
	}

	tx, err := services.GetUserWithReferrals(address)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, err, ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(tx)
}
