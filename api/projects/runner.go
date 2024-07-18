package projects

import (
	"errors"
	"fmt"
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

func UpdateRunner(w http.ResponseWriter, r *http.Request) {
	oldRunner := context.Get(r, "runner").(db.Runner)

	var runner db.Runner

	if !helpers.Bind(w, r, &runner) {
		return
	}

	if runner.ID != oldRunner.ID {
		helpers.WriteErrorStatus(w,
			"Runner ID in body and URL must be the same",
			http.StatusBadRequest)
		return
	}

	if err := helpers.Store(r).UpdateRunner(runner); err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.EventLog(r, helpers.EventLogUpdate, helpers.EventLogItem{
		UserID:      helpers.UserFromContext(r).ID,
		ProjectID:   *oldRunner.ProjectID,
		ObjectType:  db.EventRunner,
		ObjectID:    oldRunner.ID,
		Description: fmt.Sprintf("Runner %d (%s) updated", runner.ID, *runner.Name),
	})
	w.WriteHeader(http.StatusNoContent)
}

func RemoveRunner(w http.ResponseWriter, r *http.Request) {
	runner := context.Get(r, "runner").(db.Runner)
	var err error

	err = helpers.Store(r).DeleteRunner(*runner.ProjectID, runner.ID)
	if errors.Is(err, db.ErrInvalidOperation) {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Runner is in use by on or more task",
			"inUse": true,
		})
		return
	}

	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.EventLog(r, helpers.EventLogDelete, helpers.EventLogItem{
		UserID:      helpers.UserFromContext(r).ID,
		ProjectID:   *runner.ProjectID,
		ObjectType:  db.EventRunner,
		ObjectID:    runner.ID,
		Description: fmt.Sprintf("Runner %d, (%s) deleted", runner.ID, *runner.Name),
	})

	w.WriteHeader(http.StatusNoContent)
}
