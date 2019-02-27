# スターティングGo言語

## 実行
### コンパイルなし(ビルド&実行)
`go run main.go`

### コンパイルあり
`go build -o main main.go`: 実行ファイル生成

`./main`

## goのコマンド
- `go fmt`: コードのフォーマット
  - `-n`: 実行されるコマンドの表示(ファイルの書き換えなし)
  - `-x`: 実行されるコマンドの表示
- `go env`: 環境変数の確認
- `go doc`: goのパッケージのドキュメントを参照
