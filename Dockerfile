FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


COPY ./ ./

RUN go build -o /main

# EXPOSE 8080

# いったんコメントアウト
# CMD ["/main"]
