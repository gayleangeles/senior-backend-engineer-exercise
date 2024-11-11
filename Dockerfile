FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./server/server.go

ENV DB_PATH="/app/employees.db"
ENV PORT=8080
ENV BIND_JSON=":${PORT}"

EXPOSE ${PORT}

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["./app"]