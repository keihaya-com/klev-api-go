package klev

type TokenID string

type Token struct {
	TokenID  TokenID  `json:"token_id"`
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
	Bearer   string   `json:"bearer,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens,omitempty"`
}

type TokenCreateParams struct {
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
}
