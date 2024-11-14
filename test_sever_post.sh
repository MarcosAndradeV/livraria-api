
# curl -X POST \
# http://localhost:8000/api/livros/create \
# -H "Content-Type: application/json" \
# -d '{"titulo": "Livro_Exemplo", "autor": "Nome"}'

# curl -X POST \
# http://localhost:8000/api/usuarios/create \
# -H "Content-Type: application/json" \
# -d '{"nome": "Marcos", "email": "Nome@exemplo.com", "telefone": "9999999999"}'


# curl -X POST \
# http://localhost:8000/api/usuarios/create \
# -H "Content-Type: application/json" \
# -d '{"nome": "Marcos", "email": "Nome@exemplo.com", "telefone": "9999999999"}'

curl -X POST \
http://localhost:8000/api/emprestimos/create/Marcos/1 \
-H "Content-Type: application/json" \
-d '{"data_emprestimo":"2024-11-14T01:47:29.272Z", "data_devolucao": "2024-11-14T01:47:29.272Z"}'
