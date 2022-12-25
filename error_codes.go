// Code generated by 'make gen-errors'; DO NOT EDIT
package api

// Internal errors
var (
	ErrInternalServerError      = "ERR_KLEV_INTERNAL_0001"
	ErrInternalMaintenanceError = "ERR_KLEV_INTERNAL_0002"
)

// API errors
var (
	ErrAPIInvalidContentType      = "ERR_KLEV_API_0001"
	ErrAPINoAuthorizationHeader   = "ERR_KLEV_API_0002"
	ErrAPIInvalidToken            = "ERR_KLEV_API_0003"
	ErrAPIInvalidPayment          = "ERR_KLEV_API_0004"
	ErrAPIPathNotFound            = "ERR_KLEV_API_0005"
	ErrAPIMethodNotAllowed        = "ERR_KLEV_API_0006"
	ErrAPITokenInvalidMissing     = "ERR_KLEV_API_0007"
	ErrAPILogInvalidMissing       = "ERR_KLEV_API_0008"
	ErrAPIInvalidJson             = "ERR_KLEV_API_0009"
	ErrAPIInvalidOffset           = "ERR_KLEV_API_0010"
	ErrAPIInvalidLen              = "ERR_KLEV_API_0011"
	ErrAPIInvalidEncoding         = "ERR_KLEV_API_0012"
	ErrAPIWebhookInvalidMissing   = "ERR_KLEV_API_0013"
	ErrAPIInsufficientPermissions = "ERR_KLEV_API_0014"
	ErrAPIInvalidACLItemFormat    = "ERR_KLEV_API_0015"
	ErrAPIInvalidACLSubject       = "ERR_KLEV_API_0016"
	ErrAPIInvalidACLAction        = "ERR_KLEV_API_0017"
	ErrAPIInvalidACLObjectAction  = "ERR_KLEV_API_0018"
	ErrAPIInvalidACLObjectToken   = "ERR_KLEV_API_0019"
	ErrAPIInvalidACLObjectLog     = "ERR_KLEV_API_0020"
	ErrAPIInvalidACLObjectWebhook = "ERR_KLEV_API_0021"
	ErrAPIOffsetInvalidMissing    = "ERR_KLEV_API_0022"
	ErrAPIOffsetLogMismatch       = "ERR_KLEV_API_0023"
	ErrAPIInvalidACLObjectOffset  = "ERR_KLEV_API_0024"
	ErrAPIInvalidPoll             = "ERR_KLEV_API_0025"
)

// APIWebhooks errors
var (
	ErrAPIWebhooksPathNotFound          = "ERR_KLEV_API_WEBHOOKS_0001"
	ErrAPIWebhooksWebhookInvalidMissing = "ERR_KLEV_API_WEBHOOKS_0002"
)

// Logs errors
var (
	ErrLogsNotFound         = "ERR_KLEV_LOGS_0001"
	ErrLogsMaxMetadata      = "ERR_KLEV_LOGS_0002"
	ErrLogsAgeCompactExpire = "ERR_KLEV_LOGS_0003"
	ErrLogsMaxCount         = "ERR_KLEV_LOGS_0004"
)

// Messages errors
var (
	ErrMessagesMaxPublish            = "ERR_KLEV_MESSAGES_0001"
	ErrMessagesMaxKey                = "ERR_KLEV_MESSAGES_0002"
	ErrMessagesMaxValue              = "ERR_KLEV_MESSAGES_0003"
	ErrMessagesMaxConsume            = "ERR_KLEV_MESSAGES_0004"
	ErrMessagesConsumeInvalid        = "ERR_KLEV_MESSAGES_0005"
	ErrMessagesGetByOffsetNotFound   = "ERR_KLEV_MESSAGES_0006"
	ErrMessagesGetByOffsetInvalid    = "ERR_KLEV_MESSAGES_0007"
	ErrMessagesGetByKeyNotCompacting = "ERR_KLEV_MESSAGES_0008"
	ErrMessagesGetByKeyNotFound      = "ERR_KLEV_MESSAGES_0009"
	ErrMessagesMaxPoll               = "ERR_KLEV_MESSAGES_0010"
)

// Offsets errors
var (
	ErrOffsetsNotFound         = "ERR_KLEV_OFFSETS_0001"
	ErrOffsetsMaxMetadata      = "ERR_KLEV_OFFSETS_0002"
	ErrOffsetsMaxValueMetadata = "ERR_KLEV_OFFSETS_0003"
	ErrOffsetsMaxCount         = "ERR_KLEV_OFFSETS_0004"
)

// Tokens errors
var (
	ErrTokensNotFound     = "ERR_KLEV_TOKENS_0001"
	ErrTokensInvalidToken = "ERR_KLEV_TOKENS_0002"
	ErrTokensMaxMetadata  = "ERR_KLEV_TOKENS_0003"
	ErrTokensMaxCount     = "ERR_KLEV_TOKENS_0004"
)

// Webhooks errors
var (
	ErrWebhooksNotFound    = "ERR_KLEV_WEBHOOKS_0001"
	ErrWebhooksMaxMetadata = "ERR_KLEV_WEBHOOKS_0002"
	ErrWebhooksUnknownType = "ERR_KLEV_WEBHOOKS_0003"
	ErrWebhooksMaxCount    = "ERR_KLEV_WEBHOOKS_0004"
)
