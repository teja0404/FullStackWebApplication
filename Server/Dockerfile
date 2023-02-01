FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8081

# Command to run the executable
CMD ["./main"]