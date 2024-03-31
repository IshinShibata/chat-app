# Go言語の公式イメージをベースとする
FROM golang:1.22.1 AS builder

# アプリケーションのソースコードを含むディレクトリを作成
WORKDIR /app

# モジュールファイルをコピー
COPY go.mod go.sum ./

RUN go mod download && go install github.com/cosmtrek/air@v1.44.0
