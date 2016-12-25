package store

import (
    "testing"
    "encoding/json"
    . "github.com/smartystreets/goconvey/convey"
)

func TestDump(t *testing.T) {

    Convey("dump%load data", t, func() {
        m := map[string]interface{} {
            "xxx": "111dfa",
            "yyyyfdasfd": "aaaaa",
        }
        Dump(m)
        d_str, _ := json.Marshal(Load())
        e_str, _ := json.Marshal(m)
        So(string(d_str), ShouldEqual, string(e_str))
    })
}
