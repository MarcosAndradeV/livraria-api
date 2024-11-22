package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./biblioteca.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}
	log.Println("Connected to database successfully!")

	createTableLivrosStmt := `
    CREATE TABLE IF NOT EXISTS livros (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        titulo TEXT NOT NULL,
        autor TEXT NOT NULL
    );
    `
	_, err = db.Exec(createTableLivrosStmt)
	if err != nil {
		log.Printf("Error creating table: %v\nStatement: %s\n", err, createTableLivrosStmt)
	} else {
		log.Println("Table 'livros' is ready.")
	}

	createTableUsuariosStmt := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id INTEGER PRIMARY KEY,
        nome TEXT NOT NULL,
        email TEXT,
        telefone TEXT
    );
    `
	_, err = db.Exec(createTableUsuariosStmt)
	if err != nil {
		log.Printf("Error creating table: %v\nStatement: %s\n", err, createTableUsuariosStmt)
	} else {
		log.Println("Table 'usuarios' is ready.")
	}

	createTableEmprestimosStmt := `
    CREATE TABLE IF NOT EXISTS emprestimos (
        id INTEGER PRIMARY KEY,
        id_livro INTEGER NOT NULL,
        id_usuario INTEGER NOT NULL,
        data_emprestimo TEXT,
        data_devolucao TEXT,
        FOREIGN KEY (id_livro) REFERENCES livros(id),
        FOREIGN KEY (id_usuario) REFERENCES usuarios(id)
    );
    `
	_, err = db.Exec(createTableEmprestimosStmt)
	if err != nil {
		log.Printf("Error creating table: %v\nStatement: %s\n", err, createTableEmprestimosStmt)
	} else {
		log.Println("Table 'emprestimos' is ready.")
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
