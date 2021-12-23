## Build stage
from golang:1.17-alpine AS build

# Build
WORKDIR /go/src
COPY ./ .
WORKDIR cmd
RUN go build -o supermarket


## Final stage
from alpine:3

WORKDIR /root
COPY --from=build /go/src/cmd/supermarket .

# Configure
EXPOSE 8080
ENTRYPOINT ./supermarket
