package all

import (
	"github.com/localzet/aura/main/commands/all/api"
	"github.com/localzet/aura/main/commands/all/convert"
	"github.com/localzet/aura/main/commands/all/tls"
	"github.com/localzet/aura/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
