package main

import (
	"log"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/secure"

	"github.com/livraria/api/controllers"
	"github.com/livraria/api/database"
	"github.com/livraria/api/models"
)

var livros []models.Livro

func main() {
	database.InitDB()
	defer database.CloseDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length", "Content-Type"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
	router.Use(secure.New(secure.Config{
    ContentSecurityPolicy: "default-src 'self'; connect-src 'self'",
}))


	router.GET("/api/livros/", controllers.GetLivros)
	router.GET("/api/livros/:id", controllers.GetLivro)
	router.GET("/api/livros/titulo/:title", controllers.GetLivroByTitle)
	router.POST("/api/livros/create", controllers.CreateLivro)
	router.DELETE("/api/livros/:id", controllers.DeleteLivro)

	router.GET("/api/usuarios/", controllers.GetUsuarios)
	router.GET("/api/usuarios/:name", controllers.GetUsuariosByName)
	router.POST("/api/usuarios/create", controllers.CreateGetUsuarios)
	// router.DELETE("/api/usuarios/:id", controllers.DeleteUsuarios)

	router.GET("/api/emprestimos/", controllers.GetEmprestimos)
	router.GET("/api/emprestimos/:usuario", controllers.GetEmprestimosByUsuario)
	router.POST("/api/emprestimos/create/", controllers.CreateGetEmprestimos)
	// router.DELETE("/api/emprestimos/:id", controllers.DeleteEmprestimos)

	log.Println("Server is starting on port 8000...")
	router.Run(":8000")
}
