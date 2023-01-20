FROM golang:1.17-alpine

ENV CGO_ENABLED=1

RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN pwd
RUN ls

RUN go build -mod=mod -o main .

 EXPOSE 5000

CMD [ "/app/main" ]