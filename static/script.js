// Classe Livro, Pessoa e Emprestimo para os objetos de cadastro
class Livro {
    constructor(titulo, autor) {
        this.titulo = titulo;
        this.autor = autor;
    }
}

class Pessoa {
    constructor(nome, email, telefone) {
        this.nome = nome;
        this.email = email;
        this.telefone = telefone;
    }
}

class Emprestimo {
    constructor(data_emprestimo, data_devolucao) {
        this.data_emprestimo = data_emprestimo;
        this.data_devolucao = data_devolucao;
    }
}

// Função para adicionar um livro
async function adicionarLivro() {
    const titulo = document.getElementById('titulo').value;
    const autor = document.getElementById('autor').value;

    try {
        const response = await fetch('http://localhost:8000/api/livros/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify({ titulo, autor })
        });

        if (response.ok) {
            document.getElementById('titulo').value = "";
            document.getElementById('autor').value = "";
            alert("Livro adicionado com sucesso!");
        } else {
            alert("Erro ao adicionar o livro.");
        }
    } catch (error) {
        console.error("Erro ao adicionar livro:", error);
        alert("Erro ao adicionar livro.");
    }
}

// Função para adicionar um usuário
async function adicionarUsuario() {
    let nome = document.getElementById('nome').value;
    let email = document.getElementById('email').value;
    let telefone = document.getElementById('telefone').value;

    if (!nome || !email || !telefone) {
        alert("Por favor, preencha todos os campos.");
        return;
    }

    if (!email.includes("@")) {
        email = `${email}@gmail.com`;
    }

    let usuario = {nome, email, telefone}

    try {
        const response = await fetch('http://localhost:8000/api/usuarios/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify( usuario )
        });

        if (response.ok) {
            document.getElementById('nome').value = "";
            document.getElementById('email').value = "";
            document.getElementById('telefone').value = "";
            alert("Usuário adicionado com sucesso!");
        } else {
            alert("Erro ao adicionar o usuário.");
        }
    } catch (error) {
        console.error("Erro ao adicionar usuário:", error);
        alert("Erro ao adicionar usuário.");
    }
}

// Função para criar um empréstimo
async function criarEmprestimo() {
    const usuario = document.getElementById('usuario').value;
    const titulo = document.getElementById('titulo-livro').value;
    const dataEmprestimo = document.getElementById('data-emprestimo').value;
    const dataDevolucao = document.getElementById('data-devolucao').value;
    const emprestimo = new Emprestimo(dataEmprestimo, dataDevolucao);

    try {
        const response = await fetch(`http://localhost:8000/api/emprestimos/create/${usuario}/${titulo}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify(emprestimo)
        });

        if (response.ok) {
            document.getElementById('usuario').value = "";
            document.getElementById('titulo-livro').value = "";
            document.getElementById('data-emprestimo').value = "";
            document.getElementById('data-devolucao').value = "";
            alert("Empréstimo registrado com sucesso!");
        } else {
            alert("Erro ao registrar o empréstimo.");
        }
    } catch (error) {
        console.error("Erro ao registrar empréstimo:", error);
        alert("Erro ao registrar empréstimo.");
    }
}

// Função para exibir usuários cadastrados
async function mostrarUsuarios() {
    try {
        const response = await fetch('http://localhost:8000/api/usuarios/');
        if (response.ok) {
            const usuarios_ = await response.json();
            const usuarios = usuarios_ == null ? [] : usuarios_
            const lista = document.getElementById('search-results');
          if (usuarios.length == 0) {
            lista.innerHTML = "<li>Nenhum resultado encontrado.</li>"
          } else {
            lista.innerHTML = usuarios.map(usuario =>
              `<li><strong>Nome:</strong> ${usuario.Nome}, <strong>Email:</strong> ${usuario.Email}, <strong>Telefone:</strong> ${usuario.Telefone}</li>`
            ).join('');
          }
        } else {
            console.error("Erro ao carregar a lista de usuários.");
            alert("Erro ao carregar a lista de usuários.");
        }
    } catch (error) {
        console.error("Erro ao exibir usuários:", error);
        alert("Erro ao exibir usuários.");
    }
}

// Função para exibir livros cadastrados
async function mostrarLivros() {
    try {
        const response = await fetch('http://localhost:8000/api/livros/');
        if (response.ok) {
            const livros_ = await response.json();
            const livros = livros_ == null ? [] : livros_
            const lista = document.getElementById('search-results');
          if (livros.length == 0) {
            lista.innerHTML = "<li>Nenhum resultado encontrado.</li>"
          } else {
            lista.innerHTML = livros.map(livro =>
              `<li><strong>Título:</strong> ${livro.titulo}, <strong>Autor:</strong> ${livro.autor}</li>`
            ).join('');
          }
        } else {
            alert("Erro ao carregar a lista de livros.");
        }
    } catch (error) {
        console.error("Erro ao exibir livros:", error);
        alert("Erro ao exibir livros.");
    }
}

// Função para exibir empréstimos cadastrados
async function mostrarEmprestimos() {
    try {
        const response = await fetch('http://localhost:8000/api/emprestimos/');
        if (response.ok) {
            const emprestimos_ = await response.json();
            const emprestimos = emprestimos_ == null ? [] : emprestimos_
            const lista = document.getElementById('search-results');
            if (emprestimos.length == 0) {
                lista.innerHTML = "<li>Nenhum resultado encontrado.</li>"
            } else {
              lista.innerHTML = emprestimos.map(emprestimo =>
                  `<li><strong>Usuário:</strong> ${emprestimo.Usuario.Nome}, <strong>Livro:</strong> ${emprestimo.Livro.titulo}</li>`
              ).join('');
            }
        } else {
            alert("Erro ao carregar a lista de empréstimos.");
        }
    } catch (error) {
        console.error("Erro ao exibir empréstimos:", error);
        alert("Erro ao exibir empréstimos.");
    }
}

// Função para buscar usuários e livros
async function search() {
    const query = document.getElementById("search").value.toLowerCase();
    const searchResults = [];

    try {
        // Buscar usuários
        const responseUsuarios = await fetch('http://localhost:8000/api/usuarios/');
        if (responseUsuarios.ok) {
            const usuarios_ = await responseUsuarios.json();
            const usuarios = usuarios_ == null ? [] : usuarios_
            usuarios.forEach(usuario => {
                if (usuario.Nome.toLowerCase().includes(query)) {
                    searchResults.push(
                        `<li><strong>Nome:</strong> ${usuario.Nome}, <strong>Email:</strong> ${usuario.Email}, <strong>Telefone:</strong> ${usuario.telefone}</li>`
                    );
                }
            });
        } else {
            console.error("Erro ao carregar usuários.");
        }

        // Buscar livros
        const responseLivros = await fetch('http://localhost:8000/api/livros/');
        if (responseLivros.ok) {
            const livros_ = await responseLivros.json();
            const livros = livros_ == null ? [] : livros_
            livros.forEach(livro => {
                if (
                    livro.titulo.toLowerCase().includes(query) ||
                    livro.autor.toLowerCase().includes(query)
                ) {
                    searchResults.push(
                        `<li><strong>Título:</strong> ${livro.titulo}, <strong>Autor:</strong> ${livro.autor}</li>`
                    );
                }
            });
        } else {
            console.error("Erro ao carregar livros.");
        }

        // Atualizar a interface com os resultados
        const searchResultsContainer = document.getElementById("search-results");
        if (searchResults.length > 0) {
            searchResultsContainer.innerHTML = searchResults.join('');
        } else {
            searchResultsContainer.innerHTML = "<li>Nenhum resultado encontrado.</li>";
        }
    } catch (error) {
        console.error("Erro ao realizar a busca:", error);
        alert("Erro ao realizar a busca. Tente novamente.");
    }
};
