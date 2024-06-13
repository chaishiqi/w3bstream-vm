package vm

import (
	"context"

	"github.com/machinefi/sprout-vm/wasm-server-go/vm/wasmtime"
)

func NewInstance(ctx context.Context, code []byte, id string) (*wasmtime.Instance, error) {
	return wasmtime.NewInstanceByCode(ctx, id, code)
}
