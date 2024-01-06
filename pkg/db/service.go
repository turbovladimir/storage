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
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	slog.Info("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file

	if err != nil {
		panic(err)
	}

	file.Close()
	slog.Info("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File

	return &Storage{
		sqliteDatabase,
	}
}

func (s Storage) Close() {
	s.db.Close()
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

func (s Storage) DisplayStudents() {
	row, err := s.db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		panic(err)
	}

	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)
		slog.Info("Students info.", slog.Group("student",
			slog.Int("id", id),
			slog.String("code", code),
			slog.String("name", name),
			slog.String("program", program),
		))
	}
}
