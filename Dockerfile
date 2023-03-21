FROM cgr.dev/chainguard/go AS builder
COPY . /app
RUN cd /app && go build -o kevin .

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/kevin /usr/bin/
ENTRYPOINT ["/usr/bin/kevin"]
