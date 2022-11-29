FROM golang:1.17-alpine
WORKDIR /app
COPY *.go ./
RUN go build -o app main.go
EXPOSE 8080
CMD [ "./app" ]
