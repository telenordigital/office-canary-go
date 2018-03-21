package officecanary

import (
	"log"
	"testing"
)

func TestStreamDatapoints(t *testing.T) {
	oc, err := NewClient("YOUR API TOKEN HERE")
	if err != nil {
		t.Fatal(err)
	}

	err = oc.StreamDatapoints(func(d Datapoint) {
		log.Println(d)
	})
	if err != nil {
		t.Fatal(err)
	}
}
