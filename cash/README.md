go run build.go --goos=linux build front && go run build.go --goos=linux build server && go run build.go --goos=linux build admin && go run build.go --goos=linux build wap


go run build.go --environment=production --goos=linux build front && go run build.go --environment=production --goos=linux build server && go run build.go --environment=production --goos=linux build admin && go run build.go --environment=production --goos=linux build wap

production

wiki太卡,readme里编写改动



1. 统计每日优惠`sales_discount_report`表修改记录
    
    |修改列|改动|
    |:---|:---|
    |ua_id|second_agency_id|
    |sh_id|first_agency_id|
    |cash_type|删除|
    |day_type `int`|day_type `varchar`|
    
1. 统计每日出入款 `sales_cash_count_report` 表修改记录
    
    |修改列|改动|
    |:---|:---|
    |ua_id|second_agency_id|
    |sh_id|first_agency_id|
    |cash_type|删除|
    |day_type `int`|day_type `varchar`|
    |cash_num|cash_money|

1. 公司入款记录`sales_member_company_income` 表修改记录
    
    |修改列|改动|
    |:---|:---|
    |user_name|account|
    
1. `sales_online_entry_record`线上入款列表改为线上入款记录表    
