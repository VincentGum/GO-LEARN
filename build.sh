kill -9 $(lsof -t -i:8080)
go build .
./learn