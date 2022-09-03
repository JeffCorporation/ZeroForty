package pkg

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_Operations(test *testing.T) {
	db, err := sqlx.Open("sqlite", ":memory:")
	if err != nil {
		test.Fatalf("sqlx Open() error = %v", err)
	}
	defer db.Close()

	err = CreateDatabase(db)
	if err != nil {
		test.Fatalf("createDatabase() error = %v", err)
	}

	//Create new Task
	t := Task{
		Title:    "Test",
		Priority: 1,
	}
	err = t.Update(db)
	if err != nil {
		test.Errorf("Create new Task error = %v", err)
	}
	if t.Id != 1 {
		test.Errorf("Create new Task id = %d, want 1", t.Id)
	}

	//Update Task
	t.Priority = 2
	t.Title = "ok"
	err = t.Update(db)
	if err != nil {
		test.Errorf("Update Task error = %v", err)
	}

	//Start Task
	err = t.Start(db)
	if err != nil {
		test.Errorf("Start Task error = %v", err)
	}

	//Get Active Task
	t1, err := GetActiveTask(db)
	if err != nil {
		test.Errorf("GetActiveTask() error = %v", err)
	}
	if t.Id != t1.Id {
		test.Errorf("GetActiveTask() id = %d, want %d", t1.Id, t.Id)
	}

	todo, err := GetTodoTaskList(db)
	if err != nil {
		test.Errorf("GetTodoTaskList() error = %v", err)
	}
	if len(todo) == 0 {
		test.Errorf("GetTodoTaskList() len = 0, want > 0")
	}

}
