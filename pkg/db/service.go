package db

import (
	"database/sql"
	"golang.org/x/exp/slog"
	"os"
)

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")

	return &Storage{
		sqliteDatabase,
	}
}

func (s Storage) Close() {
	s.db.Close()
}

func (s Storage) CreateDatabase() {
	slog.Info("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")

	if err != nil {
		panic(err)
	}

	file.Close()
	slog.Info("sqlite-database.db created")
}

func (s Storage) DropDatabase() {
	os.Remove("sqlite-database.db")
}

func (s Storage) FillTable() {
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
}

func (s Storage) CreateTable() {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	slog.Info("Create student table...")
	statement, err := s.db.Prepare(createStudentTableSQL) // Prepare SQL Statement

	if err != nil {
		panic(err)
	}

	statement.Exec() // Execute SQL Statements
	slog.Info("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func (s Storage) InsertStudent(code string, name string, program string) {
	slog.Info("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := s.db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(code, name, program)

	if err != nil {
		panic(err)
	}
}

type Student struct {
	Id      int
	Code    string
	Name    string
	Program string
}

func (s Storage) DisplayStudents() []Student {
	row, err := s.db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		panic(err)
	}

	defer row.Close()
	students := []Student{}

	for row.Next() { // Iterate and fetch the records from result cursor
		s := Student{}
		students = append(students, s)
		row.Scan(&s.Id, &s.Code, &s.Name, &s.Program)
		slog.Info("Students info.", slog.Group("student",
			slog.Int("id", s.Id),
			slog.String("code", s.Code),
			slog.String("name", s.Name),
			slog.String("program", s.Program),
		))

	}

	return students
}
