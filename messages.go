package api

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

type ConsumeOut struct {
	NextOffset int64               `json:"next_offset"`
	Encoding   string              `json:"encoding"`
	Messages   []ConsumeMessageOut `json:"messages,omitempty"`
}

type ConsumeMessageOut struct {
	Offset int64     `json:"offset"`
	Time   time.Time `json:"time"`
	Key    *string   `json:"key,omitempty"`
	Value  *string   `json:"value,omitempty"`
}

type PublishIn struct {
	Encoding string             `json:"encoding"`
	Messages []PublishMessageIn `json:"messages"`
}

type PublishMessageIn struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type PublishOut struct {
	NextOffset int64 `json:"next_offset"`
}

type GetOut struct {
	Encoding string    `json:"encoding"`
	Offset   int64     `json:"offset"`
	Time     time.Time `json:"time"`
	Key      *string   `json:"key,omitempty"`
	Value    *string   `json:"value,omitempty"`
}

type PostIn struct {
	Encoding string  `json:"encoding"`
	Key      *string `json:"key"`
	Value    *string `json:"value"`
}

type PostOut struct {
	NextOffset int64 `json:"next_offset"`
}

func (k *Client) PublishK(logID ksuid.KSUID, keys []string) error {
	in := PublishIn{}
	for _, k := range keys {
		k := k
		in.Messages = append(in.Messages, PublishMessageIn{
			Key: &k,
		})
	}
	var out PublishOut
	return k.base.Post(fmt.Sprintf("messages/%s", logID), in, &out)
}

func (k *Client) PublishKV(logID ksuid.KSUID, keys []string, values []string) error {
	in := PublishIn{}
	for i, k := range keys {
		k := k
		v := values[i]
		in.Messages = append(in.Messages, PublishMessageIn{
			Key:   &k,
			Value: &v,
		})
	}
	var out PublishOut
	return k.base.Post(fmt.Sprintf("messages/%s", logID), in, &out)
}

func (k *Client) PublishV(logID ksuid.KSUID, values []string) error {
	in := PublishIn{}
	for _, v := range values {
		v := v
		in.Messages = append(in.Messages, PublishMessageIn{
			Value: &v,
		})
	}
	var out PublishOut
	return k.base.Post(fmt.Sprintf("messages/%s", logID), in, &out)
}

func (k *Client) Consume(logID ksuid.KSUID, offset int64, sz int32) (ConsumeOut, error) {
	var out ConsumeOut
	err := k.base.Get(fmt.Sprintf("messages/%s?offset=%d&len=%d", logID, offset, sz), &out)
	return out, err
}

func (k *Client) GetMessage(logID ksuid.KSUID, offset int64) (GetOut, error) {
	var out GetOut
	err := k.base.Get(fmt.Sprintf("message/%s?offset=%d", logID, offset), &out)
	return out, err
}

func (k *Client) PostMessage(logID ksuid.KSUID, key string) error {
	in := PostIn{Key: &key}
	var out PostOut
	return k.base.Post(fmt.Sprintf("message/%s", logID), in, &out)
}
