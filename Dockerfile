FROM golang:1.12-alpine3.9 as builder

ENV APP_GOPATH $GOPATH/src/gitlab.com/baroprime/app/
ENV GO111MODULE=on

WORKDIR $APP_GOPATH

RUN apk add git
COPY go.* $APP_GOPATH
RUN go mod download
#compile app
COPY . $APP_GOPATH
RUN CGO_ENABLED=0 GOOS=linux go build -o rest ./cmd/prod-rest/prod-rest.go 

#resulting app
FROM scratch as final
COPY --from=builder go/src/gitlab.com/baroprime/app/rest /app/rest
WORKDIR /app
ENTRYPOINT ["./rest"]
