package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nqmt/goerror"
	"log"
)

var (
	ErrBadRequestValidateInput = goerror.DefineBadRequest("ErrBadRequestValidateInput", "input not correct format")
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func RespError(ctx *fiber.Ctx, err error) error {
	log.Println(err.(*goerror.GoError).Code + ": " + err.(*goerror.GoError).Msg)
	if err.(*goerror.GoError).Cause != "" {
		log.Println(err.(*goerror.GoError).Cause)
	}

	return ctx.Status(err.(*goerror.GoError).Status).JSON(ErrorMessage{
		Message: err.(*goerror.GoError).Msg,
	})
}
