package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/pdexrefapi/models"
	"github.com/nikola43/pdexrefapi/services"
	"github.com/nikola43/pdexrefapi/utils"
)

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

func GetOrCreate(ctx *fiber.Ctx) error {
	req := new(models.CreateUserRequest)

	err := utils.ParseAndValidate(ctx, req)
	if err != nil {
		fmt.Println("error parsing", err)
		return utils.ErrorResponse(fiber.StatusBadRequest, err, ctx)
	}

	isReferrerAddressValid := utils.IsValidAddress(req.ReferrerAddress)
	if !isReferrerAddressValid {
		return utils.ErrorResponse(fiber.StatusBadRequest, fmt.Errorf("invalid referrer address"), ctx)
	}

	isReferredAddressValid := utils.IsValidAddress(req.ReferredAddress)
	if !isReferredAddressValid {
		return utils.ErrorResponse(fiber.StatusBadRequest, fmt.Errorf("invalid referred address"), ctx)
	}

	tx, err := services.GetOrCreate(req)
	if err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, err, ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(tx)
}
