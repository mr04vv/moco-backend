# moco-backend

## 環境構築

local に go 無くても動くけど入れたほうが良いです

```
brew install go
```

gvm で go のバージョン管理してもよし

```
docker-compose build
docker-compose up
```

## マイグレーション

```
docker-compose exec go /bin/bash
sql-migrate new add_colums_to_users
# 生成されたsqlファイルにSQLを書く
sql-migrate up
```

## seed

```
docker-compose exec go /bin/bash
cd db/seeds
go run seed.go
```
