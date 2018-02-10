package data_merge

import (
	"framework/render"
	"global"
	"models/back"
	"time"
)

type MesCenter struct {
	Key
	CdnUrl      string //
	SiteName    string
	NoticeTypes []back.NoticeTypes
}

func (m *MesCenter) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	m.SiteName = siteName
	if err != nil {
		return nil, err
	}
	times := new(global.Times)
	timeNow := global.GetCurrentTime()
	times.StartTime = timeNow - 7*24*3600
	times.EndTime = timeNow
	listdata, err := noticeBean.GetWapNoticeList(siteId, 0, times)
	var notice []back.NoticeTypes
	data := []back.NoticeTypes{
		{4, "视讯公告", nil},
		{5, "电子公告", nil},
		{6, "彩票公告", nil},
		{7, "体育公告", nil},
	}
	for _, v := range data {
		var noticeOne back.NoticeTypes
		noticeOne.NoticeCate = v.NoticeCate
		noticeOne.NoticeType = v.NoticeType
		for _, v1 := range listdata {
			if v.NoticeCate == v1.NoticeCate {
				var list back.WapSiteNoticeList
				list.NoticeDateStr = time.Unix(v1.NoticeDate, 10).Format("2006-01-02 15:04:05")
				list.NoticeCate = v1.NoticeCate
				list.Id = v1.Id
				list.NoticeContent = v1.NoticeContent
				list.NoticeTitle = v1.NoticeTitle
				list.NoticeAssign = v1.NoticeAssign
				noticeOne.NoticeList = append(noticeOne.NoticeList, list)
			}
		}
		notice = append(notice, noticeOne)
	}
	m.NoticeTypes = notice
	return m, nil
}

func (m *MesCenter) GetPage() []string {
	return []string{WAP_MESCENTER}
}
