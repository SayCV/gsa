// Copyright (c) 2015 by QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

// +build !windows

package ref

import (
  `os`
	`fmt`
	`time`
	`strconv`
	`io/ioutil`
	`net/http`
	`regexp`
	`strings`
	//stdLog `log`
	
	`github.com/SayCV/gsa/log`
	//`github.com/SayCV/gsa/util`
	`github.com/SayCV/gsa/portfolio/cons`
	
  //`golang.org/x/net/html`
  //`gopkg.in/xmlpath.v2`
  `github.com/beevik/etree`
)

func _fun_divi(x string) string {
  // var regex *regexp.Regexp
  regex := regexp.MustCompile(`分红(.*?)元`)
  res := regex.FindStringSubmatch(x)
  if len(res) < 1 { 
    return `0`
  } else { 
    return res[0]
  }
}

func _fun_into(x string) string {
  var res string
  regex := regexp.MustCompile(`转增(.*?)股`)
  res1 := regex.FindStringSubmatch(x)
  if len(res1) < 1 {
    res = `0`
  } else {
    res = res1[0]
  }
  res2 := regex.FindStringSubmatch(x)
  if len(res2) < 1 {
    res += `0`
  } else {
    res += res2[0]
  }
  return res
}

func _dist_cotent(year int, pageNo int, retry_count int, pause int) {
  for i:=0; i<retry_count; i++ {
    time.Sleep(time.Second * 1)
    defer func() {
  		if err := recover(); err != nil {
  			log.Error(cons.NETWORK_URL_ERROR_MSG, ` - `, err)
  		}
  	}()
  	if pageNo > 0 {
  	  cons.WriteConsole()
  	}
  	// fetch and read a web page
  	// http://quotes.money.163.com/data/caibao/fpyg.html?sort=declaredate&order=desc&reportdate=2014&sor
  	// http://quotes.money.163.com/data/caibao/fpyg.html?reportdate=2015&sort=declaredate&order=desc&page=0
  	// http://quotes.money.163.com/data/caibao/fpyg.html?reportdate=2015&sort=declaredate&order=desc&page=0
  	s := fmt.Sprintf(DP_163_URL, cons.P_TYPE[`http`], cons.DOMAINS[`163`],
                     cons.PAGES[`163dp`], strconv.Itoa(year), strconv.Itoa(pageNo))
  	log.Debug(s)
    resp, _ := http.Get(s)
    page, _ := ioutil.ReadAll(resp.Body)
    // avoid line 14 - no semicolon
    re := regexp.MustCompile("(?m)[\r\n]+^.*ntes_sitenav_link.*$")
    res := re.ReplaceAllString(string(page), "")
    re = regexp.MustCompile("(?m)[\r\n]+^.*noscript.*$")
    res = re.ReplaceAllString(res, "")
    ri := strings.NewReader(res)
    
    //log.Debug(ri)
    //log.Debug(util.GbkDecode(string(page)))
    //doc, err := html.Parse(ri)
    doc_fn_rp_list := etree.NewDocument()
    _, err := doc_fn_rp_list.ReadFrom(ri)
    //path := xmlpath.MustCompile(`//div[@class='fn_rp_list']/table`)
    //root, err := xmlpath.Parse(page)
    
  	if err != nil {
  		log.Emergency(err)
  	}
  	path_fn_rp_list := etree.MustCompilePath("//div[@class='fn_rp_list']")
  	//if value, ok := path.String(root); ok {
  	element_fn_rp_list := doc_fn_rp_list.FindElementPath(path_fn_rp_list)
  	//log.Info("Found:", element_fn_rp_list)
  	doc_table := etree.CreateDocument(element_fn_rp_list)
    //_, err := doc_table.ReadFromString()
  	path_table := etree.MustCompilePath("//table[@class='fn_cm_table']")
  	element_table := doc_table.FindElementPath(path_table)
  	list := etree.CreateDocument(element_table)
  	list.Indent(2)
  	list.WriteTo(os.Stdout)
		/*elements_table := doc_table.FindElementsPath(path_table)
  	for _, e := range elements_table {
      //log.Info("Found:", e.Text())
      list := etree.CreateDocument(e)
      list.Indent(2)
		  list.WriteTo(os.Stdout)
    }*/
    //sarr := etree.NewDocument()
}
}
/*
            html = lxml.html.parse(fmt.Sprintf(DP_163_URL, cons.P_TYPE['http'], cons.DOMAINS['163'],
                     cons.PAGES['163dp'], strconv.Itoa(year), strconv.Itoa(pageNo)))  
            res = html.xpath('//div[@class=\"fn_rp_list\"]/table')
            if ct.PY3:
                sarr = [etree.tostring(node).decode('utf-8') for node in res]
            else:
                sarr = [etree.tostring(node) for node in res]
            sarr = ''.join(sarr)
            df = pd.read_html(sarr, skiprows=[0])[0]
            df = df.drop(df.columns[0], axis=1)
            df.columns = rv.DP_163_COLS
            df['divi'] = df['plan'].map(_fun_divi)
            df['shares'] = df['plan'].map(_fun_into)
            df = df.drop('plan', axis=1)
            df['code'] = df['code'].astype(object)
            df['code'] = df['code'].map(lambda x : str(x).zfill(6))
            pages = []
            if pageNo == 0:
                page = html.xpath('//div[@class=\"mod_pages\"]/a')
                if len(page)>1:
                    asr = page[len(page)-2]
                    pages = asr.xpath('text()')
        except Exception as e:
            print(e)
        else:
            if pageNo == 0:
                return df, pages[0] if len(pages)>0 else 0
            else:
                return df
    raise IOError(ct.NETWORK_URL_ERROR_MSG)
}
*/

/*
        获取业绩预告数据
    Parameters
    --------
    year:int 年度 e.g:2014
    quarter:int 季度 :1、2、3、4，只能输入这4个季度
       说明：由于是从网站获取的数据，需要一页页抓取，速度取决于您当前网络速度
       
    Return
    --------
    DataFrame
        code,代码
        name,名称
        type,业绩变动类型【预增、预亏等】
        report_date,发布日期
        pre_eps,上年同期每股收益
        range,业绩变动范围
        
*/
/*
func profit_data(year int, top int, retry_count int, pause float32) pandas {
  if top <= 25:
        df, pages = _dist_cotent(year, 0, retry_count, pause)
        return df.head(top)
    elif top == 'all':
        ct._write_head()
        df, pages = _dist_cotent(year, 0, retry_count, pause)
        for idx in range(1,int(pages)):
            df = df.append(_dist_cotent(year, idx, retry_count,
                                        pause), ignore_index=True)
        return df
    else:
        if isinstance(top, int):
            ct._write_head()
            allPages = top/25+1 if top%25>0 else top/25
            df, pages = _dist_cotent(year, 0, retry_count, pause)
            if int(allPages) < int(pages):
                pages = allPages
            for idx in range(1, int(pages)):
                df = df.append(_dist_cotent(year, idx, retry_count,
                                            pause), ignore_index=True)
            return df.head(top)
        else:
            print(ct.TOP_PARAS_MSG)
  return 1
}
*/

func Test() {
  _dist_cotent(2015, 0, 2, 0)
}