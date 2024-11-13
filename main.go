package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/livraria/api/database"
	"database/sql"
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
	database.InitDB()
	defer database.CloseDB()

	router := gin.Default()

	router.GET("/api/livros/", getLivros)
	router.GET("/api/livros/:id", getLivro)
	router.POST("/api/livros/create", createLivro)
	router.DELETE("/api/livros/:id", deleteLivro)

	log.Println("Server is starting on port 8000...")
	router.Run(":8000")
}



func getLivros(c *gin.Context) {
	db := database.GetDB()
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
	db := database.GetDB()
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
	db := database.GetDB()
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
	db := database.GetDB()
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
