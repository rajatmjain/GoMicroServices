# Build Stage
FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Final Stage
FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 80
CMD ["./main"]
