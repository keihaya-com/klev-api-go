package messages

import (
	"github.com/klev-dev/klev-api-go/client"
)

const (
	// OffsetOldest represents the smallest offset still available
	// Use it to consume all messages, starting at the beginning of the log
	OffsetOldest int64 = -2
	// OffsetNewest represents the offset that will be used for the next produce
	// Use it to consume messages, starting from the next one produced
	OffsetNewest int64 = -1
	// OffsetInvalid is the offset returned when error is detected
	OffsetInvalid int64 = -3
)

type Client struct {
	H client.HTTP
}
