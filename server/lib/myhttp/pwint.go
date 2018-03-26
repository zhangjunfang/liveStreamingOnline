package myhttp

import (

	"fmt"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"io"
	"encoding/base64"
	"crypto/rand"
	"log"
	"github.com/zhangjunfang/liveStreamingOnline/config"
	"github.com/zhangjunfang/liveStreamingOnline/server/lib/mywebsocket"
	"math/big"
	"os"
)
var username string="";
var uzb string = "";

var logfile, err = os.OpenFile(config.ServerLog, os.O_RDWR|os.O_CREATE, 0666)
var logger = log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

type userinfo struct {
	name   string
	img    string
	isself bool
}
type Client struct {
	id   string
	conn *mywebsocket.Conn
	userinfo
}

type message struct {
	Data  string
	Mtype string
	Img   string
}

func getclient(ws *mywebsocket.Conn) string {
	for k, v := range member {
		if v.conn == ws {
			return k
		}
	}
	return ""
}
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func guid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func (m *Client) addclient(ws *mywebsocket.Conn) *Client {
	m.conn = ws
	return m
}
func getnun() string {
	rnd, _ := rand.Int(rand.Reader, big.NewInt(12))
	num := fmt.Sprintf("%v", rnd)
	return num
}

var member = make(map[string]*Client)




func Pwint(ws *mywebsocket.Conn) {
	defer func() {
		ws.Close()
	}()
	uid := guid()
	//logger.Println(i)
	if username == "" && uzb == "女主播" {
		username = uzb
	}
	user := userinfo{fmt.Sprintf("%s：", username), fmt.Sprintf("/public/images/%s.jpg", getnun()), false}
	username = ""
	client := Client{uid, ws, user}
	client.addclient(ws)
	member[uid] = &client
	for {
		var err error
		var reply string
		if err = mywebsocket.Message.Receive(ws, &reply); err != nil {
			logger.Println("LiveGoServer:", err)
			break
		}
		for k, v := range member {
			if v.conn != ws {
				var mymes message
				json.Unmarshal([]byte(reply), &mymes)
				if mymes.Mtype == "mess" {
					mymes.Data = fmt.Sprintf("%s%s", member[getclient(ws)].userinfo.name, mymes.Data)
					mymes.Img = member[getclient(ws)].userinfo.img
				}
				msg, _ := json.Marshal(mymes)
				if err = mywebsocket.Message.Send(v.conn, string(msg)); err != nil {
					delete(member, k)
					logger.Println("LiveGoServer:", err)
					break
				}
			} else {
				if !v.userinfo.isself {
					var mymesss message
					json.Unmarshal([]byte(reply), &mymesss)
					//if mymesss.Mtype == "mess" {
					mymesss.Mtype = "self"
					mymesss.Img = v.userinfo.img
					msg, _ := json.Marshal(mymesss)
					member[k].userinfo.isself = true
					if err = mywebsocket.Message.Send(ws, string(msg)); err != nil {
						delete(member, k)
						logger.Println("LiveGoServer:", err)
						break
					}
					//}
				}
			}

		}
	}
}

