// Package register registers all relevant bases and also subtype specific functions
package register

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/utils/rpc"

	"go.viam.com/rdk/component/base"

	// register bases.
	_ "go.viam.com/rdk/component/base/fake"
	_ "go.viam.com/rdk/component/base/wheeled"
	componentpb "go.viam.com/rdk/proto/api/component/v1"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/subtype"
)

func init() {
	registry.RegisterResourceSubtype(base.Subtype, registry.ResourceSubtype{
		Reconfigurable: base.WrapWithReconfigurable,
		RegisterRPCService: func(ctx context.Context, rpcServer rpc.Server, subtypeSvc subtype.Service) error {
			return rpcServer.RegisterServiceServer(
				ctx,
				&componentpb.BaseService_ServiceDesc,
				base.NewServer(subtypeSvc),
				componentpb.RegisterBaseServiceHandlerFromEndpoint,
			)
		},
		RPCClient: func(ctx context.Context, conn rpc.ClientConn, name string, logger golog.Logger) interface{} {
			return base.NewClientFromConn(ctx, conn, name, logger)
		},
	})
}