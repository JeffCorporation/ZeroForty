package pkg

import (
	"time"
)

type Tracking struct {
	//TaskId int64      `db:"taskId"`
	Start time.Time  `db:"start"`
	Stop  *time.Time `db:"stop"`
}

//Duration from start to stop. Use current time if the task is active
func (tr Tracking) Duration() time.Duration {

	if tr.Stop == nil {
		return time.Now().Sub(tr.Start)
	}
	return tr.Stop.Sub(tr.Start)
}

//StopTracking End all pending tracking by setting current time
func StopTracking(tx DBOperator) error {
	_, err := tx.Exec(sqlStopTracking, time.Now().Unix())
	return err
}
