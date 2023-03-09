FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /starsearch

EXPOSE 8000

CMD [ "/starsearch" ]