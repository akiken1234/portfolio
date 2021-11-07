＃portfolio

# サービス名
(各種バッジ)

（サービスの画面画像）

（サービスの概要や機能を書く）

## ネットワーク構成図
（prod環境までできたら図を作って記載）

## 環境
|環境|サーバー|URL|
|:---:|:---|:---|
|local|ローカルPC|http://127.0.0.1:3000|
|prod|AWS|(URL)|

## 言語・設計・フレームワーク
- GoLang v1.16
- gin

## 開発環境構築手順
```
① portfolio配下で docker　compose up -d を実行。
②　　portfolio配下で docker　compose exec front sh を実行し、npm install　を実行。
③ portfolio配下で docker　compose exec go sh を実行し、go run main.go　を実行。
④ http://localhost:3000 にアクセス。
ログインする場合、あらかじめ入力されているメールアドレスとパスワードをお使いください。
```

## テスト
go test ./...
