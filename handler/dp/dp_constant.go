package dp

var moduleNameMap = map[string]string{
	"new-media-rft":         "rft;广播电视",
	"new-media-rft-api":     "rftapi;广播电视api",
	"new-media-mp":          "mp;媒体号",
	"new-media-mp-api":      "mpapi;媒体号",
	"new-media-mpclient":    "mpclient;媒体号客户端",
	"new-media-politic":     "mpapi;问政管理",
	"new-media-politic-api": "politicapi;问政管理api",
	"new-media-politicunit": "politicunit;问政管理客户端",
	"new-media-content":     "content;新闻管理",
	"new-media-content-api": "contentapi;新闻管理api",
}

var pjNameMap = map[string]string{
	"lwe":             "升级测试项目;9527",
	"product-jiyun":   "冀云;3",
	"sxcloud":         "陕西;9",
	"product-gansu":   "甘肃;10",
	"product-ganzhou": "赣州;11",
}

var username = ""
var pwd = ""

var dpResultTpl = `
工单：{{.WorkSheetSn}}
部署升级命令： @{{.ModuleNames}}
tag:{{.Tag}}
升级项目：{{.PjName}}
升级哪些功能：{{.Msg}}
升级时长：5分钟
影响模块：{{.ModuleDesc}}
是否经过开发自测：是
是否经过测试验证： 否
是否对服务器有影响：否
`

var logInItamTpl = `
{
    "catalog_id":10,
    "service_id":11,
    "service_type":"运维资源",
    "fields":[
        {
            "type":"STRING",
            "id":369,
            "key":"title",
            "value":"{{.Msg}}",
            "choice":[

            ]
        },
        {
            "type":"SELECT",
            "id":401,
            "key":"bk_biz_id",
            "value":"{{.PjNO}}",
            "choice":[
                {
                    "key":"3",
                    "name":"冀云",
                    "can_delete":false
                },
                {
                    "key":"9",
                    "name":"陕西云",
                    "can_delete":false
                },
                {
                    "key":"10",
                    "name":"新甘肃云",
                    "can_delete":false
                },
                {
                    "key":"11",
                    "name":"赣州",
                    "can_delete":false
                }
            ]
        },
        {
            "type":"STRING",
            "id":370,
            "key":"SHENGJIGONGNENGMOKUAI",
            "value":"{{.ModuleNames}}",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":371,
            "key":"SHENGJIGONGNENGMIAOSHU",
            "value":"{{.Msg}}",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":372,
            "key":"SHIYONGSHUOMING",
            "value":"{{.Msg}}",
            "choice":[

            ]
        },
        {
            "type":"STRING",
            "id":373,
            "key":"BANBEN_YIDONGDUAN",
            "value":"",
            "choice":[

            ]
        },
        {
            "type":"STRING",
            "id":374,
            "key":"BANBEN_QIANDUAN",
            "value":"",
            "choice":[

            ]
        },
        {
            "type":"STRING",
            "id":375,
            "key":"BANBEN_HOUDUAN",
            "value":"",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":376,
            "key":"DUIYINGFUWUQI",
            "value":"{{.ModuleDesc}}对应的服务器",
            "choice":[

            ]
        },
        {
            "type":"INT",
            "id":377,
            "key":"YUJISHENGJISHICHANG",
            "value":10,
            "choice":[

            ]
        },
        {
            "type":"RADIO",
            "id":378,
            "key":"SHIFOUDUIFUWUQIYOUYINGXIANG",
            "value":"FOU",
            "choice":[
                {
                    "key":"SHI",
                    "name":"是"
                },
                {
                    "key":"FOU",
                    "name":"否"
                }
            ]
        },
        {
            "type":"TEXT",
            "id":379,
            "key":"YINGXIANGYEWUMOKUAI",
            "value":"{{.ModuleDesc}}",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":380,
            "key":"SHENGJIBUZHOU",
            "value":"部署升级命令：@{{.ModuleNames}},tag:{{.Tag}}",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":382,
            "key":"HUITUIJIHUA",
            "value":"升级前会对原工程进行备份",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":383,
            "key":"HUITUIFANGSHI",
            "value":"根据升级前的备份进行回滚",
            "choice":[

            ]
        },
        {
            "type":"TEXT",
            "id":384,
            "key":"HUITUIBUZHOU",
            "value":"启用备份内容",
            "choice":[

            ]
        },
        {
            "type":"INT",
            "id":385,
            "key":"HUITUISHICHANG",
            "value":10,
            "choice":[

            ]
        },
        {
            "type":"RADIO",
            "id":387,
            "key":"SHIFOUJINGGUOCESHIYANZHENG",
            "value":"SHI",
            "choice":[
                {
                    "key":"SHI",
                    "name":"是"
                },
                {
                    "key":"FOU",
                    "name":"否"
                }
            ]
        },
        {
            "type":"RADIO",
            "id":386,
            "key":"SHIFOUJINGGUOKAIFAZICE",
            "value":"SHI",
            "choice":[
                {
                    "key":"SHI",
                    "name":"是"
                },
                {
                    "key":"FOU",
                    "name":"否"
                }
            ]
        },
        {
            "type":"RADIO",
            "id":399,
            "key":"SHIFOUCHANPINYANZHENG",
            "value":"FOU",
            "choice":[
                {
                    "key":"SHI",
                    "name":"是"
                },
                {
                    "key":"FOU",
                    "name":"否"
                }
            ]
        },
        {
            "type":"RADIO",
            "id":400,
            "key":"SHIFOUXIANGMUYANZHENG",
            "value":"FOU",
            "choice":[
                {
                    "key":"SHI",
                    "name":"是"
                },
                {
                    "key":"FOU",
                    "name":"否"
                }
            ]
        },
        {
            "type":"FILE",
            "id":381,
            "key":"FUJIAN",
            "value":"",
            "choice":{

            }
        }
    ],
    "creator":"{{.Username}}",
    "attention":false
}
`
