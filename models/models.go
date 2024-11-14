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
	Livro     Livro
	Usuario   Usuario
	DataEmprestimo time.Time `json:"data_emprestimo"`
	DataDevolucao time.Time `json:"data_devolucao"`
}

type Emprestimo_ struct {
	ID        int `json:"ID"`
	Livro     string `json:"titulo"`
	Email   string `json:"email"`
	DataEmprestimo time.Time `json:"data_emprestimo"`
	DataDevolucao time.Time `json:"data_devolucao"`
}
