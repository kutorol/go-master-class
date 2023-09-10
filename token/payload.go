package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	if time.Now().After(p.ExpiredAt) {
		return nil, ErrExpiredToken
	}

	return &jwt.NumericDate{Time: p.ExpiredAt}, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: p.IssuedAt}, nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (p *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p *Payload) GetSubject() (string, error) {
	return "", nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token not valid")

func NewPayload(un string, d time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	p := &Payload{
		ID:        tokenID,
		Username:  un,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(d),
	}

	return p, nil
}
