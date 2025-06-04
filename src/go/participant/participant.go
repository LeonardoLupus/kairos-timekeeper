package participant

import (
	"time"
)

type Participant struct {
	ID       uint64
	Username string
	UTC      time.Time
}
