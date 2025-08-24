# dev: for development environment with hot-reloading
FROM golang:1.24-bookworm AS dev
WORKDIR /app  # 以降の作業ディレクトリをに設定
# hot-reloadingツールをインストール。ファイル更新を検知するたびに自動で再ビルド・再起動してくれる
RUN go install github.com/cosmtrek/air@latest 
CMD ["air"]

# build: to build the Go binary
# 別ステージ: クリーンな環境で本番用バイナリを作成する工場
FROM golang:1.24-bookworm AS builder

WORKDIR /src # build用の作業ディレクトリ
COPY go.mod go.sum ./ # 依存関係の定義のみ先にコピー
RUN go mod download # 依存関係のインストール

COPY . . # 残りのソースコードをコピー
# 最適化ビルド:デバッグ情報省略(-s -w)、出力を/out/appに。mainが./cmd/appにある想定
RUN go build -ldflags="-s -w" -o /out/app ./cmd/app


# runtime: to run the Go binary
# 実行専用の最終ステージ: シェル等がない最小かつ高セキュリティなランタイム
FROM gcr.io/distroless/static-debian12 AS DEPLOYMENT
COPY --from=builder /out/app /app # builderステージからビルド済バイナリのみをコピー
USER nonroot:nonroot # root権限を避ける本番セキュリティの基本
ENTRYPOINT ["/app"] # コンテナ起動時に実行されるコマンド