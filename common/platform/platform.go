package platform // import "github.com/localzet/aura/common/platform"

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	PluginLocation  = "aura.location.plugin"
	ConfigLocation  = "aura.location.config"
	ConfdirLocation = "aura.location.confdir"
	ToolLocation    = "aura.location.tool"
	AssetLocation   = "aura.location.asset"
	CertLocation    = "aura.location.cert"

	UseReadV         = "aura.buf.readv"
	UseFreedomSplice = "aura.buf.splice"
	UseVmessPadding  = "aura.vmess.padding"
	UseCone          = "aura.cone.disabled"

	BufferSize           = "aura.ray.buffer.size"
	BrowserDialerAddress = "aura.browser.dialer"
	XUDPLog              = "aura.xudp.show"
	XUDPBaseKey          = "aura.xudp.basekey"
)

type EnvFlag struct {
	Name    string
	AltName string
}

func NewEnvFlag(name string) EnvFlag {
	return EnvFlag{
		Name:    name,
		AltName: NormalizeEnvName(name),
	}
}

func (f EnvFlag) GetValue(defaultValue func() string) string {
	if v, found := os.LookupEnv(f.Name); found {
		return v
	}
	if len(f.AltName) > 0 {
		if v, found := os.LookupEnv(f.AltName); found {
			return v
		}
	}

	return defaultValue()
}

func (f EnvFlag) GetValueAsInt(defaultValue int) int {
	useDefaultValue := false
	s := f.GetValue(func() string {
		useDefaultValue = true
		return ""
	})
	if useDefaultValue {
		return defaultValue
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int(v)
}

func NormalizeEnvName(name string) string {
	return strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(name)), ".", "_")
}

func getExecutableDir() string {
	exec, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exec)
}

func getExecutableSubDir(dir string) func() string {
	return func() string {
		return filepath.Join(getExecutableDir(), dir)
	}
}

func GetPluginDirectory() string {
	pluginDir := NewEnvFlag(PluginLocation).GetValue(getExecutableSubDir("plugins"))
	return pluginDir
}

func GetConfigurationPath() string {
	configPath := NewEnvFlag(ConfigLocation).GetValue(getExecutableDir)
	return filepath.Join(configPath, "config.json")
}

// GetConfDirPath reads "aura.location.confdir"
func GetConfDirPath() string {
	configPath := NewEnvFlag(ConfdirLocation).GetValue(func() string { return "" })
	return configPath
}
