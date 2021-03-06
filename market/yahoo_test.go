package market

import (
	"os"
	"strings"
	"testing"
	"time"

	"log"

	"github.com/nzai/stockrecorder/config"
	"github.com/nzai/go-utility/io"
)

func init() {

	//	读取配置文件
	err := config.Init()
	if err != nil {
		log.Fatal("读取配置文件错误: ", err)
		return
	}
}

func TestParse60(t *testing.T) {

	var u1 int64 = 1444829400
	t.Logf("%d is %s", u1, time.Unix(u1, 0).Format("2006-01-02 15:04:05"))
}



func TestReplace(t *testing.T) {
	path := `c:\data\America\AAOI\20150826_raw.txt`
	regular := strings.Replace(path, rawSuffix, regularSuffix, -1)
	t.Log(strings.Replace(path, rawSuffix, errorSuffix, -1))
	t.Log(regular)

	_, err := os.Stat(regular)
	if err == nil || os.IsExist(err) {
		t.Log("Exists")
	} else {
		t.Log("Not Exists")
	}
}
