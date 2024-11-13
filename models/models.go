package models

import "time"

type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

type Usuario struct {
	ID       int `json:"ID"`
	Nome     string `json:"Nome"`
	Email    string `json:"Email"`
	Telefone string `json:"Telefone"`
}

type Emprestimo struct {
	ID        int `json:"ID"`
	Livro     Livro `json:"Livro"`
	Usuario   Usuario`json:"Usuario"`
	DataEmprestimo time.Time
	DataDevolucao time.Time
}
