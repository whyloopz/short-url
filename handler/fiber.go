package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nqmt/short-url/service"
)

func (h FiberHandler) SetupFiberRouter(app *fiber.App) {
	app.Post("/", func(ctx *fiber.Ctx) error {
		input := new(service.CreateShortUrlInput)

		if err := ctx.BodyParser(input); err != nil {
			return RespError(ctx, ErrBadRequestValidateInput.WithCause(err))
		}

		shortUrl, err := h.sv.CreateShortUrl(input)
		if err != nil {
			return RespError(ctx, err)
		}
		shortUrl.SetHostName(ctx.Protocol() + "://" + ctx.Hostname())

		return ctx.JSON(shortUrl)
	})
	app.Get("/:shortCode", func(ctx *fiber.Ctx) error {
		originUrl, err := h.sv.GetOriginUrl(ctx.Params("shortCode"))
		if err != nil {
			return RespError(ctx, err)
		}

		return ctx.Redirect(originUrl, fiber.StatusFound)
	})

	admin := app.Group("/admin")
	admin.Get("/shortUrl", func(ctx *fiber.Ctx) error {
		urls, err := h.sv.AdminGetShortUrls("", "", "")
		if err != nil {
			return err
		}

		return ctx.JSON(&service.AdminGetShortUrlOutput{Urls: urls})
	})
	admin.Delete("/shortUrl", func(ctx *fiber.Ctx) error {
		if err := h.sv.AdminDeleteShortUrls("", []string{""}); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusOK)
	})
}
