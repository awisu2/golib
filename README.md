# golib

## overview
汎用的に使えそうなパッケージ群
たぶん車輪の再開発になっていると思う

## history

### 20170131

- commandにInt取得の関数を追加
- shiftjisをutf8に変換する関数の追加
- fileに、利用の際のパラメータコメント追加
- fileに削除関数追加
- sql型にbigint追加

### 20170130

- 文字列数値を日本語数値に変換する機能追加
- mpa[string]interface{}をとりあえず宣言
- os.Argsがあんまり行けてる感じに使えてなかったので、パッケージだけ作成
- logがExtraStringを出力していたので修正
- sql/tableにDBConfigNameを追加

### 20170128

- 【追加】type MapStringString map[string]string

### 20170127

- golib/sqlにtableInfosを追加  
一度gositeにセットしたtable.infoを再セットする処理があったが、直観的でないので、  
golib/sqlが直で持つように変更  
(パッケージ直のインスタンスなのはどうかと思うので、用途に合わせた区分と持ち方を思いついたらそちらに変更)  
- NewSessionでConfigインスタンスを取得するように修正

### 20170126

- golib/dbをgolib/sqlに変更
- golib/sqlにSessionを追加  
Session経由でdbを開くことで2重オープンを防止できる  
別途sqlを利用するパッケージ側にSession機能を持たせることもできたが、パッケージをまたいで管理させたくなかったため  
 