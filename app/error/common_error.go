package error

import (
	"github.com/gofiber/fiber/v2"
)

type BadInputError struct {
	Message string
	Err     error
}

type MissingRequestError struct {
	Message string
	Err     error
}

type UnauthorizedError struct {
	Message string
	Err     error
}

type ResultNotFoundError struct {
	Message string
	Err     error
}

type ApiError struct {
	Message string
	Err     error
}

func (e *BadInputError) Error() string {
	return e.Message
}

func (e *MissingRequestError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func (e *ResultNotFoundError) Error() string {
	return e.Message
}

func (e *ApiError) Error() string {
	if e.Message == "" {
		return "Api Error"
	}
	return e.Message
}

func CustomError(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *UnauthorizedError:
		return ctx.Status(401).JSON(fiber.Map{
			"error":   "UnauthorizedError",
			"message": e.Message,
		})
	case *MissingRequestError:
		return ctx.Status(400).JSON(fiber.Map{
			"error":   "MissingRequestError",
			"message": e.Message,
		})
	case *BadInputError:
		return ctx.Status(401).JSON(fiber.Map{
			"error":   "BadInputError",
			"message": e.Message,
		})
	case *ResultNotFoundError:
		return ctx.Status(400).JSON(fiber.Map{
			"error":   "ResultNotFound",
			"message": e.Message,
		})
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "InternalServerError",
			"message": e.Error(),
		})
	}
}
