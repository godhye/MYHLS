package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "127.0.0.1:8000", "http service address")

func main() {

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connect to %s", u.String())

	//client Connect
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("FATAL :", err)

	}
	defer c.Close()

	//값의 교환이 불필요한 경우 빈 struct , size = 0
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("read : %s", err)
				return
			}
			log.Printf("recv: %s", message)

		}
	}()

	//1초마다 타이머
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write :", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt ")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close: ", err)
				return
			}
			select {
			//? 버퍼에 남아있는 경우 위한 처리?
			case <-done:
			//1초마다 지나면 Ready 되는 채널을 리턴
			// 서버와 연결이 종료되기 기다림 (with timeout)
			case <-time.After(time.Second):
			}
			return
		}

	}

}
