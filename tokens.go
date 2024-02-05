package klev

type TokenID string

func ParseTokenID(id string) (TokenID, error) {
	if err := validate(id, "tok"); err != nil {
		return TokenID(""), err
	}
	return TokenID(id), nil
}

type Token struct {
	TokenID  TokenID  `json:"token_id"`
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
	Bearer   string   `json:"bearer,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
}

type TokenCreateParams struct {
	Metadata string   `json:"metadata,omitempty"`
	ACL      []string `json:"acl,omitempty"`
}

type TokenUpdateParams struct {
	Metadata *string   `json:"metadata,omitempty"`
	ACL      *[]string `json:"acl,omitempty"`
}
