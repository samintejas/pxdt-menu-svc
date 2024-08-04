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
    <title>PXDT - Menu service</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        header {
            background-color: #333;
            color: white;
            padding: 10px 0;
            text-align: center;
        }
        header h1 {
            margin: 0;
            font-size: 2.5em;
        }
        .about {
            background-color: white;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        .about h2 {
            margin-top: 0;
        }
        .social-links {
            text-align: center;
            margin-top: 20px;
        }
        .social-links a {
            margin: 0 10px;
            text-decoration: none;
            color: #333;
        }
        .social-links a:hover {
            color: #007BFF;
        }
        footer {
            text-align: center;
            padding: 20px 0;
            background-color: #333;
            color: white;
            position: fixed;
            width: 100%;
            bottom: 0;
        }
    </style>
</head>
<body>
    <header>
        <h1>Project X - Menu service</h1>
    </header>
    <div class="container">
        <div class="about">
            <h2>Available API's</h2>
            <li>User Management</li>
            <li>Item Management</li>
            <li>Category Management</li>
            </div>
        <div class="social-links">
            <a href="https://www.linkedin.com/in/samin-tejas" target="_blank"><i class="fab fa-linkedin"></i></a>
            <a href="https://github.com/samin-tejas" target="_blank"><i class="fab fa-github"></i></a>
            <a href="https://twitter.com/samin-tejas" target="_blank"><i class="fab fa-twitter"></i></a>
        </div>
    </div>
    <footer>
        <p>&copy; 2024 Samin Tejas. All rights reserved.</p>
    </footer>
</body>
</html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
