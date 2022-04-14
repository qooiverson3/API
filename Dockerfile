FROM        golang
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o ces-api
ENTRYPOINT  ["./ces-api","api"]