package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/livraria/api/database"
	"github.com/livraria/api/models"
)

func GetLivros(c *gin.Context) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, titulo, autor FROM livros")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot query books: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var livros []models.Livro

	for rows.Next() {
		var livro models.Livro
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

func GetLivro(c *gin.Context) {

	param := c.Param("id")
	var livro models.Livro
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

func GetLivroByTitle(c *gin.Context) {
	param := c.Param("title")
	var livro models.Livro
	db := database.GetDB()
	err := db.QueryRow("SELECT titulo, autor FROM livros WHERE titulo = ?", param).Scan(&livro.Titulo, &livro.Autor)

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

func CreateLivro(c *gin.Context) {
	var livro models.Livro
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

func DeleteLivro(c *gin.Context) {
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
		c.JSON(404, gin.H{"error": "models.Livro não encontrado"})
		return
	}

	c.JSON(200, gin.H{"message": "models.Livro deletado com sucesso"})
}

func GetUsuarios(c *gin.Context) {
	db := database.GetDB()
	rows, err := db.Query("SELECT nome, email, telefone FROM usuarios")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot query books: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario
		if err := rows.Scan(&usuario.Nome, &usuario.Email, &usuario.Telefone); err != nil {
			c.JSON(500, gin.H{
				"error": "cannot scan book data: " + err.Error(),
			})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": "error occurred during row iteration: " + err.Error(),
		})
		return
	}

	c.JSON(200, usuarios)
}

func GetUsuariosByName(c *gin.Context) {
	param := c.Param("nome")
	var usuario models.Usuario
	db := database.GetDB()
	err := db.QueryRow("SELECT nome, email, telefone FROM usuarios WHERE nome = ?", param).Scan(&usuario.Nome, &usuario.Email, &usuario.Telefone)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "usuario não encontrado",
			})
		} else {
			c.JSON(500, gin.H{
				"error": "usuario não pode ser encontrado " + err.Error(),
			})
		}
		return
	}

	c.JSON(200, usuario)
}

func CreateGetUsuarios(c *gin.Context) {

	var usuario models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": "cannot bind JSON: " + err.Error()})
		return
	}

	db := database.GetDB()
	result, err := db.Exec(
		"INSERT INTO usuarios (nome, email, telefone) VALUES (?, ?, ?)",
		usuario.Nome,
		usuario.Email,
		usuario.Telefone,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possivel criar o Usuario: " + err.Error()})
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot retrieve last insert ID: " + err.Error()})
		return
	}
	usuario.ID = int(lastID)

	c.JSON(200, usuario)

}

func GetEmprestimos(c *gin.Context) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id_livro, id_usuario, data_emprestimo, data_devolucao FROM emprestimos")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var emprestimos []models.Emprestimo

	for rows.Next() {
		var emprestimo models.Emprestimo
		var livroId int
		var usuarioId int
		if err := rows.Scan(&livroId, &usuarioId, &emprestimo.DataEmprestimo, &emprestimo.DataDevolucao); err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
			})
			return
		}
		if err := db.QueryRow("SELECT titulo, autor FROM livros WHERE id = ?", livroId).Scan(&emprestimo.Livro.Titulo, &emprestimo.Livro.Autor); err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
			})
			return
		}
		emprestimo.Livro.ID = livroId
		if err := db.QueryRow("SELECT id, nome, email, telefone FROM usuarios WHERE id = ?", usuarioId).Scan(
			&emprestimo.Usuario.ID,
			&emprestimo.Usuario.Nome,
			&emprestimo.Usuario.Email,
			&emprestimo.Usuario.Telefone); err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel selecionar o emprestimo Usuario não encontrado: " + err.Error(),
			})
			return
		}
		emprestimos = append(emprestimos, emprestimo)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": "error occurred during row iteration: " + err.Error(),
		})
		return
	}

	c.JSON(200, emprestimos)
}

func GetEmprestimosByUsuario(c *gin.Context) {

	param := c.Param("usuario")
	var usuario models.Usuario
	db := database.GetDB()
	err := db.QueryRow("SELECT id, nome, email, telefone FROM usuarios WHERE nome = ?", param).Scan(&usuario.ID,
	&usuario.Nome,
	&usuario.Email,
	&usuario.Telefone)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "usuario não encontrado" + err.Error(),
			})
		} else {
			c.JSON(500, gin.H{
				"error": "usuario não pode ser encontrado " + err.Error(),
			})
		}
		return
	}

	rows, err := db.Query("SELECT id_livro, id_usuario, data_emprestimo, data_devolucao FROM emprestimos WHERE id_usuario = ?", usuario.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var emprestimos []models.Emprestimo

	for rows.Next() {
		var emprestimo models.Emprestimo
		var usuarioID int
		var livroId int
		if err := rows.Scan(&livroId, &usuarioID, &emprestimo.DataEmprestimo, &emprestimo.DataDevolucao); err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
			})
			return
		}
		if err := db.QueryRow("SELECT titulo, autor FROM livros WHERE id = ?", livroId).Scan(&emprestimo.Livro.Titulo,
			&emprestimo.Livro.Autor); err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel selecionar o emprestimo: " + err.Error(),
			})
			return
		}
		emprestimo.Usuario = usuario
		emprestimos = append(emprestimos, emprestimo)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{
			"error": "error occurred during row iteration: " + err.Error(),
		})
		return
	}
	if len(emprestimos) > 0 {
		c.JSON(200, emprestimos)

	} else {
		c.JSON(200, gin.H{
			"message": "Sem Emprestimos",
		})
	}
}

func CreateGetEmprestimos(c *gin.Context) {
	var emprestimo_ models.Emprestimo_

	if err := c.BindJSON(&emprestimo_); err != nil {
		c.JSON(400, gin.H{"error": "cannot bind JSON: " + err.Error()})
		return
	}

	var usuario models.Usuario
	db := database.GetDB()
	err := db.QueryRow("SELECT id, nome, email, telefone FROM usuarios WHERE email = ?", emprestimo_.Email).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.Telefone)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "usuario não encontrado" + err.Error(),
			})
		} else {
			c.JSON(500, gin.H{
				"error": "usuario não pode ser encontrado " + err.Error(),
			})
		}
		return
	}
	var livro models.Livro
	err = db.QueryRow("SELECT id, titulo, autor FROM livros WHERE titulo = ?", emprestimo_.Livro).Scan(&livro.ID, &livro.Titulo, &livro.Autor)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "livro não encontrado " + err.Error(),
			})
		} else {
			c.JSON(500, gin.H{
				"error": "livro não pode ser encontrado " + err.Error(),
			})
		}
		return
	}

	var emprestimo models.Emprestimo
	emprestimo.Livro = livro
	emprestimo.Usuario = usuario
	emprestimo.DataEmprestimo = emprestimo_.DataEmprestimo
	emprestimo.DataDevolucao = emprestimo_.DataDevolucao

	result, err := db.Exec(
  		// id INTEGER PRIMARY KEY,
        // id_livro INTEGER NOT NULL,
        // id_usuario INTEGER NOT NULL,
        // data_emprestimo DATE,
        // data_devolucao DATE,
		"INSERT INTO emprestimos (id_livro, id_usuario, data_emprestimo, data_devolucao) VALUES (?, ?, ?, ?)",
		emprestimo.Livro.ID,
		emprestimo.Usuario.ID,
		emprestimo.DataEmprestimo,
		emprestimo.DataDevolucao,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possivel criar o Emprestimo: " + err.Error()})
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot retrieve last insert ID: " + err.Error()})
		return
	}
	emprestimo.ID = int(lastID)

	c.JSON(200, emprestimo)
}
