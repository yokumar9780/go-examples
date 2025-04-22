package middleware

import (
	"encoding/json"
	"github.com/MicahParks/keyfunc/v3"
	"io"
)
import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwks keyfunc.Keyfunc

func SetupJWKS() {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		log.Fatal("JWT_ISSUER environment variable is not set")
	}
	jwksURL := issuer + "/.well-known/jwks.json"

	resp, err := http.Get(jwksURL)
	if err != nil {
		log.Fatalf("Failed to fetch JWKS from %s: %v", jwksURL, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read JWKS response body: %v", err)
	}

	jwks, err = keyfunc.NewJWKSetJSON(json.RawMessage(body))

	if err != nil {
		log.Fatalf("Failed to parse JWKS JSON: %v", err)
	}

}

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// ✅ Skip JWT validation for Swagger UI and docs
		path := c.Request().URL.Path
		if strings.HasPrefix(path, "/swagger/") || path == "/swagger/index.html" || strings.HasSuffix(path, "swagger.json") {
			return next(c)
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Load config
		//jwtIssuer := os.Getenv("JWT_ISSUER")
		//jwtAudience := os.Getenv("JWT_AUDIENCE")
		//jwtSecret := []byte(os.Getenv("JWT_SECRET")) // Use your secret OR public key if RS256

		// Parse token manually without validating `aud`
		parser := jwt.NewParser(
			jwt.WithoutClaimsValidation(), // disables automatic claims validation
		)
		claims := jwt.RegisteredClaims{}
		token, err := parser.ParseWithClaims(tokenString, &claims, jwks.Keyfunc)
		if err != nil {
			log.Error("Token parse error:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token: "+err.Error())
		}
		log.Info(token.Claims)

		/*token, err := jwt.Parse(tokenString, jwks.Keyfunc,
			jwt.WithAudience(jwtAudience),
			jwt.WithIssuer(jwtIssuer), dis)
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token: "+err.Error())
		}

		// Optional: set user claims in context
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user", claims)
		*/
		return next(c)
	}
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		stop := time.Now()
		log.WithFields(log.Fields{
			"level":     "info",
			"time":      time.Now().Format(time.RFC3339),
			"method":    c.Request().Method,
			"path":      c.Request().URL.Path,
			"status":    c.Response().Status,
			"duration":  stop.Sub(start).String(),
			"requestId": c.Response().Header().Get("X-Request-ID"),
		}).Info("Handled request")
		return err
	}
}
func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.New().String()
		c.Request().Header.Set("X-Request-ID", reqID)
		c.Response().Header().Set("X-Request-ID", reqID)
		return next(c)
	}
}
