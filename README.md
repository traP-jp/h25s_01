# 25 春ハッカソン 01 班

## 同時起動

```bash
docker compose up -d
```

ポート 8080 で Web サーバーが起動し、`/api/*`を server、その他を client に流す。
8081 番で adminer、3306 番で MySQL が起動する。
