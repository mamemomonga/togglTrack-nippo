# toggl trackで日報

toggle trackの特定Projectの日報をSlackの特定チャンネルに投げるツールです。

* Slack Appの設定とTokenの取得が必要です。
* Toggl TrackのTokenが必要です。
* あまり古いタスクは取得できません。

# 設定ファイル

config.yaml

	togglWorkspace:
	  token: 00000000000000000000000000000000
	  workspace: "mamemomonga's Workspace"

	tasks:
	  - slack:
	      token: "xoxp-000000000000-000000000000-0000000000000-00000000000000000000000000000000"
	      channel: "#channelName1"
	    toggl:
	      - client: "お客様1"
	        project: "なにかの作業A"
	      - client: "お客様1"
	        project: "なにかの作業B"
	  - slack:
	      token: "xoxp-000000000000-000000000000-0000000000000-00000000000000000000000000000000"
	      channel: "#channelName2"
	    toggl:
	      - client: "お客様2"
	        project: "なにかの作業B"
	      - client: "お客様2"
	        project: "なにかの作業B"

# 実行例

ヘルプ

	./toggl-nippo --help

今日の日報を取得する

	./toggl-nippo -config ./config.yaml

今日の日報を指定Slackチャンネルに送る

	./toggl-nippo -config ./config.yaml -slack

昨日の日報を指定Slackチャンネルに送る

	./toggl-nippo -config ./config.yaml -days 1 -slack

# License

MIT
