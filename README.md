# iggy










# Stage 1 - Build
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o myapp .

# Stage 2 - Final Image
FROM gcr.io/distroless/static
COPY --from=builder /app/myapp /myapp
ENTRYPOINT ["/myapp"]