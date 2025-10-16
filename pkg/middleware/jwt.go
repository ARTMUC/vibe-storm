package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTService handles JWT token operations
type JWTService struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTService creates a new JWT service instance
func NewJWTService(secretKey string, tokenDuration time.Duration) *JWTService {
	return &JWTService{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

// GenerateToken generates a new JWT token for the given user information
func (s *JWTService) GenerateToken(userID, username, email string) (string, error) {
	claims := &JWTClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "vibe-storm",
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token string and returns the claims
func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// Validate token timing claims
	if err := s.validateTokenTiming(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// IsTokenExpired checks if a token is expired without full validation
func (s *JWTService) IsTokenExpired(tokenString string) bool {
	claims, err := s.parseTokenWithoutValidation(tokenString)
	if err != nil {
		return true // Consider invalid tokens as expired
	}

	if claims.ExpiresAt == nil {
		return false // No expiration set, consider not expired
	}

	now := time.Now()
	return claims.ExpiresAt.Time.Before(now)
}

// GetTokenExpiration returns the expiration time of a token
func (s *JWTService) GetTokenExpiration(tokenString string) (*time.Time, error) {
	claims, err := s.parseTokenWithoutValidation(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt == nil {
		return nil, jwt.ErrInvalidKey
	}

	expiryTime := claims.ExpiresAt.Time
	return &expiryTime, nil
}

// GetTimeUntilExpiration returns the duration until token expiration
func (s *JWTService) GetTimeUntilExpiration(tokenString string) (time.Duration, error) {
	expiryTime, err := s.GetTokenExpiration(tokenString)
	if err != nil {
		return 0, err
	}

	return time.Until(*expiryTime), nil
}

// parseTokenWithoutValidation parses a token without validating claims (for inspection only)
func (s *JWTService) parseTokenWithoutValidation(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// validateTokenTiming validates the timing claims of a token
func (s *JWTService) validateTokenTiming(claims *JWTClaims) error {
	now := time.Now()

	// Check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(now) {
		return jwt.ErrTokenExpired
	}

	// Check not before
	if claims.NotBefore != nil && claims.NotBefore.Time.After(now) {
		return jwt.ErrTokenNotValidYet
	}

	// Check issued at (should not be in future)
	if claims.IssuedAt != nil && claims.IssuedAt.Time.After(now) {
		return jwt.ErrTokenInvalidClaims
	}

	return nil
}

// RefreshToken generates a new token with updated expiration time
func (s *JWTService) RefreshToken(claims *JWTClaims) (string, error) {
	// Create new claims with updated expiration
	newClaims := &JWTClaims{
		UserID:   claims.UserID,
		Username: claims.Username,
		Email:    claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "vibe-storm",
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
