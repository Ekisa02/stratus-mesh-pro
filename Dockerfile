FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN apk add --no-cache git
COPY . .
RUN go build -o /bin/stratus-mesh ./cmd/stratus-mesh

FROM alpine:3.17
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/stratus-mesh /usr/local/bin/stratus-mesh
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/stratus-mesh"]
