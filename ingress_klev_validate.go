package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/klev-dev/kleverr"
)

var retMessage = kleverr.Ret1[ConsumeMessage]

func IngressWebhookKlevValidateMessage(w http.ResponseWriter, r *http.Request, now func() time.Time, secret string) (ConsumeMessage, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return retMessage(ErrKlevInvalidContentTypeJson())
	}

	payload, err := ingressWebhookKlevValidate(w, r, now, secret)
	if err != nil {
		return retMessage(err)
	}

	var out GetOut
	if err := json.Unmarshal(payload, &out); err != nil {
		return retMessage(err)
	}
	msg, err := out.Decode()
	if err != nil {
		return retMessage(err)
	}
	return ConsumeMessage{
		Offset: msg.Offset,
		Time:   msg.Time,
		Key:    msg.Key,
		Value:  msg.Value,
	}, nil
}

func IngressWebhookKlevValidateData(w http.ResponseWriter, r *http.Request, now func() time.Time, secret string) ([]byte, error) {
	if r.Header.Get("Content-Type") != "application/octet-stream" {
		return nil, ErrKlevInvalidContentTypeOctet()
	}
	return ingressWebhookKlevValidate(w, r, now, secret)
}

func ingressWebhookKlevValidate(w http.ResponseWriter, r *http.Request, now func() time.Time, secret string) ([]byte, error) {
	r.Body = http.MaxBytesReader(w, r.Body, 128*1024)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	hs := r.Header.Get("X-Klev-Signature")
	if len(hs) == 0 {
		return nil, ErrKlevNotSigned()
	}

	var ts string
	var sigs [][]byte
	for _, part := range strings.Split(hs, ";") {
		k, v, found := strings.Cut(part, "=")
		if !found {
			continue
		}

		switch k {
		case "t":
			timestamp, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, ErrKlevTimestampInvalid(v)
			}
			if now().Sub(time.Unix(timestamp, 0)) > 5*time.Minute {
				return nil, ErrKlevTimestampExpired(v)
			}
			ts = v
		case "v1":
			sig, err := hex.DecodeString(v)
			if err != nil {
				continue
			}
			sigs = append(sigs, sig)
		default:
			continue
		}
	}

	if len(ts) == 0 {
		return nil, ErrKlevSignatureTimeMissing(hs)
	}

	if len(sigs) == 0 {
		return nil, ErrKlevSignatureMissing(hs)
	}

	expectedSig := IngressWebhookKlevSignature(ts, payload, secret)

	for _, sig := range sigs {
		if hmac.Equal(expectedSig, sig) {
			return payload, nil
		}
	}

	return nil, ErrKlevSignatureMismatch()
}

func IngressWebhookKlevSignature(ts string, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(ts))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}
