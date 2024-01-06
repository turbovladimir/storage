package main

import (
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/turbovladimir/storage.git/pkg/db"
	mylog "github.com/turbovladimir/storage.git/pkg/log"
	"golang.org/x/exp/slog"
)

func main() {
	mylog.Init()

	defer func() {
		if err := recover(); err != nil {
			slog.Error("App get panic", err)
		}
	}()

	s := db.New()
	s.CreateTable()

	// INSERT RECORDS
	s.InsertStudent("0001", "Liana Kim", "Bachelor")
	s.InsertStudent("0002", "Glen Rangel", "Bachelor")
	s.InsertStudent("0003", "Martin Martins", "Master")
	s.InsertStudent("0004", "Alayna Armitage", "PHD")
	s.InsertStudent("0005", "Marni Benson", "Bachelor")
	s.InsertStudent("0006", "Derrick Griffiths", "Master")
	s.InsertStudent("0007", "Leigh Daly", "Bachelor")
	s.InsertStudent("0008", "Marni Benson", "PHD")
	s.InsertStudent("0009", "Klay Correa", "Bachelor")

	// DISPLAY INSERTED RECORDS
	s.DisplayStudents()

	s.Close()
}
