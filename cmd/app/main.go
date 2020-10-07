package main
import (
    "fmt"
	"net/http"
	
    "github.com/aaronmack/simple_app/app/router"
	"github.com/aaronmack/simple_app/config"
	lr "github.com/aaronmack/simple_app/util/logger"
)

func main() {
	appConf := config.AppConfig()
	print("conf Debug: $appConf.Debug")
	logger := lr.New(appConf.Debug)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	appRouter := router.New()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	
	logger.Info().Msgf("Starting server %v", address)

    s := &http.Server{
        Addr:         address,
		// Handler:      mux,
		Handler:      appRouter,
        ReadTimeout:  appConf.Server.TimeoutRead,
        WriteTimeout: appConf.Server.TimeoutWrite,
        IdleTimeout:  appConf.Server.TimeoutIdle,
    }
    if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
    }
}
func Greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}