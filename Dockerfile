# より新しいglibcバージョンを使用する互換性のあるベースイメージを使用
FROM debian:bookworm-slim AS base

# Goの環境変数を設定
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 必要な依存関係をインストール
RUN apt-get update && apt-get install -y --no-install-recommends \
    wget gcc build-essential ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# ステージ1: Goアプリケーションのビルド
FROM golang:1.20 AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係ファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY ./cmd ./cmd
COPY ./internal ./internal

# Goアプリケーションをビルド
WORKDIR /app/cmd
RUN go build -o /app/main .

# ステージ2: 最小限の最終イメージを作成
FROM base AS final

# 作業ディレクトリを設定
WORKDIR /app

# ビルダーからアプリケーションバイナリをコピー
COPY --from=builder /app/main /app/main

# アプリケーションのポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["/app/main"]
