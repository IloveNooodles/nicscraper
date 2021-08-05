FROM golang:1.15-buster as build
WORKDIR /src
COPY . .
RUN go get -d -v ./...
RUN go build -o /src/nicscraper main.go

FROM gcr.io/distroless/base-debian10
COPY --from=build /src/nicscraper /nicscraper
ENTRYPOINT ["/nicscraper"]