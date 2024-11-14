package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/livraria/api/database"
	"github.com/livraria/api/models"
	"github.com/livraria/api/controllers"
)



var livros []models.Livro

func main() {
	database.InitDB()
	defer database.CloseDB()

	router := gin.Default()

	router.GET("/api/livros/", controllers.GetLivros)
	router.GET("/api/livros/:id", controllers.GetLivro)
	router.GET("/api/livros/titulo/:title", controllers.GetLivroByTitle)
	router.POST("/api/livros/create", controllers.CreateLivro)
	router.DELETE("/api/livros/:id", controllers.DeleteLivro)

	router.GET("/api/usuarios/", controllers.GetUsuarios)
	router.GET("/api/usuarios/:name", controllers.GetUsuariosByName)
	router.POST("/api/usuarios/create", controllers.CreateGetUsuarios)
	// router.DELETE("/api/livros/:id", controllers.DeleteLivro)

	router.GET("/api/emprestimos/", controllers.GetEmprestimos)
	router.GET("/api/emprestimos/:usuario", controllers.GetEmprestimosByUsuario)
	router.POST("/api/emprestimos/create/:usuario/:titulo", controllers.CreateGetEmprestimos)
    // router.DELETE("/api/livros/:id", controllers.DeleteLivro)


	log.Println("Server is starting on port 8000...")
	router.Run(":8000")
}
