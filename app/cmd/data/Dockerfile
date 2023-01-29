FROM golang:1.19.4-alpine 

WORKDIR /app

COPY go.mod go.sum ./

COPY app ./app
COPY db ./db

RUN ls

RUN go build -o ./data app/cmd/data/main.go

RUN ls

EXPOSE 9090
CMD ["./data"]
