# application name
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/
mysql:
	docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -p 3307:3306 -d mysql:8.0
