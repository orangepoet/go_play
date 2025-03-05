package main

var content1 string = `
{
  "code": 200,
  "status": 200,
  "msg": "成功",
  "data": [
    {
      "creator": "chenxiaoyu",
      "materialId": 1245,
      "materialName": "通用商品弹窗-跳会场",
      "materialType": 13,
      "materialTypeDesc": "通用商品弹窗",
      "img": "",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "o_zhouhui",
      "materialId": 1236,
      "materialName": "ZH-1-居底-首单补贴",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10612097/20250210-27f6998459bf7a6b-w990h396.jpeg",
      "quoteStatus": 0,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "songhuan01",
      "materialId": 1122,
      "materialName": "testtest2",
      "materialType": 13,
      "materialTypeDesc": "通用商品弹窗",
      "img": "",
      "quoteStatus": 0,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "chenxiaoyu",
      "materialId": 1119,
      "materialName": "通用商品弹窗",
      "materialType": 13,
      "materialTypeDesc": "通用商品弹窗",
      "img": "",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "tianyu01",
      "materialId": 1118,
      "materialName": "ty-1234",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10119464/20250124-79a30cf90420e063-w963h1290.png",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "songhuan01",
      "materialId": 1116,
      "materialName": "1111",
      "materialType": 13,
      "materialTypeDesc": "通用商品弹窗",
      "img": "",
      "quoteStatus": 0,
      "isAlgoSync": 1,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "tianyu01",
      "materialId": 1114,
      "materialName": "ty-1",
      "materialType": 13,
      "materialTypeDesc": "通用商品弹窗",
      "img": "",
      "quoteStatus": 0,
      "isAlgoSync": 1,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "qiaomingming01",
      "materialId": 1005,
      "materialName": "通用图片弹窗素材 test",
      "materialType": 1,
      "materialTypeDesc": "通用图片弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10136210/20241206-6228108923bbc86a-w810h1116.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 1,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "zhaoguangcheng",
      "materialId": 989,
      "materialName": "520弹窗登录注册页面（会场专用 慎动）",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10711758/20241129-4721d9a5694fedc2-w963h1290.png",
      "quoteStatus": 0,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "xuanjun",
      "materialId": 885,
      "materialName": "520领券_居底（xj）",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10657112/20240306-62f441f5c3061099-w990h396.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "",
      "materialId": 470,
      "materialName": "通用图片弹窗_52",
      "materialType": 1,
      "materialTypeDesc": "通用图片弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10114019/20230706-a00ace2587e6fec9-w810h1116.png",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "zhangmingxin01",
      "materialId": 401,
      "materialName": "520领券_居底（临时）",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10657112/20240306-62f441f5c3061099-w990h396.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "tangxinjing",
      "materialId": 195,
      "materialName": "txj测试",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10110907/20241218-5805554e3f96fe58-w963h1290.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "chenxiaoyu",
      "materialId": 120,
      "materialName": "抽免单",
      "materialType": 1,
      "materialTypeDesc": "通用图片弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/jGPb4wZ3/20230901-a3470f7a9c1da9f2-w810h1116.png",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "chenxiaoyu",
      "materialId": 28,
      "materialName": "520领券_居底_自动化勿动",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10657112/20240304-5334a9ae2bead8bc-w990h396.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "chenxiaoyu",
      "materialId": 27,
      "materialName": "520领券_居中_自动化勿动",
      "materialType": 6,
      "materialTypeDesc": "520领券弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10114019/20230711-bd837dc9c2d42f75-w963h1290.jpeg",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期",
      "couponStatus": "0"
    },
    {
      "creator": "chenxiaoyu",
      "materialId": 26,
      "materialName": "通用弹窗",
      "materialType": 1,
      "materialTypeDesc": "通用图片弹窗",
      "img": "https://t1-h5cdn.dewu.net/growth/growth-react/10114019/20230706-a00ace2587e6fec9-w810h1116.png",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    },
    {
      "creator": "",
      "materialId": 5,
      "materialName": "通用弹窗_自动化勿动",
      "materialType": 1,
      "materialTypeDesc": "通用图片弹窗",
      "img": "https://cdn.poizon.com/node-common/TGFyazIwMjAxMTE5LTE3MzgxNDE2MDU4NTM3NDY1MDI=.gif",
      "quoteStatus": 1,
      "isAlgoSync": 0,
      "couponStatusText": "未领取/已过期,券在期",
      "couponStatus": "0,1"
    }
  ],
  "timestamp": 1740540061411,
  "traceId": "0a58743b67be889dfd0b03d0b05cefce"
}
	`

var content2 string = `
	{
  "code": 200,
  "status": 200,
  "msg": "已过期",
  "data": [
    {
      "date": "2025-03-09",
      "availableCount": 2000
    },
    {
      "date": "2025-03-10",
      "availableCount": 1999
    },
    {
      "date": "2025-03-11",
      "availableCount": 2000
    },
    {
      "date": "2025-03-12",
      "availableCount": 1999
    },
    {
      "date": "2025-03-13",
      "availableCount": 2000
    },
    {
      "date": "2025-03-14",
      "availableCount": 2000
    },
    {
      "date": "2025-03-15",
      "availableCount": 2000
    }
  ],
  "timestamp": 1740518242185,
  "traceId": "0a4882a167be3362e2372d43586c7df4"
}
	`
