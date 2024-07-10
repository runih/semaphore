package db

type RunnerState string

//const (
//	RunnerOffline RunnerState = "offline"
//	RunnerActive  RunnerState = "active"
//)

type Runner struct {
	ID          int    `db:"id" json:"-"`
	Token       string `db:"token" json:"-"`
	ProjectID   *int   `db:"project_id" json:"project_id"`
	InventoryId *int   `db:"inventory_id" json:"inventory_id"`
	//State            RunnerState `db:"state" json:"state"`
	Webhook          string `db:"webhook" json:"webhook"`
	MaxParallelTasks int    `db:"max_parallel_tasks" json:"max_parallel_tasks"`
}
