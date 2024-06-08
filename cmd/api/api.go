package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/tanishqtrivedi27/ecom/logger"
	"github.com/tanishqtrivedi27/ecom/service/cart"
	"github.com/tanishqtrivedi27/ecom/service/order"
	"github.com/tanishqtrivedi27/ecom/service/product"
	"github.com/tanishqtrivedi27/ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(router)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: logger.RequestLoggingMiddleWare(v1),
	}

	log.Printf("Server started at %s", s.addr)
	return server.ListenAndServe()
}
