package jwt

// generate by claude (TODO: fix and run)
//func TestJWT_CreateAndVerifyToken(t *testing.T) {
//	// Setup
//	cfg := &config.Config{
//		Token: config.TokenConfig{
//			Key:      "your-test-secret-key",
//			Issuer:   "test-issuer",
//			Audience: "test-audience",
//			Duration: config.TokenDuration{
//				Access:  "15m",
//				Refresh: "24h",
//			},
//		},
//	}
//
//	jwtService, err := NewJWT(cfg)
//	assert.NoError(t, err)
//
//	testUser := &domain.User{
//		ID:     1,
//		Social: domain.SocialKakao,
//		Role:   domain.RoleUser,
//	}
//
//	tests := []struct {
//		name    string
//		setup   func() (string, error)
//		wantErr bool
//	}{
//		{
//			name: "should create and verify valid access token",
//			setup: func() (string, error) {
//				return jwtService.CreateAccessToken(testUser)
//			},
//			wantErr: false,
//		},
//		{
//			name: "should create and verify valid refresh token",
//			setup: func() (string, error) {
//				return jwtService.CreateRefreshToken(testUser)
//			},
//			wantErr: false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Create token
//			token, err := tt.setup()
//			assert.NoError(t, err)
//			assert.NotEmpty(t, token)
//
//			// Verify token
//			payload, err := jwtService.VerifyToken(token)
//			if tt.wantErr {
//				assert.Error(t, err)
//				return
//			}
//
//			assert.NoError(t, err)
//			assert.NotNil(t, payload)
//			assert.Equal(t, cfg.Token.Issuer, payload.Issuer)
//			assert.Equal(t, cfg.Token.Audience, payload.Audience)
//			assert.Equal(t, "1", payload.Subject)
//			assert.Equal(t, domain.SocialKakao, payload.Social)
//			assert.Equal(t, domain.RoleUser, payload.Role)
//		})
//	}
//}
