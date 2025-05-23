# birthinfo

`birthinfo` は、あなたの生年月日から曜日を計算するシンプルなGoツールです。
ただのリポジトリ数稼ぎで作りました。

## 概要

このツールを使うことで、指定した生年月日からその曜日を簡単に取得できます。
例えば、`2007/01/19`の曜日を求めることができます。

## インストール方法

Goがインストールされている環境で以下の手順を実行してください。

1. リポジトリをクローンします：

```zsh
git clone https://github.com/raiton-boo/birthinfo.git
cd birthinfo
```

2. 必要な依存関係をインストールします（`go mod`が自動で管理します）：

```zsh
go mod tidy
```

## 使い方

`main.go` を実行するだけで、指定した生年月日から曜日が計算されます。

1. `main.go` を開き、生年月日を設定します：

```go
year := 2007
month := 1
day := 19
```

2. 以下のコマンドで実行します：

```zsh
go run main.go
```

結果が表示されます：

```
2007/01/19 あなたの産まれた日は Friday です
```

## ライセンス

MIT License - 詳しくは [LICENSE](LICENSE) ファイルを参照してください。