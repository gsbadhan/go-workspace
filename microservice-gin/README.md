# Build microservice with gin web framework

# setup module
go mod init microservice-gin

# add gin dependency
go get github.com/gin-gonic/gin

# start app
go run main.go


# save album
`curl http://localhost:8090/albums     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'`

# get album by id
`curl http://localhost:8090/albums/1     --include     --header "Content-Type: application/json"     --request "GET"`

# get all albums
`curl http://localhost:8090/albums     --include     --header "Content-Type: application/json"     --request "GET"`

# check service health
`curl http://localhost:8090/health     --include     --header "Content-Type: application/json"     --request "GET"`




