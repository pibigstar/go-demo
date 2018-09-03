package cygo

import (
	"time"
)

type CyTime struct {
	time.Time
}

func Now() CyTime {
	t := CyTime{}
	t.Time = time.Now()
	return t
}

func (t CyTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"` + DATE_LAYOUT_SOLR_MINUTE + `"`)), nil
}

func (t *CyTime) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse(`"`+time.RFC3339+`"`, string(data))
	return
}

const DATE_LAYOUT_SOLR_MINUTE = "2006-01-02T15:04:00Z"
const DATE_LAYOUT_CHINA = "2006年1月2日 15点4分"
const DATE_DAY = "2006-01-02"
