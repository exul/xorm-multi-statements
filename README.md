# xorm multi statements

This is an example that demonstrates how the `multiStatements` parameter is relevant when using a MySQL DSN.

This code will panic with: `MYSQL_DSN="root:secret@tcp(localhost:3306)/test-db" go run main.go`

This code will exit with code 0 with: `MYSQL_DSN="root:secret@tcp(localhost:3306)/test-db?multiStatements=true" go run main.go`

This example was created in the context of: https://github.com/grafana/grafana/issues/24721.

Example command to spin-up a MySQL instance for testing: `docker run -p 3306:3306 --name mysql-test -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=test-db -e MYSQL_USER=user -e MYSQL_PASSWORD=secret mysql:latest`
