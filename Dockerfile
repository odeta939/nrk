FROM golang:1.25.1-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/nrk-CLI ./cmd/nrk-CLI

FROM gcr.io/distroless/static:nonroot
COPY --from=build /out/nrk-CLI /usr/local/bin/nrk-CLI
USER nonroot
ENTRYPOINT ["/usr/local/bin/nrk-CLI"]
