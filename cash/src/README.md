### 前端测试地址
> 代理管理系统 http://10.10.10.186:8080<br/>
> 平台管理系统 http://10.10.10.186:8088

### api测试地址
> 代理管理系统 http://10.10.10.186:9696<br/>
> 平台管理系统 http://10.10.10.186:9797

### 获取测试版本信息接口地址
> 代理管理系统 http://10.10.10.186:9696/version/test<br/>
> 平台管理系统 http://10.10.10.186:9797/version/test

### api登录账号
```
1- whw_k whw                        开户人
    a(默认)                          站点
        5- whw_a_g_1(默认)           股东
            11- whw_a_z_1(默认)      总代理
                30- whw_a_d_1(默认)  代理
                    1- whw_a_m_1     会员
                    2- whw_a_m_2     会员
                    3- whw_a_m_3     会员
                31- whw_a_d_2        代理
                    4- whw_a_m_4     会员
                    5- whw_a_m_5     会员
                    6- whw_a_m_6     会员
                32- whw_a_d_3        代理
            12- whw_a_z_2            总代理
            13- whw_a_z_3            总代理
            14- whw_a_z_4            总代理
        6- whw_a_g_2                 股东
            15- whw_a_z_a            总代理
            16- whw_a_z_b            总代理
            17- whw_a_z_c            总代理
        7- whw_a_g_3                 股东
            18- whw_a_z_A            总代理
            19- whw_a_z_B            总代理
            20- whw_a_z_C            总代理
        8- whw_a_g_4                 股东
            21- whw_a_z_11            总代理
            22- whw_a_z_22            总代理
            23- whw_a_z_33            总代理
        9- whw_a_g_5                 股东
            24- whw_a_z_aa            总代理
            25- whw_a_z_bb            总代理
            26- whw_a_z_cc            总代理
        10- whw_a_g_6                 股东
            27- whw_a_z_AA            总代理
            28- whw_a_z_BB            总代理
            29- whw_a_z_CC            总代理
    b                                 站点
        33- whw_b_g_1(默认)           股东
            36- whw_b_z_1(默认)       总代理
                39- whw_b_d_1(默认)   代理
                40- whw_b_d_2         代理
                41- whw_b_d_3         代理
            37- whw_b_z_2             总代理
            38- whw_b_z_3             总代理
        34- whw_b_g_2                 股东
        35- whw_b_g_3                 股东
2- wt_k wt                            开户人
    a(默认)                           站点
        42- wt_a_g_1(默认)            股东
            44- wt_a_z_1(默认)        总代理
                46- wt_a_d_1(默认)    代理
                    7- wt_a_m_1       会员
                    8- wt_a_m_2       会员
                47- wt_a_d_2          代理
                    9- wt_a_m_3       会员
                    10- wt_a_m_4      会员
            45- wt_a_z_2             总代理
        43- wt_a_g_2
3- lmg_k lmg
4- zym_k zmy
```

### bug文档书写规范

#### 格式

标题样式:

<类型>-<状态>-<功能模块>

例如:

bug-紧急-会员列表

opt-一般-会员列表下拉框

```
版本:xxxx.xx
路由:xxxx
问题:xxxx
补充描述:xxxx
```
#### 例子
```
版本:170927.01
路由:http://10.10.10.186:9696/member
问题:返回数据缺少一个会员订单数量字段
补充描述:
```