package gitcmd

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

const testGitDir = "."

func TestGetCommitLog(t *testing.T) {
	type args struct {
		detail  bool
		recentN int16
		dir     string
		author  string
		start   string
		end     string
		branchs bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				detail:  false,
				recentN: 3,
				dir:     testGitDir,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommitLog(tt.args.detail, tt.args.recentN, tt.args.dir, tt.args.author, tt.args.start, tt.args.end, tt.args.branchs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, log := range *got {
				fmt.Printf("Branch:%s-----------Hash:%s-----------\n", log.Branch, log.CommitHash)
				fmt.Printf("@%s  %s commit msg: %s\n\n", log.Username, log.CommitAt, log.CommitMsg)
			}
		})
	}
}

func TestGetChangedFile(t *testing.T) {
	type args struct {
		commitId string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				//current lwe git repository commit id 6f635d7
				commitId: "6f635d7",
			},
			want:    []string{".gitignore", "README.md"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChangedFile(tt.args.commitId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChangedFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChangedFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllGitRepoCommitLog2(t *testing.T) {

	resLogs, _ := GetAllGitRepoCommitLog(true, 3, testGitDir, "", "", "", false)

	//控制台
	console := ConsoleOutput{}
	console.Output(resLogs)

	defer func() {
		os.Remove(REPORT_PATH)
	}()
	//写文件
	file := FileOutput{}
	file.Output(resLogs)

	if f, err := os.Stat(REPORT_PATH); err != nil || f.Size() == 0 {
		t.Error("file not exist >>>", REPORT_PATH)
	}
}

func TestGetAllGitRepoCommitLog(t *testing.T) {
	type args struct {
		detail  bool
		recentN int16
		dir     string
		author  string
		start   string
		end     string
		branchs bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				detail:  false,
				recentN: 3,
				dir:     testGitDir,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllGitRepoCommitLog(tt.args.detail, tt.args.recentN, tt.args.dir, tt.args.author, tt.args.start, tt.args.end, tt.args.branchs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllGitRepoCommitLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer func() {
				os.Remove(REPORT_PATH)
			}()
			//控制台
			console := ConsoleOutput{}
			console.Output(got)

			//写文件
			file := FileOutput{}
			file.Output(got)
			if f, err := os.Stat(REPORT_PATH); err != nil || f.Size() == 0 {
				t.Error("file not exist >>>", REPORT_PATH)
			}
		})
	}
}
