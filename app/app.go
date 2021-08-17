package app

import (
	"fmt"
)

const (
	errRunnerExist = "runner exist. Received: %s"
)

type Config struct {
	Name string
}

type App struct {
	runners []*Runner
	Name     string
}

//New create new app instance
func New(c Config) *App {
	return &App{
		Name: c.Name,
	}
}

func (a *App) Service(runs ...Runner) error {
	for _, run := range runs {
		if a.runnerExist(run) {
			return fmt.Errorf(errRunnerExist, run.Name())
		}

		a.runners = append(a.runners, &run)
	}
	return nil
}

func (a *App) Start() {
	fmt.Println("Starting app\n", a.Name)
}

func (a *App) runnerExist(r Runner) bool {
	for _, run := range a.runners {
		if run == &r {
			return true
		}
	}
	return false
}