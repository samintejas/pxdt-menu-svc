package api

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"projectx.io/drivethru/service/category"
	"projectx.io/drivethru/service/item"
	"projectx.io/drivethru/service/page"
	"projectx.io/drivethru/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {

	apiServ := APIServer{
		addr: addr,
		db:   db,
	}
	return &apiServ
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	rtpl, err1 := template.ParseFiles("./templates/root.html")
	etpl, err2 := template.ParseFiles("./templates/error.html")

	if err1 != nil || err2 != nil {
		log.Fatalf("could not load templates")
	}

	pageshandler := page.NewHandler(rtpl, etpl)
	pageshandler.RegisterRoutes(router)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	itemStore := item.NewStore(s.db)
	itemHandler := item.NewHandler(itemStore)
	itemHandler.RegisterRoutes(subrouter)

	categoryHandler := category.NewHandler()
	categoryHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
