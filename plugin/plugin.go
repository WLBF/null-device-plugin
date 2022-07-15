package plugin

import (
	"context"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type DevicePlugin interface {
	Start() error
	Stop() error
	Serve() error
	Register() error
	GetDevicePluginOptions(context.Context, *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error)
	ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error
	GetPreferredAllocation(ctx context.Context, r *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error)
	Allocate(ctx context.Context, reqs *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error)
	PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error)
}
