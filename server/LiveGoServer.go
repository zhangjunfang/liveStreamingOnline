package main

import (

	"github.com/zhangjunfang/liveStreamingOnline/config"

	"net/http"
)
import (
	"github.com/zhangjunfang/liveStreamingOnline/server/lib/myhttp"

	"github.com/zhangjunfang/liveStreamingOnline/server/lib/mywebsocket"
)


func main() {

	go func() {
		http.Handle("/chat", mywebsocket.Handler(myhttp.Pwint))
	}()

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/live", myhttp.Live)

	http.HandleFunc("/camera", myhttp.Camera)

	http.Handle("/", http.RedirectHandler("/index", 301))

	http.HandleFunc("/index", myhttp.Index)

	var config = config.ServerHost + ":" + config.ServerPort

	if err := http.ListenAndServe(config, nil); err != nil {
		logger.Println("LiveGoServer:", err)
		logfile.Close()
	}

}
