// Copyright 2026 PingCAP, Inc. Licensed under Apache-2.0.

// Package noauth provides a simple authenticator for clusters without TiDB instances.
// It accepts a username and password without SQL verification.
package noauth

import (
	"go.uber.org/fx"

	"github.com/pingcap/tidb-dashboard/pkg/apiserver/user"
	"github.com/pingcap/tidb-dashboard/pkg/apiserver/utils"
	"github.com/pingcap/tidb-dashboard/pkg/config"
)

const typeID utils.AuthType = 0

type Authenticator struct {
	user.BaseAuthenticator
	authService *user.AuthService
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func registerAuthenticator(a *Authenticator, authService *user.AuthService) {
	authService.RegisterAuthenticator(typeID, a)
	a.authService = authService
}

var Module = fx.Options(
	fx.Provide(NewAuthenticator),
	fx.Invoke(registerAuthenticator),
)

func (a *Authenticator) Authenticate(f user.AuthenticateForm) (*utils.SessionUser, error) {
	// In no-TiDB mode, we accept the credentials directly without SQL verification.
	// The default root user with empty password is accepted.
	plainPwd := f.Password
	if a.authService != nil && a.authService.RsaPrivateKey != nil {
		decrypted, err := user.Decrypt(f.Password, a.authService.RsaPrivateKey)
		if err == nil {
			plainPwd = decrypted
		}
		// If decryption fails, use the password as-is (might be plain text)
	}

	// Accept root user with any password in no-TiDB mode
	username := f.Username
	if username == "" {
		username = "root"
	}

	_ = plainPwd // password is not verified in no-TiDB mode

	return &utils.SessionUser{
		Version:      utils.SessionVersion,
		HasTiDBAuth:  false, // No TiDB auth - this is important to skip TiDB connection middleware
		TiDBUsername: username,
		TiDBPassword: "",
		DisplayName:  username + " (no-tidb)",
		IsShareable:  true,
		IsWriteable:  true,
	}, nil
}

// NeedConfig checks whether noauth should be used based on config.
func NeedConfig(cfg *config.Config) bool {
	return cfg.NoTiDB
}
