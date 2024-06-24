run:
	sudo docker build -t backend-golang -f api.Dockerfile .
	sudo docker compose up

test:
	sudo docker build -t tests-golang-test -f tests/tests.Dockerfile .
	sudo docker compose -f tests/docker-compose-tests.yml up 

local:
	sudo docker compose up postgres -d
	sudo docker compose up redis -d
	go build cmd/api/main.go
	./main