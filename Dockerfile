FROM golang:1.23.0

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY configs ./configs
COPY entities ./entities
COPY errors ./errors
COPY handlers ./handlers
COPY migrations ./migrations
COPY pkg ./pkg
COPY repos ./repos
COPY services ./services
COPY static ./static

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Run
ENTRYPOINT [ "/src/app" ]