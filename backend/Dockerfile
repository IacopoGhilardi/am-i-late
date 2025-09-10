# Dockerfile.dev
FROM golang:1.25

WORKDIR /app

# Copio solo go.mod e go.sum per cache dei moduli
COPY go.mod go.sum ./
RUN go mod download

# Copio tutto (ma in compose userai volume override)
COPY . .

# Lancia in sviluppo
CMD ["go", "run", "cmd/server/main.go"]
