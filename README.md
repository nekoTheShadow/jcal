# jcal

自分用calコマンド、略してjal

# Usage

```
Usage: jcal [year month]
  - year : 1955..2021
  - month: 1..12
```

# Install

https://github.com/nekoTheShadow/jcal/releases

# 利用イメージ

引数なしで実行すると、実行した月のカレンダーが表示されます。

![screenshot1.PNG]](https://raw.githubusercontent.com/nekoTheShadow/jcal/master/screenshot1.png)

引数として年と月を指定すると、指定した月のカレンダーが表示されます。

![screenshot2.PNG]](https://raw.githubusercontent.com/nekoTheShadow/jcal/master/screenshot2.png)

# 特徴と制約

- 月単位のカレンダーが日曜始まりで表示されます。
- 土曜日は青文字、日曜日と国民の祝日は赤色で表示されます。
    - 国民の祝日データは[内閣府Webページ](https://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html)より取得しています
    - 国民の祝日をサポートする関係上、1955年から2021年までサポートしています。
- 国民の祝日がある月の場合は、どの祝日であるかも表示しします。