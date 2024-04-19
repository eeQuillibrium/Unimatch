package jwt_test

import (
	"context"
	"testing"
	"time"

	jwtmy "github.com/eeQuillibrium/Unimatch/auth_service/internal/jwt"
)


func TestGenerateToken(t *testing.T) {
	ctx := context.Background()

	tokenStr, err := jwtmy.GenerateToken(ctx, 1, time.Hour)
	t.Log(tokenStr)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("Token: %s", tokenStr)

	if err != nil {
		t.Errorf("parsing error: %v", err)
	}

}
