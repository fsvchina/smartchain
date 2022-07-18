package module

import (
	"fmt"

	"github.com/gogo/protobuf/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)





type Configurator interface {



	MsgServer() grpc.Server



	QueryServer() grpc.Server




	//





	RegisterMigration(moduleName string, forVersion uint64, handler MigrationHandler) error
}

type configurator struct {
	cdc         codec.Codec
	msgServer   grpc.Server
	queryServer grpc.Server


	migrations map[string]map[uint64]MigrationHandler
}


func NewConfigurator(cdc codec.Codec, msgServer grpc.Server, queryServer grpc.Server) Configurator {
	return configurator{
		cdc:         cdc,
		msgServer:   msgServer,
		queryServer: queryServer,
		migrations:  map[string]map[uint64]MigrationHandler{},
	}
}

var _ Configurator = configurator{}


func (c configurator) MsgServer() grpc.Server {
	return c.msgServer
}


func (c configurator) QueryServer() grpc.Server {
	return c.queryServer
}


func (c configurator) RegisterMigration(moduleName string, forVersion uint64, handler MigrationHandler) error {
	if forVersion == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, "module migration versions should start at 1")
	}

	if c.migrations[moduleName] == nil {
		c.migrations[moduleName] = map[uint64]MigrationHandler{}
	}

	if c.migrations[moduleName][forVersion] != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic, "another migration for module %s and version %d already exists", moduleName, forVersion)
	}

	c.migrations[moduleName][forVersion] = handler

	return nil
}



func (c configurator) runModuleMigrations(ctx sdk.Context, moduleName string, fromVersion, toVersion uint64) error {

	if toVersion <= 1 {
		return nil
	}

	moduleMigrationsMap, found := c.migrations[moduleName]
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "no migrations found for module %s", moduleName)
	}


	for i := fromVersion; i < toVersion; i++ {
		migrateFn, found := moduleMigrationsMap[i]
		if !found {
			return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "no migration found for module %s from version %d to version %d", moduleName, i, i+1)
		}
		ctx.Logger().Info(fmt.Sprintf("migrating module %s from version %d to version %d", moduleName, i, i+1))

		err := migrateFn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
