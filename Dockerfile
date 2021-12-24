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
# Copy binary into final image
COPY --from=build /go/src/cmd/supermarket .
# Copy default config file - this can be overridden
# with a volume mount (such as with a config map)
COPY --from=build /go/src/config.json .

# Configure
EXPOSE 8080
ENTRYPOINT ./supermarket
