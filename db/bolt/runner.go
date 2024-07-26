package bolt

import (
	"github.com/ansible-semaphore/semaphore/db"
	"github.com/ansible-semaphore/semaphore/util"
)

func (d *BoltDb) GetAllRunners() (runners []db.Runner, err error) {
	return
}

func (d *BoltDb) GetRunner(projectID int, runnerID int) (runner db.Runner, err error) {
	return
}

func (d *BoltDb) GetRunners(projectID int, params db.RetrieveQueryParams) (runners []db.Runner, err error) {
	return
}

func (d *BoltDb) GetRunnerRefs(projectID int, runnerID int) (db.ObjectReferrers, error) {
	return d.getObjectRefs(projectID, db.GlobalRunnerProps, runnerID)
}

func (d *BoltDb) DeleteRunner(projectID int, runnerID int) (err error) {
	return
}

func (d *BoltDb) GetGlobalRunner(runnerID int) (runner db.Runner, err error) {
	err = d.getObject(0, db.GlobalRunnerProps, intObjectID(runnerID), &runner)

	return
}

func (d *BoltDb) GetGlobalRunners() (runners []db.Runner, err error) {
	err = d.getObjects(0, db.GlobalRunnerProps, db.RetrieveQueryParams{}, nil, &runners)
	return
}

func (d *BoltDb) DeleteGlobalRunner(runnerID int) (err error) {
	return
}

func (d *BoltDb) UpdateRunner(runner db.Runner) (err error) {
	return
}

func (d *BoltDb) CreateRunner(runner db.Runner) (newRunner db.Runner, err error) {
	runner.Token = util.RandString(12)

	res, err := d.createObject(0, db.GlobalRunnerProps, runner)
	if err != nil {
		return
	}
	newRunner = res.(db.Runner)
	return
}
