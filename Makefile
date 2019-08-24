build:
	protoc --micro_out=. --go_out=. proto/vessel/vessel.proto
	go build
	docker build -t shippy-service-vessel .

run:
	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		shippy-service-vessel
