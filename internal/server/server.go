package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/adityatresnobudi/restapi/config"
	"github.com/adityatresnobudi/restapi/docs"
	"github.com/adityatresnobudi/restapi/internal/domains/transaction/handler"
	"github.com/adityatresnobudi/restapi/internal/domains/transaction/service"
	"github.com/adityatresnobudi/restapi/internal/repositories/transaction_repo/transaction_pg"
	"github.com/adityatresnobudi/restapi/pkg/postgres"
)

type server struct {
	cfg config.Config
	mux *http.ServeMux
}

func NewServer(cfg config.Config) *server {
	return &server{
		cfg: cfg,
		mux: http.NewServeMux(),
	}
}

func (s *server) Run() {
	db, err := postgres.NewDB(
		s.cfg.Postgres.Host,
		s.cfg.Postgres.Port,
		s.cfg.Postgres.User,
		s.cfg.Postgres.Password,
		s.cfg.Postgres.DBName,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = postgres.InitializeTable(db); err != nil {
		switch {
		case strings.Contains(err.Error(), "already exists"):
			fmt.Printf("Table %s already exists. Proceeding to connect...\n", strings.ToUpper(strings.Fields(err.Error())[2][1:len(strings.Fields(err.Error())[2])-1]))
		default:
			if err = db.Close(); err != nil {
				log.Printf("db graceful shutdown: %s\n", err.Error())
			} else {
				fmt.Printf("db graceful shutdown succeeded\n")
			}
			return
		}
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", s.cfg.Http.Host, s.cfg.Http.Port)

	transactionRepo := transaction_pg.NewRepo(db)

	transactionService := service.NewTransactionService(transactionRepo)

	transactionHandler := handler.NewTransactionHandler(s.mux, ctx, transactionService)

	transactionHandler.MapRoutes()

	go func() {
		log.Printf("Listening on PORT: %s\n", s.cfg.Http.Port)
		if err := s.runHTTPServer(); err != nil {
			log.Printf("s.runHTTPServer: %s\n", err.Error())
		}

	}()

	oscall := <-ch

	if err = db.Close(); err != nil {
		log.Printf("db graceful shutdown: %s\n", err.Error())
	} else {
		fmt.Printf("db graceful shutdown succeeded\n")
	}

	fmt.Printf("system call: %+v\n", oscall)
}
