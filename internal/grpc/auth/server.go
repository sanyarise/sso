package auth

import (
	ssov1 "github.com/sanyarise/protos"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
