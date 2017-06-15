# golang_practice_echo_todo_api
echoを使ったapi
フロント側のテストとしてajax受けるときなど便利に使えそう

## ローカルでの動作

環境変数PORTを読むので、設定しておく。
```bash
$ export PORT=8080
$ go run server.go

# request
$ curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"task":"do hogehoge"}' \
  localhost:8080/tasks

# response
{"id":1,"task":"do hogehoge"}
```
で動作確認


## デプロイ
herokuを想定
Procfileは下記で適宜書き換えする
```bash
echo "web: $(basename `pwd`)" > Procfile
```

## 課題

router.go
models/todo.go
view.go
などに分けたい。

RequestとResponseを表示するhtmlが欲しいが、その場合レンダリング要？
