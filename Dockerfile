FROM golang:1.13

WORKDIR /root/daydaytest

COPY . .
RUN go build -mod=vendor

WORKDIR /app
#COPY --from=builder /root/daydaytest .
EXPOSE 8189
ENTRYPOINT ["./daydaytest"]
#CMD ["-mode=release"]