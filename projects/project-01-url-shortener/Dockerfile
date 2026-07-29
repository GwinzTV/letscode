FROM golang:1.20-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/urlshortener ./cmd/urlshortener

FROM alpine:3.18
COPY --from=build /out/urlshortener /usr/local/bin/urlshortener
WORKDIR /app
COPY data data
EXPOSE 8080
CMD ["/usr/local/bin/urlshortener"]
