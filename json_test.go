package unixtime_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	unixtime "github.com/go-marshaltemabu/go-unixtime"
)

type dataIn struct {
	TargetValue unixtime.UnixTime `json:"target_v"`
}

type dataOut struct {
	TargetInt64 int64 `json:"target_v"`
}

func TestJSON_01(t *testing.T) {
	v := time.Unix(1604139239, 0)
	aux1 := dataIn{
		TargetValue: (unixtime.UnixTime)(v),
	}
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	var aux2 dataOut
	if err = json.Unmarshal(buf, &aux2); nil != err {
		t.Fatalf("json.Unmarshal failed: %v", err)
		return
	}
	if aux2.TargetInt64 != 1604139239 {
		log.Fatalf("cannot transcode JSON: 1604139239 vs. %d", aux2.TargetInt64)
	}
	t.Logf("JSON: %s", string(buf))
	return
}

func TestJSON_02(t *testing.T) {
	var aux1 dataIn
	buf, err := json.Marshal(&aux1)
	if nil != err {
		t.Fatalf("json.Marshal failed: %v", err)
		return
	}
	var aux2 dataOut
	if err = json.Unmarshal(buf, &aux2); nil != err {
		t.Fatalf("json.Unmarshal failed: %v", err)
		return
	}
	if aux2.TargetInt64 != 0 {
		log.Fatalf("cannot transcode JSON: 0 vs. %d", aux2.TargetInt64)
	}
	t.Logf("JSON: %s", string(buf))
	return
}
