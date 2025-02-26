package utils

import (
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func ErrorResponse(status int, err error, c *fiber.Ctx) error {
	return c.Status(status).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func SuccessResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func ParseAndValidate(c *fiber.Ctx, o interface{}) error {
	err := c.BodyParser(o)
	if err != nil {
		return err
	}

	v := validator.New()
	err = v.Struct(o)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return e
			}
		}
	}

	return nil
}
