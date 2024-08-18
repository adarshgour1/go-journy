setup-db-container:
	podman run --name mariadb -e MYSQL_ROOT_PASSWORD=password --network host -d docker.io/library/mariadb:10.3

createdb:
	./database/mysqlclient < ./database/migration/init_up.sql

removedb:
	./database/mysqlclient < ./database/migration/init_up.sql

build:
	go build -o bin/bank-management cmd/main.go

clean:
	rm -rf bin