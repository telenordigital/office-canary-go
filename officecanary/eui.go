package officecanary

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type EUI [8]byte

// EUIFromString expects a string with the format "xx-xx-xx..." (dashes optional).
func EUIFromString(s string) (EUI, error) {
	b, err := hex.DecodeString(strings.TrimSpace(strings.Replace(s, "-", "", -1)))
	if err != nil {
		return EUI{}, err
	}
	if len(b) != 8 {
		return EUI{}, fmt.Errorf("Invalid EUI: Hex string is not the correct length or format. Expected 8 bytes, got %d", len(b))
	}
	eui := EUI{}
	copy(eui[:], b)
	return eui, nil
}

func (eui EUI) String() string {
	return fmt.Sprintf("%02x-%02x-%02x-%02x-%02x-%02x-%02x-%02x", eui[0], eui[1], eui[2], eui[3], eui[4], eui[5], eui[6], eui[7])
}
