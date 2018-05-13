//go:generate browserify web/index.js web/js/ws.js -o web/bundle.js
package ui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func (ui *UI) startWeb(ws *WSServer) {
	port := ":13001"
	go func() {
		fs := http.FileServer(http.Dir("/Users/divan/src/github.com/divan/goabm/ui/webgl_grid3d/web"))
		http.Handle("/", noCacheMiddleware(fs))
		http.HandleFunc("/ws", ws.Handle)
		log.Fatal(http.ListenAndServe(port, nil))
	}()
	url := "http://localhost" + port
	time.Sleep(1 * time.Second)
	startBrowser(url)
	for val := range ui.ch {
		data, err := json.Marshal(val)
		if err != nil {
			log.Println(err)
			continue
		}

		ws.broadcastData(data)
	}
}

// startBrowser tries to open the URL in a browser
// and reports whether it succeeds.
//
// Orig. code: golang.org/x/tools/cmd/cover/html.go
func startBrowser(url string) error {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	fmt.Println("If browser window didn't appear, please go to this url:", url)
	return cmd.Start()
}

func noCacheMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, no-cache")
		h.ServeHTTP(w, r)
	})
}
