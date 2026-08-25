FROM golang:1.22
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o /main .
CMD ["/main"]
