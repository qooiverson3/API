FROM        golang
RUN         mkdir -p /app
RUN         mkdir -p /app/logs
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app","api"]