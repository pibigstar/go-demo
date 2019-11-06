package retry

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func demo() error {
	i := rand.Int31n(5)
	if i == 2 {
		return NoRetryError(fmt.Errorf("i is 3"))
	}
	return fmt.Errorf("i = %d", i)
}

func TestRetry(t *testing.T) {
	err := Retry(3, time.Second, demo)
	if err != nil {
		t.Log(err)
	}
}
