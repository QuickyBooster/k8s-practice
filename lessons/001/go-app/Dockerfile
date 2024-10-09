FROM golang:1.23.1-alpine AS build

WORKDIR /build
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run cmd/stdlib/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=build /build/run .
CMD ["/app/run"]