FROM golang:alpine
ADD . /go/src/now-go
RUN go install now-go
CMD ["/go/bin/now-go"]
EXPOSE 3000