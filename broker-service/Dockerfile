FROM golang AS build

WORKDIR /app

COPY * .

ADD data/ /app/data

RUN CGO_ENABLED=0 go build -o server .

FROM gcr.io/distroless/base-debian11

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server /app/server

EXPOSE 8080

ENTRYPOINT ["./server"]

