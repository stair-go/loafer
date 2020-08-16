FROM golang:1.13

WORKDIR /root/github.com/stair-go/loafer

COPY . .
RUN go build -mod=vendor

WORKDIR /app
#COPY --from=builder /root/github.com/stair-go/loafer .
EXPOSE 8189
ENTRYPOINT ["./github.com/stair-go/loafer"]
#CMD ["-mode=release"]