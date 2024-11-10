package main

import (
	"database/sql"
	"errors"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

type usuarios struct {
	ID       uint
	nome     string
	email    string
	telefone string
}

var livros []Livro

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()

	router.GET("/api/livros/", getLivros)
	router.GET("/api/livros/:id", getLivro)
	router.POST("/api/livros/create", createLivro)
	router.DELETE("/api/livros/:id", deleteLivro)

	log.Println("Server is starting on port 8000...")
	router.Run(":8000")
}
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

var db *sql.DB

func initDB() {
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
        data_emprestimo DATE,
        data_devolucao DATE,
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

func getLivros(c *gin.Context) {

	rows, err := db.Query("SELECT id, titulo, autor FROM livros")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot query books: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var livros []Livro

	for rows.Next() {
		var livro Livro
		if err := rows.Scan(&livro.ID, &livro.Titulo, &livro.Autor); err != nil {
			c.JSON(500, gin.H{
				"error": "cannot scan book data: " + err.Error(),
			})
			return
		}
		livros = append(livros, livro)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": "error occurred during row iteration: " + err.Error(),
		})
		return
	}

	c.JSON(200, livros)
}

func getLivro(c *gin.Context) {

	param := c.Param("id")
	var livro Livro

	err := db.QueryRow("SELECT id, titulo, autor FROM livros WHERE id = ?", param).Scan(&livro.ID, &livro.Titulo, &livro.Autor)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "book not found",
			})
		} else {
			c.JSON(500, gin.H{
				"error": "cannot find book: " + err.Error(),
			})
		}
		return
	}

	c.JSON(200, livro)
}

func createLivro(c *gin.Context) {
	var livro Livro
	if err := c.BindJSON(&livro); err != nil {
		c.JSON(400, gin.H{"error": "cannot bind JSON: " + err.Error()})
		return
	}

	log.Printf("ID: %v\n", livro.ID)
	log.Printf("Titulo: %v\n", livro.Titulo)
	log.Printf("Autor: %v\n", livro.Autor)

	result, err := db.Exec(
		"INSERT INTO livros (titulo, autor) VALUES (?, ?)",
		livro.Titulo,
		livro.Autor,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot create book: " + err.Error()})
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot retrieve last insert ID: " + err.Error()})
		return
	}
	livro.ID = int(lastID)

	c.JSON(200, livro)
}

func deleteLivro(c *gin.Context) {
	param := c.Param("id")

	result, err := db.Exec("DELETE FROM livros WHERE id = ?", param)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot delete livro: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot retrieve rows affected: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(200, gin.H{"message": "Livro deletado com sucesso"})
}
