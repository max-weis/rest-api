FROM golang:1.12-alpine3.9 as builder

ENV GO111MODULE=on

WORKDIR $GOPATH/app/

RUN apk add git
COPY go.mod .
COPY go.sum .
RUN go mod download
#compile app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

#resulting app
FROM scratch as final
COPY --from=builder go/app/server /app/server
WORKDIR /app
ENTRYPOINT [ "./server" ]