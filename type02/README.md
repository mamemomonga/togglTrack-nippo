# toggl trackで日報 Type02

当日の作業開始時刻,終了時刻,休憩時間を表示する

# 使い方

* -days 0 は当日(デフォルト)、-days 1 で昨日。
* デバッグログはSTDERRに出るので、邪魔だったら 2>/dev/null などで捨てる。

実行例

    $ cp config.example.yaml config.yaml
    $ vim config.yaml
    $ ./toggl-nippo-type02 -config ./config.yaml -days 0 2>/dev/null
