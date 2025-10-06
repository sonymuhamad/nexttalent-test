package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/sonymuhamad/nexttalent-test/config"
	"github.com/sonymuhamad/nexttalent-test/providerwire"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.LoadEnvConfig()

	server := providerwire.ProvideApiServer(cfg)

	log.Printf("Using timezone %s\n", time.Local.String())
	log.Printf("Starting api server on Port %d\n", cfg.APIPort)

	go func() {
		ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.APIPort))

		if err := server.Serve(ln); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Fatal(err.Error())
			}
		}
	}()

	<-sigs
	// terminating, graceful shutdown
	log.Println("Shutting Down HTTP Server")

	server.SetKeepAlivesEnabled(false)

	_ = server.Shutdown(context.Background())

	log.Println("Shutdown Gracefully")

}
