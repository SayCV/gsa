// Copyright (c) 2015 by QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

var DP_URL = `%sapp.finance.%s/data/stock/%s?day=&page=%s`
var DP_163_URL = `%squotes.%s/data/caibao/%s?reportdate=%s&sort=declaredate&order=desc&page=%s`
var FUND_HOLDS_URL = `%squotes.%s/hs/marketdata/service/%s?host=/hs/marketdata/service/%s&page=%s&query=start:%s;end:%s&order=desc&count=60&type=query&req=%s`
var XSG_URL = `%sdatainterface.%s/EM_DataCenter/%s?type=FD&sty=BST&st=3&sr=true&fd=%s&stat=%s`
var LHB_URL = `%sdata.%s/stock/lhb/%s.html`
var LHB_SINA_URL = `%s%s/q/go.php/vLHBData/kind/%s/%s?last=%s&p=%s`
var LHB_COLS = [...]string{`code`, `name`, `pchange`, `amount`, `buy`, `bratio`, `sell`, `sratio`, `reason`}
var NEW_STOCKS_URL = `%s%s/corp/view/%s?page=%s&cngem=0&orderBy=NetDate&orderType=desc`
var MAR_SH_HZ_URL = `%s%s/marketdata/tradedata/%s?jsonCallBack=jsonpCallback%s&isPagination=true&tabType=&pageHelp.pageSize=100&beginDate=%s&endDate=%s%s&_=%s`
var MAR_SH_HZ_REF_URL = `%s%s/market/dealingdata/overview/margin/`
var MAR_SH_MX_URL = `%s%s/marketdata/tradedata/%s?jsonCallBack=jsonpCallback%s&isPagination=true&tabType=mxtype&detailsDate=%s&pageHelp.pageSize=100&stockCode=%s&beginDate=%s&endDate=%s%s&_=%s`
var MAR_SZ_HZ_URL = `%s%s/szseWeb/%s?ACTIONID=8&CATALOGID=1837_xxpl&txtDate=%s&tab2PAGENUM=1&ENCODE=1&TABKEY=tab1`
var MAR_SZ_MX_URL = `%s%s/szseWeb/%s?ACTIONID=8&CATALOGID=1837_xxpl&txtDate=%s&tab2PAGENUM=1&ENCODE=1&TABKEY=tab2`
var MAR_SH_HZ_TAIL_URL = `&pageHelp.pageNo=%s&pageHelp.beginPage=%s&pageHelp.endPage=%s`
var TERMINATED_URL = `%s%s/%s?jsonCallBack=jsonpCallback%s&isPagination=true&sqlId=COMMON_SSE_ZQPZ_GPLB_MCJS_ZZSSGGJBXX_L&pageHelp.pageSize=50&_=%s`
var SUSPENDED_URL = `%s%s/%s?jsonCallBack=jsonpCallback%s&isPagination=true&sqlId=COMMON_SSE_ZQPZ_GPLB_MCJS_ZTSSGS_L&pageHelp.pageSize=50&_=%s`
var TERMINATED_T_COLS = [...]string{`COMPANY_CODE`, `COMPANY_ABBR`, `LISTING_DATE`, `CHANGE_DATE`}
var LHB_KINDS = [...]string{`ggtj`, `yytj`, `jgzz`, `jgmx`}
var LHB_GGTJ_COLS = [...]string{`code`, `name`, `count`, `bamount`, `samount`, `net`, `bcount`, `scount`}
var LHB_YYTJ_COLS = [...]string{`broker`, `count`, `bamount`, `bcount`, `samount`, `scount`, `top3`}
var LHB_JGZZ_COLS = [...]string{`code`, `name`, `bamount`, `bcount`, `samount`, `scount`, `net`}
var LHB_JGMX_COLS = [...]string{`code`, `name`, `date`, `bamount`, `samount`, `type`}
var TERMINATED_COLS = [...]string{`code`, `name`, `oDate`, `tDate`}
var DP_COLS = [...]string{`report_date`, `quarter`, `code`, `name`, `plan`}
var DP_163_COLS = [...]string{`code`, `name`, `year`, `plan`, `report_date`}
var XSG_COLS = [...]string{`code`, `name`, `date`, `count`, `ratio`}
var QUARTS_DIC = map[string][2]string{`1`:{`%s-12-31`, `%s-03-31`}, `2`:{`%s-03-31`, `%s-06-30`}, 
              `3`:{`%s-06-30`, `%s-09-30`}, `4`:{`%s-9-30`, `%s-12-31`}}
var FUND_HOLDS_COLS = [...]string{`count`, `clast`, `date`, `ratio`, `amount`, `nums`,`nlast`, `name`, `code`}
var NEW_STOCKS_COLS = [...]string{`code`, `name`, `ipo_date`, `issue_date`, `amount`, `markets`, `price`, `pe`,
                   `limit`, `funds`, `ballot`}
var MAR_SH_COOKIESTR = `_gscu_1808689395=27850607moztu036`
var MAR_SH_HZ_COLS = [...]string{`opDate`, `rzye`, `rzmre`, `rqyl`, `rqylje`, `rqmcl`, `rzrqjyzl`}
var MAR_SH_MX_COLS = [...]string{`opDate`, `stockCode`, `securityAbbr`, `rzye`, `rzmre`, `rzche`, `rqyl`, `rqmcl`, `rqchl`}
var MAR_SZ_HZ_COLS = [...]string{`rzmre`, `rzye`, `rqmcl`, `rqyl`, `rqye`, `rzrqye`}
var MAR_SZ_MX_COLS = [...]string{`stockCode`, `securityAbbr`, `rzmre`, `rzye`, `rqmcl`, `rqyl`, `rqye`, `rzrqye`}
var MAR_SZ_HZ_MSG = `please do not input more than a year,you can obtaining the data year by year.`
var MAR_SZ_HZ_MSG2 = `start and end date all need input.`