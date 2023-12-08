package main

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/thelazylemur/hypertask/services/fe/views/components"
	"github.com/thelazylemur/hypertask/services/fe/views/pages"
	"github.com/thelazylemur/hypertask/services/task/client"
)

func pageHandler(p templ.Component) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = pages.Base(p).Render(r.Context(), w)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	task_client := client.New("localhost:8081")
	defer task_client.Close()

	r.Get("/", pageHandler(pages.Home()))

	r.Post("/hx/tasks", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		description := r.FormValue("description")
		weight := r.FormValue("weight")

		weightAsInt, err := strconv.Atoi(weight)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t, err := task_client.CreateTask(name, description, int32(weightAsInt))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		components.Task(*t).Render(r.Context(), w)
	})

	r.Get("/hx/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := task_client.GetTasks()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		components.TaskList(tasks).Render(r.Context(), w)
	})

	slog.Info("Starting server on :3001")
	http.ListenAndServe(":3001", r)
}
