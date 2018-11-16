package gate

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpGate(t *testing.T) {
	sigChan := make(chan string, 1)

	go func() {
		c := &Config{
			Port:10000,
			WriteWait:20,
			PongWait:60,
			PingPeriod:54,
			MaxMessageSize:512,
			MessageBufferSize:256,
		}
		gate := NewGate(c)
		gate.UseHttp()
		gate.RegisterHttpRouter("/index", func(writer http.ResponseWriter, r *http.Request) {
			fmt.Println("this is a http request")
			sigChan <- "over"
		})
		gate.StartServer()
	}()

	Get("http://localhost:10000/index")
	fmt.Println(<- sigChan)
}

func Get(url string) {
	client := &http.Client{}
	client.Get(url)
}