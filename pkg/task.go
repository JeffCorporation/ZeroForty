package pkg

import (
	"database/sql"
	"time"
)

type Task struct {
	Id        int64      `db:"id"`
	Title     string     `db:"title"`
	ProjectId int64      `db:"projectId"`
	Tracking  []Tracking `db:"-"`
	Created   *time.Time `db:"created"`
	Due       *time.Time `db:"due"`
	Completed *time.Time `db:"completed"`
	Priority  uint8      `db:"priority"`
	TimeSpan  int64      `db:"timespan"`
}

//Update task in database. Create new record if id 0
func (t *Task) Update(tx DBOperator) error {
	var err error
	var result sql.Result

	if t.Id <= 0 {
		//New record
		if t.Created == nil {
			now := time.Now()
			t.Created = &now
		}
		if t.ProjectId <= 0 {
			result, err = tx.Exec(sqlInsertTask, t.Title, nil, t.Created.Unix(), t.Due, t.Completed, t.Priority)
		} else {
			result, err = tx.Exec(sqlInsertTask, t.Title, t.ProjectId, t.Created.Unix(), t.Due, t.Completed, t.Priority)
		}
		if err != nil {
			return err
		}

		t.Id, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else {
		//Update record
		if t.ProjectId <= 0 {
			result, err = tx.Exec(sqlUpdateTask, t.Title, nil, t.Due, t.Completed, t.Priority, t.Id)
		} else {
			result, err = tx.Exec(sqlUpdateTask, t.Title, t.ProjectId, t.Due, t.Completed, t.Priority, t.Id)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//Start tracking on this Task
func (t *Task) Start(tx DBOperator) error {
	var err error

	//stop all active tracking
	err = StopTracking(tx)
	if err != nil {
		return err
	}

	//Task need to exist in database
	if t.Id <= 0 {
		err = t.Update(tx)
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(sqlStartTracking, t.Id, time.Now().Unix())

	return err
}

//Stop tracking on this Task
//(stop all pending tasks as only one must be active)
func (t *Task) Stop(tx DBOperator) error {
	return StopTracking(tx)
}

func (t *Task) FetchTracking(db DBOperator, from, to time.Time) error {
	return db.Select(t.Tracking, sqlSelectTracking, t.Id, from, to)
}

func GetActiveTask(db DBOperator) (*Task, error) {
	var t Task
	err := db.Get(&t, sqlSelectActiveTask)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func GetTodoTaskList(db DBOperator) ([]Task, error) {
	var tasks []Task
	err := db.Select(&tasks, sqlSelectTodoTasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
