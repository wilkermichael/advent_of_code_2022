echo "Generating simple code directory for $1"

mkdir -p $1
cd ./$1

echo "package main" > main.go
go mod init $2
go mod tidy