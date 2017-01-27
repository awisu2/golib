# golib

## overview
汎用的に使えそうなパッケージ群
たぶん車輪の再開発になっていると思う

## history

### 20170126

- golib/dbをgolib/sqlに変更
- golib/sqlにSessionを追加  
Session経由でdbを開くことで2重オープンを防止できる  
別途sqlを利用するパッケージ側にSession機能を持たせることもできたが、パッケージをまたいで管理させたくなかったため  
 