
FROM golang:1.17.5
WORKDIR /ethservice/eth
# Copy everything from this project into the filesystem of the container.
COPY . .
# Compile the binary exe for our app.
RUN go build -o main .
# Start the application.
EXPOSE 80
CMD ["./main"]