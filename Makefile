generate-proto:
			cd src/proto && protoc --go_out=../grpc_server/pb --go_opt=paths=source_relative --go-grpc_out=../grpc_server/pb --go-grpc_opt=paths=source_relative ./*

build-grpc-server:
			cd src/grpc_server/main && go build -o ../../../bin/london-express-server.exe main.go

build-grpc-client:
			cd src/grpc_client && go build -o ../../bin/london-express-client.exe client.go

build-all:
		cd src/grpc_server/main && go build -o ../../../bin/london-express-server.exe main.go
		cd src/grpc_client && go build -o ../../bin/london-express-client.exe client.go
