api:
	docker-compose -f ./docker/docker-compose.yml up -d --build api

stop:
	docker-compose -f ./docker/docker-compose.yml stop api

remove: stop
	docker-compose -f ./docker/docker-compose.yml rm api

clean: remove
	docker images -f dangling=true -q | xargs docker rmi

pull:
	docker pull golang:1.19-alpine3.16