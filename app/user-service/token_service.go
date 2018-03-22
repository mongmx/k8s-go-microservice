package main

// Authable interface
type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

// TokenService struct
type TokenService struct {
	repo Repository
}

// Decode function
func (srv *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

// Encode function
func (srv *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
