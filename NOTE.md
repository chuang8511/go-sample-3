Ref:
- https://www.youtube.com/watch?v=d_L64KT3SFM&ab_channel=LaithAcademy

bug
- main.go:6:2: no required module provides package github.com/gin-gonic/gin: go.mod file not found in current directory or any parent directory; see 'go help modules'

Solution
add this line in your ~/.zprofile
export GO111MODULE=off

go get -u github.com/gin-gonic/gin

go run main.go



Testing command
curl -d '{"id": "4", "title": "HTTP", "completed": false}' -X POST http://localhost:9090/todos

curl "http://localhost:9090/todos"

curl -X PATCH "http://localhost:9090/todos/1"