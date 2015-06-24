// Copyright (c) 2015 by QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	`fmt`
)

// var VERSION = `0.1.0`
var K_LABELS = [...]string{`D`, `W`, `M`}
var K_MIN_LABELS = [...]string{`5`, `15`, `30`, `60`}
var K_TYPE = map[string]string{`D`: `akdaily`, `W`: `akweekly`, `M`: `akmonthly`}
var INDEX_LABELS = [...]string{`sh`, `sz`, `hs300`, `sz50`, `cyb`, `zxb`}
var INDEX_LIST = map[string]string{`sh`: `sh000001`, `sz`: `sz399001`, `hs300`: `sz399300`,
              `sz50`: `sh000016`, `zxb`: `sz399005`, `cyb`: `sz399006`}
var P_TYPE = map[string]string{`http`: `http://`, `ftp`: `ftp://`}
var PAGE_NUM = [...]uint8{38, 60, 80, 100}
var FORMAT = `lambda x: %.2f % x`
var DOMAINS = map[string]string{
           `sina`: `sina.com.cn`, `sinahq`: `sinajs.cn`,
           `ifeng`: `ifeng.com`, `sf`: `finance.sina.com.cn`,
           `vsf`: `vip.stock.finance.sina.com.cn`, 
           `idx`: `www.csindex.com.cn`, `163`: `money.163.com`,
           `em`: `eastmoney.com`, `sseq`: `query.sse.com.cn`,
           `sse`: `www.sse.com.cn`, `szse`: `www.szse.cn`,
           `oss`: `tudata.oss-cn-beijing.aliyuncs.com`,
           `shibor`: `www.shibor.org`}
var PAGES = map[string]string{`fd`: `index.phtml`, `dl`: `downxls.php`, `jv`: `json_v2.php`,
         `cpt`: `newFLJK.php`, `ids`: `newSinaHy.php`, `lnews`:`rollnews_ch_out_interface.php`,
         `ntinfo`:`vCB_BulletinGather.php`, `hs300b`:`000300cons.xls`,
         `hs300w`:`000300closeweight.xls`,`sz50b`:`000016cons.xls`,
         `dp`:`all_fpya.php`, `163dp`:`fpyg.html`,
         `emxsg`:`JS.aspx`, `163fh`:`jjcgph.php`,
         `newstock`:`vRPD_NewStockIssue.php`, `zz500b`:`000905cons.xls`,
         `t_ticks`:`vMS_tradedetail.php`, `dw`: `downLoad.html`,
         `qmd`:`queryMargin.do`, `szsefc`:`FrontController.szse`,
         `ssecq`:`commonQuery.do`}
var TICK_COLUMNS = [...]string{`time`, `price`, `change`, `volume`, `amount`, `type`}
var TODAY_TICK_COLUMNS = [...]string{`time`, `price`, `pchange`, `change`, `volume`, `amount`, `type`}
var DAY_TRADING_COLUMNS = [...]string{`code`, `symbol`, `name`, `changepercent`,
                       `trade`, `open`, `high`, `low`, `settlement`, `volume`, `turnoverratio`}
var REPORT_COLS = [...]string{`code`, `name`, `eps`, `eps_yoy`, `bvps`, `roe`,
               `epcf`, `net_profits`, `profits_yoy`, `distrib`, `report_date`}
var FORECAST_COLS = [...]string{`code`, `name`, `type`, `report_date`, `pre_eps`, `range`}
var PROFIT_COLS = [...]string{`code`, `name`, `roe`, `net_profit_ratio`,
               `gross_profit_rate`, `net_profits`, `eps`, `business_income`, `bips`}
var OPERATION_COLS = [...]string{`code`, `name`, `arturnover`, `arturndays`, `inventory_turnover`,
                  `inventory_days`, `currentasset_turnover`, `currentasset_days`}
var GROWTH_COLS = [...]string{`code`, `name`, `mbrg`, `nprg`, `nav`, `targ`, `epsg`, `seg`}
var DEBTPAYING_COLS = [...]string{`code`, `name`, `currentratio`,
                   `quickratio`, `cashratio`, `icratio`, `sheqratio`, `adratio`}
var CASHFLOW_COLS = [...]string{`code`, `name`, `cf_sales`, `rateofreturn`,
                 `cf_nm`, `cf_liabilities`, `cashflowratio`}
var DAY_PRICE_COLUMNS = [...]string{`date`, `open`, `high`, `close`, `low`, `volume`, `price_change`, `p_change`,
                     `ma5`, `ma10`, `ma20`, `v_ma5`, `v_ma10`, `v_ma20`, `turnover`}
var INX_DAY_PRICE_COLUMNS = [...]string{`date`, `open`, `high`, `close`, `low`, `volume`, `price_change`, `p_change`,
                         `ma5`, `ma10`, `ma20`, `v_ma5`, `v_ma10`, `v_ma20`}
var LIVE_DATA_COLS = [...]string{`name`, `open`, `pre_close`, `price`, `high`, `low`, `bid`, `ask`, `volume`, `amount`,
                  `b1_v`, `b1_p`, `b2_v`, `b2_p`, `b3_v`, `b3_p`, `b4_v`, `b4_p`, `b5_v`, `b5_p`,
                  `a1_v`, `a1_p`, `a2_v`, `a2_p`, `a3_v`, `a3_p`, `a4_v`, `a4_p`, `a5_v`, `a5_p`, `date`, `time`, `s`}
var FOR_CLASSIFY_B_COLS = [...]string{`code`,`name`}
var FOR_CLASSIFY_W_COLS = [...]string{`date`,`code`,`weight`}
var THE_FIELDS = [...]string{`code`,`symbol`,`name`,`changepercent`,`trade`,`open`,`high`,`low`,`settlement`,`volume`,`turnoverratio`}
var TICK_PRICE_URL = `%smarket.%s/%s?date=%s&symbol=%s`
var TODAY_TICKS_PAGE_URL = `%s%s/quotes_service/api/%s/CN_Transactions.getAllPageTime?date=%s&symbol=%s`
var TODAY_TICKS_URL = `%s%s/quotes_service/view/%s?symbol=%s&date=%s&page=%s`
var DAY_PRICE_URL = `%sapi.finance.%s/%s/?code=%s&type=last`
var LIVE_DATA_URL = `%shq.%s/rn=%s&list=%s`
var DAY_PRICE_MIN_URL = `%sapi.finance.%s/akmin?scode=%s&type=%s`
var SINA_DAY_PRICE_URL = `%s%s/quotes_service/api/%s/Market_Center.getHQNodeData?num=80&sort=changepercent&asc=0&node=hs_a&symbol=&_s_r_a=page&page=%s`
var REPORT_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/mainindex/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var FORECAST_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/performance/%s?s_i=&s_a=&s_c=&s_type=&reportdate=%s&quarter=%s&p=%s&num=%s`
var PROFIT_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/profit/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var OPERATION_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/operation/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var GROWTH_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/grow/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var DEBTPAYING_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/debtpaying/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var CASHFLOW_URL = `%s%s/q/go.php/vFinanceAnalyze/kind/cashflow/%s?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s`
var SHIBOR_TYPE = map[string]string{`Shibor`: `Shibor数据`, `Quote`: `报价数据`, `Tendency`: `Shibor均值数据`,
              `LPR`: `LPR数据`, `LPR_Tendency`: `LPR均值数据`}
var SHIBOR_DATA_URL = `%s%s/shibor/web/html/%s?nameNew=Historical_%s_Data_%s.xls&downLoadPath=data&nameOld=%s%s.xls&shiborSrc=http://www.shibor.org/shibor/`
var ALL_STOCK_BASICS_FILE = fmt.Sprintf(`%s%s/all.csv`, P_TYPE[`http`], DOMAINS[`oss`])
var SINA_CONCEPTS_INDEX_URL = `%smoney.%s/q/view/%s?param=class`
var SINA_INDUSTRY_INDEX_URL = `%s%s/q/view/%s`
var SINA_DATA_DETAIL_URL = `%s%s/quotes_service/api/%s/Market_Center.getHQNodeData?page=1&num=400&sort=symbol&asc=1&node=%s&symbol=&_s_r_a=page`
var INDEX_C_COMM = `sseportal/ps/zhs/hqjt/csi`
var HS300_CLASSIFY_URL = `%s%s/%s/%s`
var HIST_FQ_URL = `%s%s/corp/go.php/vMS_FuQuanMarketHistory/stockid/%s.phtml?year=%s&jidu=%s`
var HIST_INDEX_URL = `%s%s/corp/go.php/vMS_MarketHistory/stockid/%s/type/S.phtml?year=%s&jidu=%s`
var HIST_FQ_FACTOR_URL = `%s%s/api/json.php/BasicStockSrv.getStockFuQuanData?symbol=%s&type=hfq`
var INDEX_HQ_URL = `%shq.%s/rn=xppzh&list=sh000001,sh000002,sh000003,sh000008,sh000009,sh000010,sh000011,sh000012,sh000016,sh000017,sh000300,sz399001,sz399002,sz399003,sz399004,sz399005,sz399006,sz399100,sz399101,sz399106,sz399107,sz399108,sz399333,sz399606`
var SSEQ_CQ_REF_URL = `%s%s/assortment/stock/list/name`
var ALL_STK_URL = `%s%s/all.csv`
var SHIBOR_COLS = [...]string{`date`, `ON`, `1W`, `2W`, `1M`, `3M`, `6M`, `9M`, `1Y`}
var QUOTE_COLS = [...]string{`date`, `bank`, `ON_B`, `ON_A`, `1W_B`, `1W_A`, `2W_B`, `2W_A`, `1M_B`, `1M_A`,
                    `3M_B`, `3M_A`, `6M_B`, `6M_A`, `9M_B`, `9M_A`, `1Y_B`, `1Y_A`}
var SHIBOR_MA_COLS = [...]string{`date`, `ON_5`, `ON_10`, `ON_20`, `1W_5`, `1W_10`, `1W_20`,`2W_5`, `2W_10`, `2W_20`,
                  `1M_5`, `1M_10`, `1M_20`, `3M_5`, `3M_10`, `3M_20`, `6M_5`, `6M_10`, `6M_20`,
                  `9M_5`, `9M_10`, `9M_20`,`1Y_5`, `1Y_10`, `1Y_20`}
var LPR_COLS = [...]string{`date`, `1Y`}
var LPR_MA_COLS = [...]string{`date`, `1Y_5`, `1Y_10`, `1Y_20`}
var INDEX_HEADER = `code,name,open,preclose,close,high,low,0,0,volume,amount,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,d,c,3\n`
var INDEX_COLS = [...]string{`code`, `name`, `change`, `open`, `preclose`, `close`, `high`, `low`, `volume`, `amount`}
var HIST_FQ_COLS = [...]string{`date`, `open`, `high`, `close`, `low`, `volume`, `amount`, `factor`}
var HIST_FQ_FACTOR_COLS = [...]string{`code`,`value`}
var DATA_GETTING_TIPS = `string{Getting data:}`
var DATA_GETTING_FLAG = `#`
var DATA_ROWS_TIPS = `%s rows data found.Please wait for a moment.`
var DATA_INPUT_ERROR_MSG = `date input error.`
var NETWORK_URL_ERROR_MSG = `获取失败，请检查网络和URL`
var DATE_CHK_MSG = `年度输入错误：请输入1989年以后的年份数字，格式：YYYY`
var DATE_CHK_Q_MSG = `季度输入错误：请输入1、2、3或4数字`
var TOP_PARAS_MSG = `top有误，请输入整数或all.`
var LHB_MSG = `周期输入有误，请输入数字5、10、30或60`