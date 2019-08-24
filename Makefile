build:
	protoc --micro_out=. --go_out=. proto/vessel/vessel.proto
	go build

run:
	./shippy-service-vessel
