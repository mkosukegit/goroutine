FROM golang:1.17.6-alpine

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/cmd

CMD ["go", "run", "/go/src/app/main.go"]