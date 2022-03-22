server:
	go run main.go
db:
	cd Docker-Compose && sudo docker-compose up 

.PHONY: server db 