// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package bandwidth

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/hive/cell"
	"github.com/cilium/cilium/pkg/option"
)

var Cell = cell.Module(
	"bandwidth-manager",
	"Linux Bandwidth Manager for EDT-based pacing",

	cell.Config(Config{false, false}),
	cell.Provide(newBandwidthManager),
)

type ManagerInterface interface {
	BBREnabled() bool
	DeleteEndpointBandwidthLimit(epID uint16) error
	Enabled() bool
	ResetQueues() bool
}

type Config struct {
	// EnableBandwidthManager enables EDT-based pacing
	EnableBandwidthManager bool

	// EnableBBR enables BBR TCP congestion control for the node including Pods
	EnableBBR bool
}

func (def Config) Flags(flags *pflag.FlagSet) {
	flags.Bool("enable-bandwidth-manager", def.EnableBandwidthManager, "Enable BPF bandwidth manager")
	flags.Bool(EnableBBR, def.EnableBBR, "Enable BBR for the bandwidth manager")
}

func newBandwidthManager(lc hive.Lifecycle, p bandwidthManagerParams) ManagerInterface {
	m := &Manager{params: p}
	lc.Append(m)
	return m
}

func (m *Manager) Start(hive.HookContext) error {
	err := m.probe()
	if err != nil {
		return err
	} else if !m.enabled {
		return nil
	}

	return m.init()
}

func (*Manager) Stop(hive.HookContext) error {
	return nil
}

type bandwidthManagerParams struct {
	cell.In

	Log          logrus.FieldLogger
	Config       Config
	DaemonConfig *option.DaemonConfig
}
