package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/thelazylemur/hypertask/services/fe/views/components"
	"github.com/thelazylemur/hypertask/services/fe/views/pages"
	"github.com/thelazylemur/hypertask/services/task/client"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	task_client := client.New("localhost:8081")
	defer task_client.Close()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = pages.Home().Render(r.Context(), w)
	})

	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
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

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := task_client.GetTasks()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		components.TaskList(tasks).Render(r.Context(), w)
	})

	http.ListenAndServe(":3001", r)
}
