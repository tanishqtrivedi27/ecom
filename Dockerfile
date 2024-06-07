FROM golang:1.22-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ecom ./cmd/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -o db ./cmd/migrate/main.go

# FROM alpine as release-stage
# WORKDIR /app
# COPY --from=build-stage /app/ecom /app/ecom
# COPY --from=build-stage /app/db /app/db
EXPOSE 8080
# COPY entrypoint.sh /app/entrypoint.sh
# USER root
# RUN chmod +x /app/entrypoint.sh /app/db /app/ecom
ENTRYPOINT [ "/app/ecom" ]
