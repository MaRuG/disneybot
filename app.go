package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
        "xxxxxxxxxxxxxxxxxxxxxxxxxx",
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch msg := event.Message.(type) {
				case *linebot.TextMessage:
					if msg.Text == "test" {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("success")).Do()
					} else if msg.Text == "ビッグサンダーマウンテン" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#WL_BTMOUNTAIN .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "スプラッシュマウンテン" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#CC_SPLASH .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "プーさんのハニーハント" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_POOH .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ホーンテッドマンション" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_HAUNTED .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "スターツアーズ" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TL_TOURS .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "バズライトイヤー" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TL_BUZZ .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "モンスターズインク" {
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TL_MONSTER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ウエスタンリバー鉄道"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AL_WESTERN .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "カリブの海賊"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AL_CARIB .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ジャングルクルーズ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AL_JUNGLE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ツリーハウス"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AL_SWISS .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "魅惑のチキルーム"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AL_TIKI .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						} 
					} else if msg.Text == "ガジェットのゴーコースター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_GADGET .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "グーフィー"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_GOOFY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "チップとデール"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_CHIP .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トゥ―ンパーク"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_TOON .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ドナルドのボート"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_DONALD .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ミニーの家"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_MINNIE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ロジャーラビット"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TT_RABBIT .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "スティッチエンカウンター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#TL_STITCH .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "シューティングギャラリー"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#WL_GALLERY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "カントリーベアシアター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#WL_COUNTRY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トムソーヤ島いかだ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#WL_TOM .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "マークトウェイン号"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#WL_MARK .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "カヌー探検"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#CC_BEAVER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "スモールワールド"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_SMALL .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "アリスのティーパーティー"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_ALICE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "キャッスルカルーセル"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_CASTLE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "シンデレラ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_CINDERELLA .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ピノキオ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_PINOCCHIO .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ピーターパン"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_PETER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "フィルハーマジック"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_PHILHAR .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "白雪姫"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_SNOW .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "空飛ぶダンボ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#FL_DUMBO .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トイストーリーマニア"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_TOY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "インディジョーンズ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#LD_INDIANA .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "センターオブジアース"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#MI_CENTER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "タワーオブテラー"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_TOT .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トランジットアメリカン発"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_TRANSIT .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トランジットロスト発"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#LD_TRANSIT .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "フライングフィッシュコースター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_FLOUNDER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "バルーンレース"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_BALLOON .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ワールプール"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_POOL .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ニモ&フレンズ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#PD_SEARIDER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "マジックランプシアター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AC_MAGIC .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "タートルトーク"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_TURTLE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "フライングカーペット"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AC_JASMINE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ジェリーフィッシュ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_JELLY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "海底2万マイル"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#MI_20000 .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "キャラバンカルーセル"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AC_CARAVAN .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "シンドバット"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AC_SINDBAD .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "スカットルのスクーター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_SCUTTLE .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ヴェネツィアンゴンドラ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#MH_GONDOLAS .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ビッグシティヴィークル"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_BIG .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "マーメイドラグーンシアター"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_THEATER .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "レイジングスピリッツ"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#LD_RAGING .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "アリエル"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#ML_PLAY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "フォートレス"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#MH_FORTRESS .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "アクアトピア"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#PD_AQUA .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "レールウェイアメリカ発"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#AW_RAILWAY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "レールウェイポート発"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#PD_RAILWAY .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "トランジットメディテ発"{
						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						test := doc.Find("#MH_TRANSIT .area p").Text()
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(test)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ランド"{
						var all []string

						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneyland-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						doc.Find(".wait-time .area").Each(func(i int, s *goquery.Selection) {
							about := s.Find(".about > h4").Text()
							fp := s.Find(".about > .fp").Text()
							time := s.Find(".time > p").Text()
							act := "Name: " + about + "\n" + "Time: " + time + "\n" + "FP: " + fp + "\n" + "----"
							all = append(all, act)
							// fmt.Println(app)
						})
						re_all := strings.Join(all, "\n")
						
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(re_all)).Do()
						if err != nil {
							log.Fatal(err)
						}
					}  else if msg.Text == "シー"{
						var all []string

						doc, err := goquery.NewDocument("http://disneyreal.asumirai.info/realtime/disneysea-wait-today.html")
						if err != nil {
							log.Fatal(err)
						}
						doc.Find(".wait-time .area").Each(func(i int, s *goquery.Selection) {
							about := s.Find(".about > h4").Text()
							fp := s.Find(".about > .fp").Text()
							time := s.Find(".time > p").Text()
							act := "Name: " + about + "\n" + "Time: " + time + "\n" + "FP: " + fp + "\n" + "----"
							all = append(all, act)
							// fmt.Println(app)
						})
						re_all := strings.Join(all, "\n")
						
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(re_all)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "ユニバ" {
						var all []string

						doc, err := goquery.NewDocument("http://usjinfo.com/wait/realtime.php")
						if err != nil {
							log.Fatal(err)
						}
					
						doc.Find("#contents table tr").Each(func(i int, s *goquery.Selection) {
							name := s.Find("td:nth-child(1) a  b").Text()
							time := s.Find("td:nth-child(2) b").Text()
							time = strings.TrimSpace(time)
							act := "Name: " + name + "\n" + "Time: " + time + "\n" + "----"
							all = append(all, act)
						})
						re_all := strings.Join(all, "\n")
						//fmt.Println(re_all)
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(re_all)).Do()
						
						if err != nil {
							log.Fatal(err)
						}

					}else {
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("間違っています!困ったら'ランド'か'シー'ッて打ちな！")).Do()
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	})
	
	if err := http.ListenAndServe(":2415", nil); err != nil {
        log.Fatal(err)
    }
}
