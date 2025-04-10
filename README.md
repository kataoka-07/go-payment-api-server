# go-payment-api-server

法人向け請求支払い支援サービス「スーパー支払い君.com」のバックエンドAPI。

---

## プロジェクト構成

```shell
go-payment-api-server/ 
├── cmd/
│ └── payment-server/   # サーバアプリ起動
├── internal/
│ ├── di/               # 依存関係の管理
│ ├── domain/
│ │ └── model/          # Entity
│ ├── infrastructure/
│ │ └── mysql/          # DB 操作
│ ├── usecase/          # ビジネスロジック
│ └── interface/
│   └── handler/        # HTTP Handler
├── pkg/
│ └── logger/           # slogベースのLogger
├── migration/          # マイグレーション
├── docker-compose.yml  # 実行コンテナ環境(Local)
└── go.mod
```

## 起動方法（Docker DB）

動作環境: MacOS Intel Chip

```bash
# 初回起動（MySQL + API）
make up
```

## Local 起動方法(go server)

```bash
make run
```
