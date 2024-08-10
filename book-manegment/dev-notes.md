# Developer Notes

## Basic setup
  ```
  dnf install -y mysql podman go
  ```

## Database
 1.	Download container
    ```
    podman pull docker.io/library/mariadb:10.3
    ```
 2.	To bring up container
    ```
    podman run --name mariadb -e MYSQL_ROOT_PASSWORD=password --network host -d docker.io/library/mariadb:10.3
    ```
 4. Test connection
	```
	 ./script/mysqlclient -e "SELECT 1"
	```
 3. To connect mysql
	```
	./script/mysqlclient
	```
    [Mariadb container setup](https://mariadb.com/kb/en/installing-and-using-mariadb-via-docker/)