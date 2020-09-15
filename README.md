# AtCoder LibaryのGo言語用です

## 1. how to install

以下のコマンドでインストールできます。

```bash
go get github.com/aruaru0/go-atcoder-lib
```

## 2. how to use (local)

インポートすれば使えます

```go
import "github.com/aruaru0/go-atcoder-lib/acl"
```

インポートして使う場合は、

```go
d := acl.NewDsu(n)
```

という形でaclをつけて呼び出します。

## 3. how to use (online judge)
ソースコードのpackage文以下をカットして、ソースコードにコピー＆ペーストします。複数コピペしてもなるべくぶつからないように気を付けて実装する予定です。

## 4. Other

AtCoder Library Practice Contestを通していますが、正しく実装されていない部分があるかもしれません。また、Go言語とC++言語の違いから、一部利用方法や実装が異なりま（もしかしたら遅い部分があるかもしれません）。使い方については、examplを参照ください。

Go言語にはジェネリクスがないので、テンプレートを使った部分については適当な型で実装しています。その他Goでは実装できないものは、適当に妥協して実装していま（とりあえず、全部そろえることが目標です。手伝い歓迎）

### 5. Memo
SCCをC++と同じ実装をするとかなり重いです。まったく違う実装を行ったほうがよいと思われる