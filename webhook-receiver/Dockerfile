FROM golang:1.15-alpine as build

ENV CGO_ENABLED=0
WORKDIR /go/src

RUN apk update \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates

COPY ./go.* ./
RUN go mod download

COPY . .
RUN go build -o /out/app .

FROM scratch AS bin
COPY --from=build /out/app /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]
