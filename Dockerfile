# STAGE 1) Setup the golang server
FROM golang:alpine

# Add our files
ADD . .

RUN go build -o rmc-web .

# Expose the container's port 8080 to our localhost
EXPOSE 8080

# Run the server!
ENTRYPOINT rmc-web
