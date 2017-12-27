package hostlink

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

func parseReadResponse(resp string) (string, error) {
	reg := regexp.MustCompile(`^.{7}(.*).{3}$`)
	matches := reg.FindStringSubmatch(resp)

	if len(matches) <= 1 {
		return "", fmt.Errorf("Wrong response format: %s", resp)
	}

	decoded, err := hex.DecodeString(matches[1])
	if err != nil {
		return "", fmt.Errorf("Error during decoding response: %s", resp)
	}

	return string(decoded), nil
}

func parseWriteResponse(resp string) (bool, error) {
	decoded, err := parseReadResponse(resp)
	if err != nil {
		return false, err
	}

	return decoded == "00", nil
}
