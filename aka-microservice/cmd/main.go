package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ilyakaznacheev/cleanenv"
	handlers "microservice"
	"microservice/internal/service"

	"log"
	"microservice/internal/storage"
	"os"
)

type Config struct {
	Service service.Config `env-prefix:"SERVICE_"`
}

func main() {
	var err error

	cfg := &Config{}

	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.Fatalln("read env from os err: ", err)
		}
	} else {
		if err := cleanenv.ReadConfig(".env", cfg); err != nil {
			log.Fatalln("read env from .env file err:", err)
		}
	}

	st, err := storage.New()

	if err != nil {
		log.Fatalln("init storage err", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	handlers.InitHandlers(app)

	s, err := service.New(&cfg.Service, st)
	if err != nil {
		log.Fatalln("init service err", err)
	}
	s.Run(app)
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := 500

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(err)
}
