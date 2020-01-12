package pg

import (
	"time"

	"github.com/lib/pq"
)

func ToNullTime(t time.Time) pq.NullTime {
	return pq.NullTime{
		Time:  t,
		Valid: true,
	}
}

func NullTime() pq.NullTime {
	return pq.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
}
