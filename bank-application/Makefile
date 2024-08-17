setup-db-container:
	podman run --name mariadb -e MYSQL_ROOT_PASSWORD=password --network host -d docker.io/library/mariadb:10.3

createdb:
	./database/mysqlclient < ./database/init_up.sql

removedb:
	./database/mysqlclient < ./database/init_up.sql