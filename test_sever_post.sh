
curl -X POST \
http://localhost:8000/api/livros/create \
-H "Content-Type: application/json" \
-d '{"titulo": "Livro Exemplo", "autor": "Nome"}'
