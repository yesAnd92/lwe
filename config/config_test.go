package config

import (
	"fmt"
	"testing"
)

func Test_loadingLweConfig(t *testing.T) {
	type args struct {
		configPath string
		configName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "template test",
			args: args{
				configPath: "./",
				configName: "config_template",
			},
		},
		{
			name: "template test",
			args: args{
				configPath: "",
				configName: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := loadingLweConfig(tt.args.configPath, tt.args.configName)
			fmt.Println(config)
		})
	}
}
