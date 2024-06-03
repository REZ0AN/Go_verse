// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadReuseportlb returns the embedded CollectionSpec for reuseportlb.
func loadReuseportlb() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ReuseportlbBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load reuseportlb: %w", err)
	}

	return spec, err
}

// loadReuseportlbObjects loads reuseportlb and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*reuseportlbObjects
//	*reuseportlbPrograms
//	*reuseportlbMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadReuseportlbObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadReuseportlb()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// reuseportlbSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type reuseportlbSpecs struct {
	reuseportlbProgramSpecs
	reuseportlbMapSpecs
}

// reuseportlbSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type reuseportlbProgramSpecs struct {
	HotStandbySelector *ebpf.ProgramSpec `ebpf:"hot_standby_selector"`
}

// reuseportlbMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type reuseportlbMapSpecs struct {
	TcpBalancingTargets *ebpf.MapSpec `ebpf:"tcp_balancing_targets"`
}

// reuseportlbObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadReuseportlbObjects or ebpf.CollectionSpec.LoadAndAssign.
type reuseportlbObjects struct {
	reuseportlbPrograms
	reuseportlbMaps
}

func (o *reuseportlbObjects) Close() error {
	return _ReuseportlbClose(
		&o.reuseportlbPrograms,
		&o.reuseportlbMaps,
	)
}

// reuseportlbMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadReuseportlbObjects or ebpf.CollectionSpec.LoadAndAssign.
type reuseportlbMaps struct {
	TcpBalancingTargets *ebpf.Map `ebpf:"tcp_balancing_targets"`
}

func (m *reuseportlbMaps) Close() error {
	return _ReuseportlbClose(
		m.TcpBalancingTargets,
	)
}

// reuseportlbPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadReuseportlbObjects or ebpf.CollectionSpec.LoadAndAssign.
type reuseportlbPrograms struct {
	HotStandbySelector *ebpf.Program `ebpf:"hot_standby_selector"`
}

func (p *reuseportlbPrograms) Close() error {
	return _ReuseportlbClose(
		p.HotStandbySelector,
	)
}

func _ReuseportlbClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed reuseportlb_bpfel.o
var _ReuseportlbBytes []byte
