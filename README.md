# Pasos de como se inicio

1. Creación de un go mod

> 
    go mod init "nombre_del_modulo"

2. Instalación del compilador de gRPC

>
    go get google.golang.org/grpc

3. Compilación del proto

>
    protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    proto/wishlist.proto