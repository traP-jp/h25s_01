# server

## Tasks

開発に用いるコマンド一覧、`xc build`のように実行できる

`xc`のインストールが必要

```bash
go install github.com/joerdav/xc/cmd/xc@latest
```

### build

アプリをビルドします。

```sh
CMD=server
go mod download
go build -o ./bin/${CMD} ./cmd/${CMD}
```

### Lint

Linter (golangci-lint) を実行します。

```sh
golangci-lint run --timeout=5m --fix ./...
```
