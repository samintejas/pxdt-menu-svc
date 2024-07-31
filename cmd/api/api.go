package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"projectx.io/drivethru/service/category"
	"projectx.io/drivethru/service/item"
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
	router.HandleFunc("/", index)
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	itemHandler := item.NewHandler()
	itemHandler.RegisterRoutes(subrouter)

	categoryHandler := category.NewHandler()
	categoryHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}

func index(w http.ResponseWriter, r *http.Request) {
	html := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>ProjectX-Drive Thru Menu</title>
        <style>
            body {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 100vh;
                margin: 0;
                font-family: Arial, sans-serif;
                background-color: #f0f0f0;
            }
            .container {
                text-align: center;
                padding: 20px;
                border-radius: 8px;
                background-color: #fff;
                box-shadow: 0 0 10px rgba(0,0,0,0.1);
            }
            h1 {
                margin: 0;
                font-size: 2.5em;
                color: #333;
            }
            p {
                font-size: 1.2em;
                color: #666;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Project X - Code DT</h1>
            <p>Backend Menu service !</p>
        </div>
    </body>
    </html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
