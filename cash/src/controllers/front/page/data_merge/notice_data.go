package data_merge

import (
	"framework/render"
	"time"
)

type NoticeData struct {
	Key
	CdnUrl string        //cdn rui
	Notice []*NoticeInfo //公告内容
}

type NoticeInfo struct {
	NoticeDate    string //公告时间
	NoticeContent string //公告内容
}

//得到数据
func (m *NoticeData) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//查询公告
	notices, err := noticeBean.GetNoticeAndDateBySiteId(siteId)
	if err != nil {
		return nil, err
	}
	if len(notices) > 0 {
		for k, _ := range notices {
			tm := time.Unix(notices[k].NoticeDate, 0)
			tm.Format("2006-01-02 03:04:05")
			m.Notice = append(m.Notice, &NoticeInfo{tm.String()[:19], notices[k].NoticeContent})
		}
	}
	return m, nil
}

//得到页面
func (m *NoticeData) GetPage() []string {
	return []string{NOTICE_DATA}
}

//得到<视讯电子彩票体育的页面>
func (m *NoticeData) GetSubPage() map[string]string {
	return nil
}
