package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateId(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Millisecond)
		id := GenID()
		idStr := fmt.Sprintf(`%v`, id)

		// Generate and print, all in one.
		fmt.Printf("id: %d, len:%v\n", id, len(idStr))
	}

}
