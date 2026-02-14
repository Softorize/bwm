package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// BingTime handles Bing's /Date(milliseconds)/ JSON format.
type BingTime struct {
	time.Time
}

func (bt *BingTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		return nil
	}
	// Format: /Date(1234567890000)/
	s = strings.TrimPrefix(s, "/Date(")
	s = strings.TrimSuffix(s, ")/")
	// Handle timezone offset like +0000
	if idx := strings.IndexAny(s, "+-"); idx > 0 {
		s = s[:idx]
	}
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("parsing BingTime %q: %w", string(b), err)
	}
	bt.Time = time.Unix(0, ms*int64(time.Millisecond)).UTC()
	return nil
}

func (bt BingTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"/Date(%d)/"`, bt.UnixMilli())), nil
}

func (bt BingTime) String() string {
	if bt.IsZero() {
		return ""
	}
	return bt.Format("2006-01-02")
}
