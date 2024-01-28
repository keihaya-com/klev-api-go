// Code generated by 'make gen-errors'; DO NOT EDIT
package logs

import "github.com/klev-dev/klev-api-go/errors"

const (
	ErrLogPathInvalid = "ERR_KLEV_LOGS_API_0001"
)

func IsErrLogPathInvalid(err error) bool {
	return errors.IsError(err, ErrLogPathInvalid)
}
