package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/remotetodo/handler"
	"github.com/remotetodo/middleware"
	"net/http"
)

type Server struct {
	chi.Router
}

func Route() *Server {

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {

		r.Route("/welcome", func(welcome chi.Router) {
			welcome.Post("/signup", handler.Signup)
			welcome.Post("/login", handler.Login)
			welcome.Post("/reset", handler.ResetPassword)
		})

		r.Route("/todo", func(todo chi.Router) {
			todo.Use(middleware.AuthMiddleware)
			todo.Post("/create", handler.CreateTodo)
			todo.Put("/update", handler.Update)
			todo.Get("/show", handler.Showalltodo)
			todo.Post("/upcoming", handler.Upcoming)
			todo.Get("/expired", handler.Expired)
			todo.Delete("/deleted", handler.Deletetodo)
			todo.Get("/completed", handler.Completed)
		})

	})
	return &Server{router}
}

func (svc *Server) Run() error {
	err := http.ListenAndServe(":8888", svc)
	if err != nil {
		return err
	}
	return nil
}
