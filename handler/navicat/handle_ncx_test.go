package navicat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseNcx(t *testing.T) {

	tests := []struct {
		name   string
		data   []byte
		expect string
	}{
		// test cases
		{
			name:   "pwd case",
			data:   []byte(`<Connections Ver="1.5"><Connection ConnectionName="pwd case" ProjectUUID="" ConnType="MYSQL" OraConnType="" ServiceProvider="Default" Host="127.0.0.1" Port="3306" Database="" OraServiceNameType="" TNS="" MSSQLAuthenMode="" MSSQLAuthenWindowsDomain="" DatabaseFileName="" UserName="root" Password="B75D320B6211468D63EB3B67C9E85933" SavePassword="true" SettingsSavePath="/Users/wangyj/Library/Application Support/PremiumSoft CyberTech/Navicat CC/Common/Settings/0/0/MySQL/冀云测试 2" SessionLimit="0" ClientCharacterSet="" ClientEncoding="65001" Keepalive="false" KeepaliveInterval="240" Encoding="65001" MySQLCharacterSet="true" Compression="false" AutoConnect="false" NamedPipe="false" NamedPipeSocket="" OraRole="" OraOSAuthen="false" SQLiteEncrypt="false" SQLiteEncryptPassword="" SQLiteSaveEncryptPassword="false" UseAdvanced="false" SSL="false" SSL_Authen="false" SSL_PGSSLMode="REQUIRE" SSL_ClientKey="" SSL_ClientCert="" SSL_CACert="" SSL_Clpher="" SSL_PGSSLCRL="" SSL_WeakCertValidation="false" SSL_AllowInvalidHostName="false" SSL_PEMClientKeyPassword="" SSH="false" SSH_Host="" SSH_Port="22" SSH_UserName="" SSH_AuthenMethod="PASSWORD" SSH_Password="" SSH_SavePassword="false" SSH_PrivateKey="" SSH_Passphrase="" SSH_SavePassphrase="false" SSH_Compress="false" HTTP="false" HTTP_URL="" HTTP_PA="" HTTP_PA_UserName="" HTTP_PA_Password="" HTTP_PA_SavePassword="" HTTP_EQ="" HTTP_CA="" HTTP_CA_ClientKey="" HTTP_CA_ClientCert="" HTTP_CA_CACert="" HTTP_CA_Passphrase="" HTTP_Proxy="" HTTP_Proxy_Host="" HTTP_Proxy_Port="" HTTP_Proxy_UserName="" HTTP_Proxy_Password="" HTTP_Proxy_SavePassword="" Compatibility="false" Compatibility_OverrideLowerCaseTableNames="false" Compatibility_LowerCaseTableNames="" Compatibility_OverrideSQLMode="false" Compatibility_SQLMode="" Compatibility_OverrideIsSupportNDB="false" Compatibility_IsSupportNDB="false" Compatibility_OverrideDatabaseListingMethod="false" Compatibility_DatabaseListingMethod="" Compatibility_OverrideViewListingMethod="false" Compatibility_ViewListingMethod=""/></Connections>`),
			expect: `{"Conns":[{"ConnectionName":"pwd case","ConnType":"MYSQL","Host":"127.0.0.1","UserName":"root","Port":"3306","Password":"This is a test"}],"Version":"1.5"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCons, err := ParseNcx(tt.data)
			result, err := json.Marshal(gotCons)
			if err != nil {
				t.Errorf("ParseNcx() error %e", err)
				return
			}
			if !reflect.DeepEqual(string(result), tt.expect) {
				t.Errorf("except:%s\n	actual:%s", tt.expect, result)
			}
		})
	}
}

func Test_decryptPwd(t *testing.T) {
	type args struct {
		encryptTxt string
		expect     string
	}
	tests := []struct {
		name string
		args args
	}{
		// test cases.
		{
			name: "case 1",
			args: args{
				encryptTxt: "833E4ABBC56C89041A9070F043641E3B",
				expect:     "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decryptPwd(tt.args.encryptTxt)
			if err != nil {
				t.Errorf("decryptPwd() error %e", err)
				return
			}
			if !reflect.DeepEqual(got, tt.args.expect) {
				t.Errorf("except:%s\n actual:%s", tt.args.expect, got)
			}
		})
	}
}
