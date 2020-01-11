package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/Royalsspirit/cisco-api/internal/validator"
)

// Server represent the API
type Server struct {
	Port        string
	Environment string
	Logger      *log.Entry
	DB          *sql.DB
	Validator   *validator.Val
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Environment string
	Port        string
	Logger      *log.Entry
	DB          *sql.DB
	Validator   *validator.Val
}

// NewServer create a new server
func NewServer(conf *ServerConfig) *Server {
	return &Server{
		Port:        conf.Port,
		Logger:      conf.Logger,
		Environment: conf.Environment,
		Validator:   conf.Validator,
		DB:          conf.DB,
	}
}

// Run the server
func (s *Server) Run() {
	// Init router
	r := s.Routes()

	addr := fmt.Sprintf("0.0.0.0:%s", s.Port)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		s.Logger.Debug("Server start on ", addr)

		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	s.Logger.Debug("Graceful stop.")

}
