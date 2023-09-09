package sync

import (
	"testing"
)

func TestFsync_compareDirDiff(t *testing.T) {

	tests := []struct {
		name  string
		fsync *Fsync
	}{
		// TODO: Add test cases.
		{
			name:  "mac",
			fsync: InitFsync("/Users/wangyj/ideaProject/my/lwe", "/Users/wangyj/Desktop/lwe_copy"),
		},
		{
			name:  "win",
			fsync: InitFsync("D:\\ideaProject\\my\\go_works\\src\\lwe", "C:\\Users\\Administrator\\Desktop\\lwe_copy"),
		},
		{
			name:  "empty dir",
			fsync: InitFsync("/Users/wangyj/Desktop/a", "/Users/wangyj/Desktop/a"),
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			f := tt.fsync
			f.DiffDir()
			f.Sync(&DisplayCompareThenDo{})
		})
	}
}

func TestFsync_CopyCompareThenDo(t *testing.T) {

	tests := []struct {
		name  string
		fsync *Fsync
	}{
		// TODO: Add test cases.
		{
			name:  "mac",
			fsync: InitFsync("/Users/wangyj/ideaProject/my/lwe", "/Users/wangyj/Desktop/lwe_copy"),
		},
		{
			name:  "win",
			fsync: InitFsync("D:\\ideaProject\\my\\go_works\\src\\lwe", "C:\\Users\\Administrator\\Desktop\\lwe_copy"),
		},
		{
			name:  "empty dir",
			fsync: InitFsync("/Users/wangyj/Desktop/a", "/Users/wangyj/Desktop/a"),
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			f := tt.fsync
			f.DiffDir()
			f.Sync(&CopyCompareThenDo{})
		})
	}
}
