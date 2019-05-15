# Use an official go image as a build image
FROM golang:alpine AS build

# Set the working directory
WORKDIR /go/src/youtube-mosaic/

# Copy the current directory contents in the container workdir
COPY cmd/server/main.go .

# Define environment variable
#RUN go build hello.go 
RUN go build -o server


FROM alpine:latest
EXPOSE 80

WORKDIR /bin/youtube-mosaic/

COPY --from=build /go/src/youtube-mosaic/server .
COPY configs/server-config.json ./configs/
COPY www/* ./www/

RUN apk add ca-certificates

ENTRYPOINT ["/bin/youtube-mosaic/server"]