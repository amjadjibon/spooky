package echo

import (
	"context"
	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
)

type Echo struct {

}

func (e *Echo) Name() string {
	return "abesh_example_echo"
}

func (e *Echo) Version() string {
	return "0.0.1"
}

func (e *Echo) Category() string {
	return string(constant.CategoryService)
}

func (e *Echo) ContractId() string {
	return "abesh:ex_echo"
}

func (e *Echo) GetConfigMap() iface.ConfigMap {
	panic("implement me")
}

func (e *Echo) Setup() error {
	panic("implement me")
}

func (e *Echo) SetConfigMap(values iface.ConfigMap) error {
	panic("implement me")
}

func (e *Echo) New() iface.ICapability {
	return &Echo{}
}

func (e *Echo) Serve(_ context.Context, _ iface.ICapabilityRegistry, input *model.Event) (*model.Event, error) {
	panic("implement me")
}

func init() {
	registry.GlobalRegistry().AddCapability(&Echo{})
}