# Use an official Go runtime as a parent image
FROM golang:1.22
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -v -t -d ./...
CMD ["go", "run", "ssr.go"]