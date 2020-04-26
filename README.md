# 環境構築

## go導入

### Windows
- https://golang.org/dl/ からWIndows用のmsiファイルをダウンロードする
- ダウンロードしたmsiファイルを実行し，指示にそってインストールする
- インストールしたbinフォルダにパスを通す
- - プロジェクトはGOPATH/src/github.com の下に作成する
- インストールのテストを行う
```
> go version
go version go1.14.2 windows/amd64
```

## ANTLR導入

### Windows
[参考サイト](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)
- javaをインストールする
- https://www.antlr.org/download/ から antlr-4.7.2-complete.jarをダウンロードする
- ダウンロードしたファイルにパスを通す
- antlr4コマンドが使えるようにバッチファイルを作成してパスを通す

```
java org.antlr.v4.Tool %*
```
- インストールのテストを行う

```
> java org.antlr.v4.Tool
ANTLR Parser Generator Version 4.7.1
-o ___ specify output directory where all output is generated
-lib ___ specify location of .tokens files
...
```

## ANTLRのパーサをgoで出力

- goのruntimeをインストールする
```
go get github.com/antlr/antlr4/runtime/Go/antlr
```
- goのruntimeとantlrのバージョンを合わせる
```
cd ~/go/src/github.com/antlr/antlr4
git checkout tags/4.7.2
```

- 出力言語を指定してパーサを生成する
```
antlr4 -Dlanguage=Go -o ./generated example.g4
```
[antlrのオプション一覧](https://github.com/antlr/antlr4/blob/master/doc/tool-options.md)
