package api

import (
	"github.com/localzet/aura/main/commands/base"
)

// CmdAPI calls an API in an Aura process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "Call an API in an Aura process",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate Aura via its API.
`,
	Commands: []*base.Command{
		cmdRestartLogger,
		cmdGetStats,
		cmdQueryStats,
		cmdSysStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds,
		cmdListInbounds,
		cmdListOutbounds,
		cmdInboundUser,
		cmdInboundUserCount,
		cmdAddRules,
		cmdRemoveRules,
		cmdSourceIpBlock,
		cmdOnlineStats,
		cmdOnlineStatsIpList,
	},
}
