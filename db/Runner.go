package db

type RunnerState string

//const (
//	RunnerOffline RunnerState = "offline"
//	RunnerActive  RunnerState = "active"
//)

type Runner struct {
	ID            int     `db:"id" json:"id"`
	Token         string  `db:"token" json:"token"`
	ProjectID     *int    `db:"project_id" json:"project_id"`
	ProjectName   *string `db:"project_name" json:"project_name"`
	InventoryId   *int    `db:"inventory_id" json:"inventory_id"`
	InventoryName *string `db:"inventory_name" json:"inventory_name"`
	//State            RunnerState `db:"state" json:"state"`
	Webhook          string `db:"webhook" json:"webhook"`
	MaxParallelTasks int    `db:"max_parallel_tasks" json:"max_parallel_tasks"`
}
