FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]