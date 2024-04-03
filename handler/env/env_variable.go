package env

import (
	"fmt"
	"github.com/yesAnd92/lwe/utils"
	"strings"
	"time"
)

type EnvInfo struct {
	Path     string
	envItems []string
}

// MacEvnPath MAC env variable file
var MacEvnPath = []string{
	"~/.bash_profile", "~/.zshrc", "~/.bashrc"}

// LinuxEvnPath linux env variable file
var LinuxEvnPath = []string{
	"~/.zshrc", "~/.bashrc", "/etc/environment"}

type AbstractEnvV struct {
}

func EnvCat(envInfos []*EnvInfo) {
	for _, info := range envInfos {
		fmt.Printf("%s >>>\n", info.Path)
		for _, item := range info.envItems {
			fmt.Println(item)
		}

		fmt.Printf("\n\n")
	}
}

type IEnvVariable interface {
	FindEnvInfo() []*EnvInfo
}

type WinEVnVariable struct {
}

func (w *WinEVnVariable) FindEnvInfo() []*EnvInfo {

	//TODO implement me
	panic("implement me")
}

type MacEVnVariable struct {
}

func (m *MacEVnVariable) FindEnvInfo() (envInfos []*EnvInfo) {

	for _, path := range MacEvnPath {
		cmd := fmt.Sprintf("cat %s", path)
		result := utils.RunCmd(cmd, 2*time.Second)
		if result.Err() != nil {
			continue
		}

		var envItems []string
		envLineSplit := strings.Split(result.String(), "\n")
		for _, envLine := range envLineSplit {
			envLine = strings.TrimSpace(envLine)
			//  filter comment line
			if strings.HasPrefix(envLine, "#") || len(envLine) == 0 {
				continue
			}
			envItems = append(envItems, envLine)

		}

		if len(envItems) == 0 {
			continue
		}
		envInfos = append(envInfos, &EnvInfo{
			Path:     path,
			envItems: envItems,
		})

	}

	return
}
