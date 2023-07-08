FROM golang:1.20 AS build
WORKDIR /usr/local/src/greeter-service
COPY . .
RUN CGO_ENABLED=0 go build -v

FROM scratch
COPY --from=build /usr/local/src/greeter-service/greeter-service /usr/local/bin/greeter-service
ENTRYPOINT ["/usr/local/bin/greeter-service"]
