package handlers

import (
	"github.com/fahmifan/shortly/gen/models"
	errors "github.com/go-openapi/errors"
	jwt "github.com/golang-jwt/jwt"
)

type roleClaims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

type AuthHandler struct {
	Context *Context
}

func NewAuthHandler(a *AuthHandler) *AuthHandler {
	return a
}

func (a *AuthHandler) HasRole(token string, scopes []string) (*models.Principal, error) {
	claims, err := a.parseAndCheckToken(token)
	if err != nil {
		return nil, errors.New(401, "Unauthorized: invalid Bearer token: %v", err)
	}

	scopeMap := make(map[string]struct{}, len(scopes))
	for _, scope := range scopes {
		scopeMap[scope] = struct{}{}
	}

	var claimedRoles []string
	hasScope := false
	for _, role := range claims.Roles {
		if _, ok := scopeMap[role]; ok {
			hasScope = true
			claimedRoles = append(claimedRoles, role)
		}
	}

	if !hasScope {
		return nil, errors.New(403, "Forbidden: insufficient privileges")
	}

	return &models.Principal{
		Name:  claims.Id,
		Roles: claimedRoles,
	}, nil
}

func (a *AuthHandler) parseAndCheckToken(token string) (*roleClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &roleClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		return a.Context.JWTSecretKey, nil
	})

	if err == nil {
		if claims, ok := parsedToken.Claims.(*roleClaims); ok && parsedToken.Valid {
			return claims, nil
		}
	}
	return nil, err
}
