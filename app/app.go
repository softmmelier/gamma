package app

import (
	"fmt"
)

type Config struct {
	Name string
}

type App struct {
	services []Service
	Name     string
}

func (a App) Use(s Service) {
	fmt.Printf("Adding --> %s\n", s.Name())
	s.Run()
}

func (a App) Start() {
	fmt.Println("Starting app\n", a.Name)
}

//New create new app instance
func New(c Config) *App {
	return &App{
		Name: c.Name,
	}
}
