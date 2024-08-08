package handlers

import (
	"github.com/gofiber/fiber/v2"
	"microservice/internal/service"
	"microservice/internal/storage"
	"net/url"
)

func GetItems(ctx *fiber.Ctx) error {
	inst, err := service.GetInstance()
	if err != nil {
		return err
	}

	parsedURL, err := url.Parse(ctx.OriginalURL())
	if err != nil {
		return err
	}

	var res []storage.Item

	for k, values := range parsedURL.Query() {
		if k == "id" {
			res, err = inst.Storage.GetInfoByIds(values)
			if err != nil {
				return err
			}
		}
	}

	return ctx.Status(200).JSON(res)
}
