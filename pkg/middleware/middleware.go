package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Custom middleware for request logging
func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			stop := time.Now()
			req := c.Request()
			res := c.Response()

			status := res.Status
			size := res.Size
			ip := c.RealIP()
			method := req.Method
			path := req.URL.Path

			logrus.WithFields(logrus.Fields{
				"ip":       ip,
				"method":   method,
				"path":     path,
				"status":   status,
				"size":     size,
				"duration": stop.Sub(start).String(),
			}).Info("HTTP Request")

			return err
		}
	}
}

// Custom middleware for CORS
func CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set CORS headers
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Handle preflight requests
			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}

// Custom middleware for error recovery
func CustomRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
				"path":  c.Request().URL.Path,
				"stack": string(stack),
			}).Error("Panic recovered")

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Internal server error",
			})
		},
	})
}

// JWTClaims represents the JWT claims structure
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// ContextKey is used for context keys to avoid collisions
type ContextKey string

const (
	UserContextKey ContextKey = "user"
)

// JWTMiddleware validates JWT tokens and injects user info into context
func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Authorization header is required",
				})
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Authorization header must start with Bearer",
				})
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse and validate the token
			token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				logrus.WithError(err).Error("Failed to parse JWT token")
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token",
				})
			}

			// Extract claims
			claims, ok := token.Claims.(*JWTClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token claims",
				})
			}

			// Validate token timing claims
			now := time.Now()
			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(now) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Token has expired",
				})
			}

			if claims.NotBefore != nil && claims.NotBefore.After(now) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Token not yet valid",
				})
			}

			if claims.IssuedAt != nil && claims.IssuedAt.After(now) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Token issued in the future",
				})
			}

			// Inject user info into context
			ctx := context.WithValue(c.Request().Context(), UserContextKey, claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// OptionalJWTMiddleware validates JWT tokens if present, but doesn't require them
func OptionalJWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				// No token provided, continue without user context
				return next(c)
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				// Invalid format, continue without user context
				return next(c)
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse and validate the token
			token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				// Invalid token, continue without user context
				logrus.WithError(err).Warn("Invalid JWT token in optional middleware")
				return next(c)
			}

			// Extract claims
			claims, ok := token.Claims.(*JWTClaims)
			if !ok || !token.Valid {
				// Invalid claims, continue without user context
				return next(c)
			}

			// Validate token timing claims
			now := time.Now()
			if claims.ExpiresAt != nil && claims.ExpiresAt.Before(now) {
				// Token expired, continue without user context
				logrus.Warn("Expired JWT token in optional middleware")
				return next(c)
			}

			if claims.NotBefore != nil && claims.NotBefore.After(now) {
				// Token not yet valid, continue without user context
				logrus.Warn("JWT token not yet valid in optional middleware")
				return next(c)
			}

			if claims.IssuedAt != nil && claims.IssuedAt.After(now) {
				// Token issued in the future, continue without user context
				logrus.Warn("JWT token issued in future in optional middleware")
				return next(c)
			}

			// Inject user info into context
			ctx := context.WithValue(c.Request().Context(), UserContextKey, claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// GetUserFromContext retrieves user claims from the request context
func GetUserFromContext(c echo.Context) (*JWTClaims, bool) {
	user, ok := c.Request().Context().Value(UserContextKey).(*JWTClaims)
	return user, ok
}

// InitializeEcho initializes Echo with all middleware
func InitializeEcho() *echo.Echo {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(RequestLogger())
	e.Use(CORS())
	e.Use(CustomRecover())

	// Security middleware
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())

	// Body limit middleware
	e.Use(middleware.BodyLimit("10M"))

	// Hide Echo banner
	e.HideBanner = true

	return e
}
