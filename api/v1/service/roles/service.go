package serviceroles

import (
	"context"
	reporoles "github.com/RyaWcksn/nann-e/api/v1/repository/roles"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type RolesService struct {
	rolesRepo reporoles.IRepository
	l         logger.ILogger
}

func NewRolesService(rolesRepo reporoles.IRepository, l logger.ILogger) *RolesService {
	return &RolesService{
		rolesRepo: rolesRepo,
		l:         l,
	}
}

type IService interface {
	CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) (*entityroles.CreateRoleDetails, error)
}

var _ IService = (*RolesService)(nil)
