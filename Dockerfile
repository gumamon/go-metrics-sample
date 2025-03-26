# --- Build Stage ---
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# 依存関係を先にコピーしてキャッシュ活用
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# --- Final Stage ---
FROM gcr.io/distroless/static:nonroot

# 実行バイナリをコピー（非rootユーザで実行）
COPY --from=builder /app/app /app/app

# ポート（ドキュメント用途）
EXPOSE 8080
EXPOSE 8000

# エントリポイント
ENTRYPOINT ["/app/app"]
