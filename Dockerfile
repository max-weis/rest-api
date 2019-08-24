FROM golang:1.12-alpine3.9 as builder

ARG APP_NAME

ENV GO111MODULE=on

WORKDIR $GOPATH/app

RUN apk add git
COPY go.* ./
RUN go mod download
#compile app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o $APP_NAME ./cmd/$APP_NAME/$APP_NAME.go 

#resulting app
FROM scratch as final
COPY --from=builder go/$APP_NAME /app/$APP_NAME
WORKDIR /app
CMD ./$APP_NAME
