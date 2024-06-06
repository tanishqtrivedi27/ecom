FROM golang:1.22-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ecom ./cmd/main.go
RUN go build -o db ./cmd/migrate/main.go

FROM alpine AS release-stage
WORKDIR /app
COPY --from=build-stage /app/ecom ecom
COPY --from=build-stage /app/db db
EXPOSE 8080
COPY entrypoint.sh entrypoint.sh
RUN chmod +x ./entrypoint.sh
ENTRYPOINT [ "/app/entrypoint.sh" ]
