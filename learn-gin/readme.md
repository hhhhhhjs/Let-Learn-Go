# 学习 Gin 框架
## 终端发送请求命令
`$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'`
    可以在终端中直接看到接口返回的响应