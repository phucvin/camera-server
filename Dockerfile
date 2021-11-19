FROM golang:1.17
WORKDIR /
COPY main.go /
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /main ./
CMD ["./main"]  