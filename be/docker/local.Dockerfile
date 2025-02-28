FROM golang:1.23 AS build

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y unzip

RUN make setup build

FROM gcr.io/distroless/base-debian12:nonroot

USER nonroot

COPY --from=build /go/src/app/dist /

ENTRYPOINT ["/markdown-blog-server"]
