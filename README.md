### Docker run postgres
`docker run --name=user-restapi -e POSTGRES_PASSWORD="123" -p 5436:5432 -d --rm postgres`

### Docker run redis
`docker run --name redis -p 6369:6379 -d redis`

### Migrate Up
`migrate -path ./migrate -database 'postgres://postgres:123@localhost:5436/postgres?sslmode=disable' up`