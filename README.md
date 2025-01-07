# Go-Docker-CRUD アプリケーション

このプロジェクトは、効率的なリソース管理のために設計された Go ベースの CRUD API です。Docker を使用したコンテナ化と PostgreSQL をデータベースバックエンドとして活用しています。本システムは、業界のベストプラクティスに従い、拡張性、保守性、展開の容易さを重視して構築されています。


# 特徴

## コンテナ化されたアプリケーション:

- マルチステージ Docker ビルドを使用した軽量でセキュアなプロダクションイメージ。
- 最新の glibc 要件を満たし、ポータビリティを最適化。

## クリーンアーキテクチャ:

- モジュール設計による関心の分離: Handler、Service、および Repository レイヤー。
- SOLID 原則に従った高いテスト可能性を備えたコード構造。

## データベース統合:

- PostgreSQL をリレーショナルデータストレージとして採用。
- docker-compose を使用した自動データベース初期化。

## 拡張可能な設計:

- 高速なビルドと展開を可能にする最適化された Docker イメージレイヤー。
- Kubernetes などのツールを使用したオーケストレーションに対応。


# 前提条件

## 以下のソフトウェアがシステムにインストールされていることを確認してください:

- Docker >= `24.0.0`
- Docker Compose >= `2.0.0`


# 開始手順

## リポジトリのクローン

- `$ git clone https://github.com/ThePrettyN/go-docker-crud.git`
- `$ cd go-docker-crud`

## アプリケーションのビルドと実行

1. Docker Compose を使用してアプリケーションをビルドおよび起動:

- `$ docker-compose up --build`

2. アプリケーションは http://localhost:8080 でアクセス可能。

3. PostgreSQL データベースが以下の認証情報で起動していることを確認:

- Username: `postgres`
- Password: `password`

## API エンドポイント

- Base URL: http://localhost:8080
  ```
  メソッド             エンドポイント               説明
  GET                 /resources                  全リソースの取得
  GET                 /resources/:id              ID に基づくリソースの取得
  POST                /resources                  新規リソースの作成
  PUT                 /resources/:id              既存リソースの更新
  DELETE              /resources/:id              リソースの削除
  ```
## アプリケーションの停止

アプリケーションを停止するには以下を実行:

- `$ docker-compose down`


# プロジェクト構成
```
📁 go-docker-crud
├── 📂 cmd
│   └── main.go              # アプリケーションエントリポイント
├── 📂 internal
│   ├── 📂 handler           # API エンドポイントの HTTP ハンドラー
│   ├── 📂 service           # ビジネスロジック層
│   └── 📂 repository        # データベース操作層
├── Dockerfile               # Docker ビルド設定
├── docker-compose.yml       # Docker Compose 設定
├── go.mod                   # Go モジュール依存関係
├── go.sum                   # 依存関係チェックサム
└── README.md                # プロジェクトドキュメント
```

# アーキテクチャの説明

## Handler:

- HTTP リクエストとレスポンスを管理。
- データ検証とルーティングを担当。

## Service:

- ビジネスロジックを含み、ハンドラーとリポジトリ間のやり取りを調整。

## Repository:

- データベースとの直接的な操作を担当。
- gorm を使用したスムーズなデータベース統合。


# 主要技術

## Go:

- 高速で静的型付けされた言語で、優れた並行処理のサポートがあります。
- go modによる依存関係管理でモジュール化されたコードを実現。

## Docker:

- 最適化されたイメージのためのマルチステージビルド。
- デプロイメントを簡素化し、一貫した環境を保証します。

## PostgreSQL:

- 強力で信頼性の高いオープンソースのリレーショナルデータベース。
- 永続的なデータ保存に使用されます。


# 高度なトピック

## Kubernetes を使ったスケーリング

1. **Kubernetes ディレクトリを作成し、以下を定義:**

- Deployments
- Services
- ConfigMaps/Secrets

2. **以下のコマンドでアプリケーションをデプロイ:**

- $ kubectl apply -f <manifest-file>

## CI/CD 統合

- GitHub Actions を使用した自動テストとデプロイ。
- Docker Hub または AWS ECR を活用したコンテナイメージの保存。


# テスト (Postman 例)

以下は API エンドポイントをテストするための具体例です。curl コマンドを使用してアプリケーションの動作を検証できます。

1. **ユーザーの作成 (エンドポイント: POST /users)**

- 説明: 新しいユーザーをデータベースに作成します。
  ```
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 30
    }'
  ```
- 期待されるレスポンス (HTTP 201 Created):
  ```
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30
  }
  ```
2. **全ユーザーの取得 (エンドポイント: GET /users)**

- 説明: データベース内の全ユーザー情報を取得します。
  ```
  curl http://localhost:8080/users
  ```
- 期待されるレスポンス (HTTP 200 OK):
  ```
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 30
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "email": "jane.smith@example.com",
      "age": 25
    }
  ]
  ```
3. **特定のユーザーの取得 (エンドポイント: GET /users/{id})**

- 説明: 指定した ID のユーザー情報を取得します。
  ```
  curl http://localhost:8080/users/1
  ```
- 期待されるレスポンス (HTTP 200 OK):
  ```
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30
  }
  ```
- ユーザーが存在しない場合のレスポンス (HTTP 404 Not Found):
  ```
  User not found (404 Not Found status code).
  ```
4. **ユーザー情報の更新 (エンドポイント: PUT /users/{id})**

- 説明: 指定した ID のユーザー情報 (名前、メールアドレス、年齢) を更新します。
  ```
  curl -X PUT http://localhost:8080/users/1 \
    -H "Content-Type: application/json" \
    -d '{
      "name": "John Updated",
      "email": "john.updated@example.com",
      "age": 31
    }'
  ```
- 期待されるレスポンス (HTTP 200 OK):
  ```
  {
    "id": 1,
    "name": "John Updated",
    "email": "john.updated@example.com",
    "age": 31
  }
  ```
5. **ユーザーの削除 (エンドポイント: DELETE /users/{id})**

- 説明: 指定した ID のユーザーをデータベースから削除します。
  ```
  curl -X DELETE http://localhost:8080/users/1
  ```
- 期待されるレスポンス: HTTP 204 No Content または類似の成功メッセージ。