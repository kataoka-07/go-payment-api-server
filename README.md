# go-payment-api-server

法人向け請求支払い支援サービス「スーパー支払い君.com」のバックエンドAPI。

---

## プロジェクト構成

```shell
go-payment-api-server/ 
├── cmd/
│ ├── gentool/          # model 生成
│ └── payment-server/   # サーバアプリ起動
├── internal/
│ ├── di/               # 依存関係の管理
│ ├── domain/
│ │ └── model/          # Entity(自動生成)
│ ├── infrastructure/
│ │ ├── generator/      # model 生成機能設定
│ │ ├── mysql/          # DB 操作
│ │ └── query/          # SQLBuilder(自動生成)
│ ├── usecase/          # ビジネスロジック
│ └── interface/
│   └── handler/        # HTTP Handler
├── pkg/
│ └── logger/           # slogベースのLogger
├── test/
│ └── e2e/              # E2E Test
├── migration/          # マイグレーション
├── docker-compose.yml  # DBコンテナ(Local)
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

## Model 生成方法(Gorm/Gen)

用途: QueryBuilder, Entity の自動生成(DB Schema first)

```bash
make model-gen
```

## E2E Test

```bash
# 準備
make test-db-init
make test-schema-init
make test-seed

# 実行
make e2e-test
```
