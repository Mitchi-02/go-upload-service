package common_test

import (
	"testing"
	"time"
	"upload-service/pkg/common"
)

func TestGenerateJWT(t *testing.T) {
	userID := "test-user-123"
	secret := "test-secret-key"

	_, err := common.GenerateJWT(userID, secret)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestValidateJWT_ValidToken(t *testing.T) {
	userID := "test-user-123"
	secret := "test-secret-key"

	token, err := common.GenerateJWT(userID, secret)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := common.ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %s, got %s", userID, claims.UserID)
	}

	if claims.Exp <= time.Now().Unix() {
		t.Error("Expected token to not be expired")
	}
}

func TestValidateJWT_InvalidFormat(t *testing.T) {
	invalidToken := "invalid.token"
	secret := "test-secret-key"
	_, err := common.ValidateJWT(invalidToken, secret)
	if err == nil {
		t.Fatal("Expected error for invalid token format")
	}
}

func TestValidateJWT_InvalidSignature(t *testing.T) {
	userID := "test-user-123"
	secret := "test-secret-key"

	token, err := common.GenerateJWT(userID, secret)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Tamper with the token
	tamperedToken := token + "tampered"

	_, err = common.ValidateJWT(tamperedToken, secret)
	if err == nil {
		t.Fatal("Expected error for tampered token")
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	//expired token generated with a past expiration time
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiY2Q5Mjk3N2YtN2M5Yy00NjM2LTg1ZTItMzg5YmIzNzUxMDc1IiwiZXhwIjoxNzUyNjk5MDA3fQ==.2fy8qhoO7W9Xx-OAHVJnnQhHDmJhH3Ax9tT1CnhB2x4="
	secret := "test-secret-key"

	_, err := common.ValidateJWT(expiredToken, secret)
	if err == nil {
		t.Fatal("Expected error for expired token")
	}
}
