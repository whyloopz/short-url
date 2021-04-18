package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h FiberHandler) SetupFiberRouter(router *fiber.App) {
	validate := validator.New()

	router.Post("/shortUrl", func(ctx *fiber.Ctx) error {
		input := new(CreateShortUrlInput)

		if err := ctx.BodyParser(input); err != nil {
			return err
		}
		if err := validate.Struct(input); err != nil {
			return err
		}

		shortUrl, err := h.sv.CreateShortUrl(input.Url, input.ExpireTime)
		if err != nil {
			return err
		}

		return ctx.JSON(&CreateShortUrlOutput{ShortUrl: shortUrl})
	})
	router.Get("/shortUrl", func(ctx *fiber.Ctx) error {
		originUrl, err := h.sv.GetShortUrl("")
		if err != nil {
			return err
		}

		return ctx.JSON(&GetShortUrlOutput{OriginUrl: originUrl})
	})

	admin := router.Group("/admin")
	admin.Get("/shortUrl", func(ctx *fiber.Ctx) error {
		urls, err := h.sv.AdminGetShortUrls("", "", "")
		if err != nil {
			return err
		}

		return ctx.JSON(&AdminGetShortUrlOutput{Urls: urls})
	})
	admin.Delete("/shortUrl", func(ctx *fiber.Ctx) error {
		if err := h.sv.AdminDeleteShortUrls("", []string{""}); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusOK)
	})
}
