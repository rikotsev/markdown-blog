FROM gcr.io/distroless/base-debian12:nonroot

ENV SERVER_ADDRESS="0.0.0.0:8080"

COPY dist/markdown-blog-server app

ENTRYPOINT ["/app"]
