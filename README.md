# go学習記録

## 概要

goを学習する上で記録をこのREADMEに記載していく。
記載方法は自由

# Day1(2021/08/24)
* goは不要な宣言を許可しないため、参照されていないパッケージがあるとコンパイルエラーが起こる。
  * 後に使用するから記載して置きたいなどの理由がある場合は、パッケージ名の左側に_を補うとコンパイルエラーが怒らない

* エントリーポイントはmainパッケージの中に定義された関数mainであると定められている。
* プログラムのbuildは以下のコマンドで可能
```sh
go build -o [出力ファイル名] [build対象goファイル]
```

上記について以下のコマンドでもbuildすることが可能
このコマンドを実行するとカレントディレクトリにあるメイン関数があるgoファイルがbuildされる。

```sh
go build
```

# Day2(2021/08/25)
