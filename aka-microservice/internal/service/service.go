package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"sync"
)

var (
	once     sync.Once
	instance *Service
)

type Config struct {
	Port int `env:"PORT"`
}

type Service struct {
	Storage IStorage
	Config  *Config
}

func New(cfg *Config, storage IStorage) (*Service, error) {
	once.Do(func() {
		instance = &Service{Storage: storage, Config: cfg}
	})
	return instance, nil
}

func (s *Service) Run(app *fiber.App) {
	err := app.Listen(fmt.Sprintf(":%v", s.Config.Port))
	if err != nil {
		log.Fatalln(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan
	s.Storage.Close()
	app.Shutdown()
}

func GetInstance() (*Service, error) {
	if instance == nil {
		return nil, errors.New("service is not initialized")
	}
	return instance, nil
}
