package helper

import (
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	t.Run("stop watch", func(t *testing.T) {
		{
			defer Elapsed("1 s", true)()
			time.Sleep(1 * time.Second)
		}
	})
}