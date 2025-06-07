package command_test

import (
	"context"
	"testing"

	"github.com/localzet/aura/app/dispatcher"
	"github.com/localzet/aura/app/log"
	. "github.com/localzet/aura/app/log/command"
	"github.com/localzet/aura/app/proxyman"
	_ "github.com/localzet/aura/app/proxyman/inbound"
	_ "github.com/localzet/aura/app/proxyman/outbound"
	"github.com/localzet/aura/common"
	"github.com/localzet/aura/common/serial"
	"github.com/localzet/aura/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
