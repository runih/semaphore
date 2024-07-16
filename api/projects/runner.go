package projects

import (
	"net/http"

	"github.com/ansible-semaphore/semaphore/api/helpers"
	"github.com/ansible-semaphore/semaphore/db"
	"github.com/gorilla/context"
)

func RunnerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := context.Get(r, "project").(db.Project)
		runnerID, err := helpers.GetIntParam("runner_id", w, r)
		if err != nil {
			return
		}

		runner, err := helpers.Store(r).GetRunner(project.ID, runnerID)

		if err != nil {
			helpers.WriteError(w, err)
			return
		}

		context.Set(r, "runner", runner)
		next.ServeHTTP(w, r)
	})
}

func GetRunner(w http.ResponseWriter, r *http.Request) {
	if repo := context.Get(r, "runner"); repo != nil {
		helpers.WriteJSON(w, http.StatusOK, repo.(db.Runner))
		return
	}

	project := context.Get(r, "project").(db.Project)

	repos, err := helpers.Store(r).GetRunners(project.ID)

	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, repos)
}
