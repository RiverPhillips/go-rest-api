FROM cgr.dev/chainguard/go AS builder

ADD . /app
RUN cd /app && go build -o server ./cmd/server/main.go

FROM cgr.dev/chainguard/static:latest

COPY --from=builder /app/server /usr/bin/

EXPOSE 8080
CMD ["/usr/bin/server"]