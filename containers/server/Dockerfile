FROM golang:1.21-alpine AS builder
ARG CONT_IMG_VER
ARG COMMIT_HASH
ENV CONT_IMG_VER=${CONT_IMG_VER:-latest}
ENV COMMIT_HASH=${COMMIT_HASH:-notfound}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags "-X 'agnos/backend/pkgs/version.Version=$CONT_IMG_VER' -X 'agnos/backend/pkgs/version.CommitHash=$COMMIT_HASH'" -o ./cmd/server/server ./cmd/server

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/cmd/server/server ./
CMD ["./server"]
