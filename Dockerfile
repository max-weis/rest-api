FROM baroprime/go as builder

WORKDIR $GOPATH/app/

# download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# compile app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# resulting app
FROM scratch as final
COPY --from=builder go/app/server /app/server
WORKDIR /app
ENTRYPOINT [ "./server" ]
