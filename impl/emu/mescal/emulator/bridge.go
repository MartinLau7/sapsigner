package emulator

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

var (
	nextBridgeAddr = bridgeAddr
)

type Bridge struct {
	uc   unicorn.Unicorn
	hook unicorn.Hook
	addr uint64
	size uint64
	uPtr unsafe.Pointer // retain reference to prevent garbage collection
}

func NewBridge(uc unicorn.Unicorn, data []byte) (*Bridge, error) {
	if len(data) == 0 {
		return &Bridge{}, nil
	}

	prot := unicorn.PROT_READ | unicorn.PROT_WRITE

	uPtr := unsafe.Pointer(&data[0])

	size := uint64(len(data))
	size += (pageSize - size) % pageSize

	addr := nextBridgeAddr
	nextBridgeAddr += size

	if err := uc.MemMapPtr(addr, size, prot, uPtr); err != nil {
		return nil, err
	}

	hookMT := unicorn.HOOK_MEM_READ | unicorn.HOOK_MEM_WRITE

	hookCB := func(uc unicorn.Unicorn, access int, addr uint64, size int, value int64) {
		err := fmt.Errorf("emulator bridge segmentation fault: 0x%016X", addr)
		panic(err)
	}

	hookBA := addr + uint64(len(data))

	hookEA := nextBridgeAddr - 1

	hook, err := uc.HookAdd(hookMT, hookCB, hookBA, hookEA)
	if err != nil {
		return nil, err
	}

	b := Bridge{
		uc:   uc,
		hook: hook,
		addr: addr,
		size: size,
		uPtr: uPtr,
	}
	return &b, nil
}

func (b *Bridge) Addr() uint64 {
	return b.addr
}

var _ io.Closer = (*Bridge)(nil)

func (b *Bridge) Close() error {
	if b.uPtr == nil {
		return nil
	}

	if err := b.uc.MemUnmap(b.addr, b.size); err != nil {
		return err
	}

	if err := b.uc.HookDel(b.hook); err != nil {
		return err
	}

	*b = Bridge{}
	return nil
}
