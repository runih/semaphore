package sql

import (
	"encoding/base64"

	"github.com/Masterminds/squirrel"
	"github.com/ansible-semaphore/semaphore/db"
	"github.com/gorilla/securecookie"
)

func (d *SqlDb) GetAllRunners() (runners []db.Runner, err error) {
	query, args, err := squirrel.Select("r.*, p.name as project_name, i.name as inventory_name").
		From("runner as r").
		LeftJoin("project as p ON p.id = r.project_id").
		LeftJoin("project__inventory as i ON i.id = r.inventory_id").
		OrderBy("r.id").
		ToSql()

	if err != nil {
		return
	}

	_, err = d.selectAll(&runners, query, args...)
	return
}

func (d *SqlDb) GetRunnerRefs(projectID int, runnerID int) (db.ObjectReferrers, error) {
	return d.getObjectRefs(projectID, db.GlobalRunnerProps, runnerID)
}

func (d *SqlDb) GetRunner(projectID int, runnerID int) (runner db.Runner, err error) {
	query, args, err := squirrel.Select("r.*, p.name as project_name, i.name as inventory_name").
		From("runner as r").
		LeftJoin("project as p ON p.id = r.project_id").
		LeftJoin("project__inventory as i ON i.id = r.inventory_id").
		Where("r.id = ?", runnerID).
		OrderBy("r.id").
		ToSql()

	err = d.selectOne(&runner, query, args...)

	if err != nil {
		return
	}

	return

}

func (d *SqlDb) GetRunners(projectID int, params db.RetrieveQueryParams) ([]db.Runner, error) {
	var runners []db.Runner
	err := d.getObjects(projectID, db.GlobalRunnerProps, params, nil, &runners)
	return runners, err
}

func (d *SqlDb) DeleteRunner(projectID int, runnerID int) (err error) {
	return d.deleteObject(projectID, db.GlobalRunnerProps, runnerID)
}

func (d *SqlDb) GetGlobalRunner(runnerID int) (runner db.Runner, err error) {
	err = d.getObject(0, db.GlobalRunnerProps, runnerID, &runner)
	return
}

func (d *SqlDb) GetGlobalRunners() (runners []db.Runner, err error) {
	err = d.getObjects(0, db.GlobalRunnerProps, db.RetrieveQueryParams{}, nil, &runners)
	return
}

func (d *SqlDb) DeleteGlobalRunner(runnerID int) (err error) {
	err = d.deleteObject(0, db.GlobalRunnerProps, runnerID)
	return
}

func (d *SqlDb) UpdateRunner(runner db.Runner) (err error) {
	_, err = d.exec(
		"update runner set name=?, project_id=?, inventory_id=?, webhook=?, max_parallel_tasks=? where id=?",
		runner.Name,
		runner.ProjectID,
		runner.InventoryId,
		runner.Webhook,
		runner.MaxParallelTasks,
		runner.ID)

	return
}

func (d *SqlDb) CreateRunner(runner db.Runner) (newRunner db.Runner, err error) {
	token := base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32))

	insertID, err := d.insert(
		"id",
		"insert into runner (name, project_id, inventory_id, token, webhook, max_parallel_tasks) values (?, ?, ?, ?, ?, ?)",
		runner.Name,
		runner.ProjectID,
		runner.InventoryId,
		token,
		runner.Webhook,
		runner.MaxParallelTasks)

	if err != nil {
		return
	}

	newRunner = runner
	newRunner.ID = insertID
	newRunner.Token = token
	return
}
