package app
import (
    "github.com/aaronmack/simple_app/util/logger"
)
type App struct {
    logger *logger.Logger
}
func New(logger *logger.Logger) *App {
    return &App{logger: logger}
}
func (app *App) Logger() *logger.Logger {
    return app.logger
}