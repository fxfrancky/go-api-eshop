FROM golang:1.20-alpine3.18 AS builder


WORKDIR /app

COPY wait-for.sh .
COPY app.env .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v3 go build  -v -a -installsuffix cgo -o ./main -tags timetzdata -trimpath cmd/http/main.go

FROM alpine:3.18  

WORKDIR /app
COPY --from=builder /app/main .

COPY   app.env .
COPY   start.sh .
COPY   wait-for.sh .

RUN mkdir -p ./images

CMD [ "/app/main" ]
ENTRYPOINT ["/app/start.sh" ]