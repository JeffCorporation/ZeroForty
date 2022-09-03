package pkg

import (
	"github.com/google/uuid"
)

type Project struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Group string `db:"groupName"` //ex: department, client
	UUID  string `db:"uuid"`
	Color string `db:"color"`
}

//generateProjectUUID update task record
func generateProjectUUID(tx DBOperator) error {
	var tasksToUpdate []int
	err := tx.Select(&tasksToUpdate, sqlSelectProjectNullUUID)
	if err != nil {
		return err
	}

	for i := range tasksToUpdate {
		_, err1 := tx.Exec(sqlUpdateProjectUUID, uuid.New().String(), tasksToUpdate[i])
		if err1 != nil {
			return err
		}
	}

	return nil
}
