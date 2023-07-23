package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		atomic.AddUint64(&number, 1)
		//m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa pagina %d vezes.", number)))
	})
	http.ListenAndServe(":3000", nil)
}
