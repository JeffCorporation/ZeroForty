package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	_ "modernc.org/sqlite"
	"punch/pkg"
)

func main() {
	db, err := sqlx.Open("sqlite", "./ZeroForty.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	err = pkg.CreateDatabase(db)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	tx, err = db.Beginx()
	if err != nil {
		log.Fatal(err)
	}

	t := pkg.Task{
		Title:    "Test1",
		Priority: 1,
	}

	//create new task
	err = t.Update(tx)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	//start task
	err = t.Start(tx)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	//get active task
	cur, err := pkg.GetActiveTask(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cur.Title)

}
