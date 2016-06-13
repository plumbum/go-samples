package main

import (
	"fmt"
	"time"

	"github.com/uber-go/zap"
	"labix.org/v2/mgo"
	"encoding/json"
)


type WriterWrapper struct {
	coll *mgo.Collection
}

func (w WriterWrapper) Write(p []byte) (n int, err error) {
	// fmt.Println(string(p))

	out := make(map[string]interface{})
	err = json.Unmarshal(p, &out)
	if err != nil {
		n = 0
		return
	}
	err = w.coll.Insert(out)
	if err != nil {
		n = 0
		return
	}
	n = len(p)
	return
}

func (w WriterWrapper) Sync() error {
	return nil
}

func NewWriter(session *mgo.Session) *WriterWrapper {
	wrap := new(WriterWrapper)
	wrap.coll = session.DB("logger").C("testlog")
	return wrap
}

const logCounts = 1e5

func main() {
	mongoSession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	mongoSession.SetMode(mgo.Monotonic, true)

	log := zap.NewJSON(
		zap.Debug,
		zap.Fields(zap.Int("count", 1)),
		zap.Output(NewWriter(mongoSession)),
	)
	url := "http://example.local"
	tryNum := 42
	startTime := time.Now()
	for i := range [logCounts]struct{}{} {
		log.Info("Failed to fetch URL.",
			zap.String("url", url),
			zap.Int("attempt", tryNum),
			zap.Duration("backoff", time.Since(startTime)),
			zap.Int("index", i),
		)
	}
	fmt.Printf("Finished in %v\n", time.Since(startTime))
}
