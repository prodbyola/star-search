FROM golang:latest
RUN go get -u github.com/cosmtrek/air
go mod download
WORKDIR /app
ENTRYPOINT ["air"]