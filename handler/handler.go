package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nqmt/short-url/service"
)

type Handler struct {
	sv service.Service
}

func New(sv service.Service) *Handler {
	return &Handler{sv: sv}
}

func (h Handler) SetupRouter(router *fiber.App) {
	router.Post("/shortUrl", func(ctx *fiber.Ctx) error {
		shortUrl, err := h.sv.CreateShortUrl("", 0)
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
