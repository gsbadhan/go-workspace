# Build microservice with gin web framework

# setup module
go mod init microservice-gin

# add gin dependency
go get github.com/gin-gonic/gin

# start app
go run main.go

# endpoints without JWT security 
## save album
`curl http://localhost:8090/private/albums     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'`

## get album by id
`curl http://localhost:8090/private/albums/235     --include     --header "Content-Type: application/json"     --request "GET"`

## get all albums
`curl http://localhost:8090/private/albums     --include     --header "Content-Type: application/json"     --request "GET"`

## check service health
`curl http://localhost:8090/public/health     --include     --header "Content-Type: application/json"     --request "GET"`


# endpoints with JWT security

## get jwt access token
`curl http://localhost:8090/public/login     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"username": "admin","password": "admin"}'`

## get all albums
`curl http://localhost:8090/private/albums \
  --include \
  --header "Content-Type: application/json" \
  --header "Authorization: Bearer <access token>" \
  --request "GET"`

## refresh jwt token
`curl http://localhost:8090/private/refresh \
  --include \
  --header "Content-Type: application/json" \
  --header "Authorization: Bearer <access token>" \
  --data '{
    "refresh_token": "<refresh token>"
  }' \
  --request "POST"`

## logout from jwt token
`curl http://localhost:8090/private/logout \
  --include \
  --header "Content-Type: application/json" \
  --header "Authorization: Bearer <access token>" \
  --request "POST"`



