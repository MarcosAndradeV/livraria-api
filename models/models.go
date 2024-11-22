package models

type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

type Usuario struct {
	ID       int    `json:"ID"`
	Nome     string `json:"Nome"`
	Email    string `json:"Email"`
	Telefone string `json:"Telefone"`
}

type Emprestimo struct {
	ID             int `json:"ID"`
	Livro          Livro
	Usuario        Usuario
	DataEmprestimo string `json:"data_emprestimo"`
	DataDevolucao  string `json:"data_devolucao"`
}

type Emprestimo_ struct {
	ID             int    `json:"ID"`
	Livro          string `json:"titulo"`
	Email          string `json:"email"`
	DataEmprestimo string `json:"data_emprestimo"`
	DataDevolucao  string `json:"data_devolucao"`
}
