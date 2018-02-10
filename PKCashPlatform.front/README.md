#### 现金网前端
```
|-- smartadmin-plugin         # 框架插件
|
|-- src                         #写代码的位置
|   |-- styles                   #样式图片
|   |-- smart                   #框架的代码
|   |   |-- _common             # SmartAdmin module
|   |       |-- module.js       #
|   |
|   |   |-- layout              # app.layout module
|   |   ...
|   |-- admin                   #总后台的代码
|   |   ...
|   |   |-- app.js                  # 程序入口函数
|   |   |-- app.service.config.js     # 全局静态常量中的说有用的的api地址
|   |   |-- service                   # 接口封装文件夹
|   |   |-- dashboard                 # 仪表盘模块
|   |   |-- Platform                  # 管理员管理
|   |   |-- PermissionManagement      # 操作记录、站点栏目、日志管理
|   |   |-- CopyTemplate              # 文案模板管理
|   |   |-- site                      # 站点管理
|   |   |-- customer                  # 客户后台管理
|   |   |-- ReportForm                # 财务报表
|   |   ...
|   |-- agency                   #客户后台的代码
|   |   ...
|   |   |-- app.js                  # 程序入口函数
|   |   |-- app.service.config.js     # 全局静态常量中的说有用的的api地址
|   |   |-- service                   # 接口封装文件夹
|   |   |-- dashboard                 # 仪表盘模块
|   |   |-- account                   # 账号管理模块
|   |       |-- administrators         # 会员账号管理
|   |   |-- moneyManagement            # 资金管理
|   |       |-- accessManagement            # 出入款管理
|   |           |-- AccessMoney             # 人工存款、人工取款、出入款账目汇总、额度转换、额度记录
|   |           |-- Income                  # 公司入款、线上入款
|   |           |-- CashManagement          # 出款管理
|   |           |-- Summary                 # 出入账目汇总
|   |       |-- auditManagement        # 稽核管理
|   |           |-- AuditQuery              # 即使稽核查询
|   |           |-- AuditLog                # 稽核日志
|   |       |-- paymentSetting         # 支付设定
|   |       |-- analysisExits          # 会员分析
|   |           |-- analysisExit            # 出入款分析、购买分析、退水分析、有效会员列表
|   |           |-- cashSystem              # 现金流水
|   |           |-- balanceStatistics       # 会员余额统计
|   |       |-- commission             # 退佣统计模块
|   |           |-- commissionStatistics    # 退佣统计、代理退佣设定、期数管理、手续费设定、退佣查询
|   |       |-- preferentialTerms      # 优惠计算
|   |           |-- PreferentialCalculation   # 返点优惠设定、优惠查询、优惠统计、自助返水查询
|   |       |-- membership             # 会员返佣
|   |           |-- membershipReturns         # 返佣查询、返佣优惠设定、会员返佣、会员推广查询、会员推广设定
|   |       |-- amountStatistics       # 额度统计
|   |           |-- quotaStatistics           # 额度统计、充值记录、掉单列表、额度充值、额度记录
|   |       |-- reportManagement        # 报表管理
|   |           |-- reportForm                # 报表查询
|   |       |-- webInformation          # 网站资讯管理
|   |           |-- websiteManagement         # 站点资料编辑
|   |           |-- graphicEditor             # 图片编辑
|   |           |-- copyEditor                # 文件编辑、案件编辑
|   |       |-- systemSetup             # 系统设置
|   |           |-- notice                    # 公告管理
|   |           |-- memberMessage             # 会员消息
|   |           |-- announcement              # 最新消息
|   |           |-- operation                 # 日志管理
|   |   ...
|   |-- httpSvc.js                   #常用类
...
|-- app.config.js             # 全局静态常量 （constant）
|-- app.scripts.json          # vendor管理
...
```