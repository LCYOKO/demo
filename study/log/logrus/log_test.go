package logrus

// https://www.liwenzhou.com/posts/Go/logrus/
// https://github.com/sirupsen/logrus
import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestInfo(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
}
