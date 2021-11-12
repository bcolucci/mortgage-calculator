FROM golang:1.17 AS bin_builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build-bin

FROM node:16 AS ui_builder
COPY --from=bin_builder /app/ui /app/ui
RUN cd /app/ui && npm i && \
    node_modules/parcel/lib/bin.js build src/index.js --no-cache --dist-dir static/dist

FROM alpine
COPY --from=bin_builder /app/bin/mortgage-server /app/mortgage-server
COPY --from=ui_builder /app/ui /app/ui
EXPOSE 3000
ENV UI_PATH=./app/ui/static
CMD ["/app/mortgage-server"]
