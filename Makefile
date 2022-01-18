proto:
	protoc -I . service.proto --go_out=plugins=grpc:.
s1mock:
	mockgen -source=s1/services/search.go -destination s1/services/mocks/mock.go
start:
	sudo docker-compose up -d
test:
	go test ./...