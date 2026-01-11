package tokens

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/rnymphaea/chronoflow/auth/internal/config"
)

var ErrEmptyUserID = errors.New("empty userID")

type Manager struct {
	accessSecret []byte
	accessTTL    time.Duration
	issuer       string
}

func NewManager(cfg config.JWTConfig) *Manager {
	return &Manager{
		accessSecret: []byte(cfg.Secret),
		accessTTL:    cfg.TTL,
		issuer:       cfg.Issuer,
	}
}

type AccessTokenClaims struct {
	jwt.RegisteredClaims

	UserID string `json:"user_id"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (m *Manager) GenerateTokenPair(userID string) (*TokenPair, error) {
	if userID == "" {
		return nil, ErrEmptyUserID
	}

	accessToken, err := m.generateAccessToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := m.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *Manager) generateAccessToken(userID string) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(m.accessTTL)

	claims := AccessTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.accessSecret)
}

func (m *Manager) generateRefreshToken() (string, error) {
	refreshToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return refreshToken.String(), nil
}
