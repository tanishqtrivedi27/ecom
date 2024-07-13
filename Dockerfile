FROM golang:1.22-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ecom ./cmd/main.go

FROM alpine AS release-stage
WORKDIR /app
COPY --from=build-stage /app/ecom /app/ecom
EXPOSE 8080
RUN chmod +x /app/ecom
CMD [ "/app/ecom" ]
