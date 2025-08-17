package logx

import (
	"log"
)

func Init() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("[poller] ")
}
