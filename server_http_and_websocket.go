package main

import (
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"

	"html/template"
)

const listenAddr = "localhost:4000"

var htmlTemplate = template.Must(template.New("root").Parse(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Golang - Http + WebSocket example</title>
			<meta charset="utf-8"/>
			<script type="text/javascript">
				var ws = new WebSocket("ws://{{.}}/socket");
				ws.onmessage = onMessage;
				ws.onclose = onClose;

				function onMessage(msg) {
					console.log("Received: ", msg)
				}

				function onClose() {
					console.log("Websocket connection closed");
				}
			</script>
		</head>
	</html>
`))

func httpHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate.Execute(w, listenAddr)
}

func socketHandler(conn *websocket.Conn) {
	var s string

	fmt.Fscan(conn, &s)
	fmt.Println("Received: ", s)
	fmt.Fprint(conn, "How do you do how how?")
}

func main() {
	http.HandleFunc("/", httpHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))

	err := http.ListenAndServe(listenAddr, nil)

	if err != nil {
		log.Fatal(err)
	}
}