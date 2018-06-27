package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"io/ioutil"
	"github.com/json-iterator/go"
	xj "github.com/basgys/goxml2json"
	"strings"
	"shanghai/utils"
	"shanghai/models"
	"github.com/PuerkitoBio/goquery"
)

var jsonf = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	//依次执行
	//抓取xml线路
	//CrawLineNameList()
	//按照线路获取line_id等信息,同时插入updown表
	//CrawLineIdUpdownList()
	//获取线路站台
    //CrawLineStation()
    //获取多少条线路
    CrawLines()
	ExampleNewDocumentFromReader_string()
    //获取线路信息
    //CrawLineStation()



}

func CrawLines() {

	res, err := http.Get("http://chengdu.8684.cn/list1")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	fmt.Println("res.Body is",res.Body)
	doc,err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Find the review items
	doc.Find("con_site_1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})

	// Find the review items
	/*doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})*/

	/*doc.Find("con_site_1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Println("s is",s.Text())
		//band := s.Find("h4").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})*/
}

func ExampleNewDocumentFromReader_string() {
	// create from a string
	data := `
<html>
	<head>
		<title>My document</title>
	</head>
	<body>
		<h1>Header</h1>
<div id="con_site_1" class="stie_list"><a href="/x_aee6271d">118路</a><a href="/x_24f5dad9">1路</a><a href="/x_8234e473">10路</a><a href="/x_94d65922">11路</a><a href="/x_27905c0e">12路</a><a href="/x_968a4331">16路</a><a href="/x_bfe4a3eb">18路</a><a href="/x_74a4f39a">19路</a><a href="/x_cffdf62f">159路</a><a href="/x_bc14618a">104路</a><a href="/x_2a9fc0ea">1115路</a><a href="/x_2218ab25">115路</a><a href="/x_4e4038ec">100路</a><a href="/x_6674c2bb">101路</a><a href="/x_87a31151">106路</a><a href="/x_face82cc">109路</a><a href="/x_94192360">112路</a><a href="/x_d0f81574">113路</a><a href="/x_b883dfef">114路</a><a href="/x_e219274c">116路</a><a href="/x_a555cd66">152路</a><a href="/x_163cda33">137路</a><a href="/x_0e74b791">136路</a><a href="/x_92abb355">105路</a><a href="/x_c431c54a">103路</a><a href="/x_361efa89">1089路</a><a href="/x_a53656c4">139路</a><a href="/x_017a31d4">123路</a><a href="/x_4177a33a">17路</a><a href="/x_d5ae2b7f">168路</a><a href="/x_67b44bca">188路</a><a href="/x_7a6c561d">153路</a><a href="/x_9390328e">154路</a><a href="/x_b75ac673">1077路</a><a href="/x_be53b447">14路</a><a href="/x_682f18cf">119路</a><a href="/x_2a23690d">122路</a><a href="/x_1df76dff">102路</a><a href="/x_e295e51a">151路</a><a href="/x_e8459e67">111路</a><a href="/x_89cd60df">163路</a><a href="/x_8acdc0e4">173路</a><a href="/x_709b4857">1119路</a><a href="/x_2cd601fd">165路</a><a href="/x_1702386f">166路</a><a href="/x_bd6c28fa">172路</a><a href="/x_8a505753">1037路</a><a href="/x_e485d149">155路</a><a href="/x_273f6d5f">178路</a><a href="/x_1866f5a7">180路</a><a href="/x_82af23ab">181路</a><a href="/x_7898317a">182路</a><a href="/x_f48371dd">13路</a><a href="/x_aa874648">198路</a><a href="/x_a24ac753">156B路</a><a href="/x_32d94193">110路</a><a href="/x_74a93310">176路</a><a href="/x_edc4ffde">1031路</a><a href="/x_7371a8e7">1041路</a><a href="/x_978afe53">1022路</a><a href="/x_fd17362d">170路</a><a href="/x_52b7ce0c">195路</a><a href="/x_e0f0aca9">184路</a><a href="/x_f371ab52">185路</a><a href="/x_4e652a14">121路</a><a href="/x_1e6e0362">187路</a><a href="/x_00e31db4">1068路</a><a href="/x_7590d7b4">150路</a><a href="/x_957667e9">193路</a><a href="/x_46fd9870">128路</a><a href="/x_14309d79">194路</a><a href="/x_07ed7df5">161路</a><a href="/x_deb6cfd4">160路</a><a href="/x_0565da64">1086路</a><a href="/x_5674204a">129路</a><a href="/x_d810a85c">133路</a><a href="/x_8432db38">164路</a><a href="/x_5a54ae14">15路</a><a href="/x_589ce89e">1048路</a><a href="/x_7b5d5ffe">127路</a><a href="/x_6f2a7370">126路</a><a href="/x_13cad322">196路</a><a href="/x_6d94bdab">1045路</a><a href="/x_800e9925">191路</a><a href="/x_0494fb39">1052路</a><a href="/x_bd252eca">1032路</a><a href="/x_642b7328">1043路</a><a href="/x_af7e1089">169路</a><a href="/x_65ed8ab0">1012路</a><a href="/x_73791cbc">120路</a><a href="/x_38064e17">1014路</a><a href="/x_16e18930">147路</a><a href="/x_d8202f46">124路</a><a href="/x_a27097d9">1060路</a><a href="/x_7fb7736a">1055路</a><a href="/x_2c45db11">1035路</a><a href="/x_41c8e442">1030路</a><a href="/x_fc63518c">1013路</a><a href="/x_d972733b">1044路</a><a href="/x_6c2241f5">146路</a><a href="/x_e074abd5">1094路</a><a href="/x_5123e995">1003路</a><a href="/x_a07d9a3a">1002路</a><a href="/x_ff5f95a6">1001路</a><a href="/x_ee8ac9e1">1006路</a><a href="/x_06ace4cf">1005路</a><a href="/x_0416b41b">1004路</a><a href="/x_e27620fb">1010路</a><a href="/x_f477448e">1009路</a><a href="/x_3f2e7174">1008路</a><a href="/x_71380ced">1011路</a><a href="/x_a7656622">1015路</a><a href="/x_2084a19b">1074路</a><a href="/x_4f9afb01">1016路</a><a href="/x_7b69e121">171B路</a><a href="/x_d8d16581">171A路</a><a href="/x_72446302">1017路</a><a href="/x_7a7957e4">162路</a><a href="/x_e049f7ee">1020路</a><a href="/x_463cca1a">1019路</a><a href="/x_97e6bcae">1018路</a><a href="/x_9552a285">1025路</a><a href="/x_f1ef6796">1024路</a><a href="/x_e2679952">1021路</a><a href="/x_015ca9c8">1023路</a><a href="/x_3021fac3">1026路</a><a href="/x_f3b18a7f">1027路</a><a href="/x_d934a63d">1028路</a><a href="/x_34228a53">1034路</a><a href="/x_bfb41f16">1036路</a><a href="/x_c4c7d28e">1040路</a><a href="/x_f02e2753">1051路</a><a href="/x_2d3aee26">1054路</a><a href="/x_7c3f42da">1053路</a><a href="/x_6e102862">1058路</a><a href="/x_ed43aa9d">1056路</a><a href="/x_071f1c88">1059路</a><a href="/x_0e755b82">1039路</a><a href="/x_03e81a9f">158路</a><a href="/x_ef058c11">1061路</a><a href="/x_902d58e5">198A路</a><a href="/x_7314fdeb">1063路</a><a href="/x_00a50c88">1131路</a><a href="/x_90587889">179路</a><a href="/x_4e1f3f98">186路</a><a href="/x_64a8585a">1067路</a><a href="/x_2eb30b51">1075路</a><a href="/x_2c1fd985">1066路</a><a href="/x_e0893c6b">1076路</a><a href="/x_aff6e2d8">183路</a><a href="/x_2522b361">1097路</a><a href="/x_a428fb57">1091路</a><a href="/x_3e5f2979">1062路</a><a href="/x_2c9adc5c">1095路</a><a href="/x_e17fdcd3">1080路</a><a href="/x_d043e4c1">1098路</a><a href="/x_24fdc294">1093路</a><a href="/x_f14c98a5">1083路</a><a href="/x_aa3749c7">1070路</a><a href="/x_d5ce3ca8">1085路</a><a href="/x_6fe9d540">1084路</a><a href="/x_d9218dd8">1078路</a><a href="/x_841a7553">1100路</a><a href="/x_8df715dc">1087路</a><a href="/x_3ee314c9">177路</a><a href="/x_b4ad1cc1">1101路</a><a href="/x_cab2fb26">1112路</a><a href="/x_ab7bfe4f">1107路</a><a href="/x_8d5f1130">1071路</a><a href="/x_89c4458d">125路</a><a href="/x_1767112f">1118路</a><a href="/x_da5c7336">138路</a><a href="/x_0fb11b31">1133路</a><a href="/x_bbb9a95e">175路</a><a href="/x_c0d42451">145路</a><a href="/x_811379a3">199路</a><a href="/x_161e0fd7">1117路</a><a href="/x_4b7b834b">1123路</a><a href="/x_a2e82f08">1033路</a><a href="/x_0affc41b">131路</a><a href="/x_3b316c4e">1090路</a><a href="/x_a24fffda">1125路</a><a href="/x_f3b83285">1124路</a><a href="/x_b126542c">1126路</a><a href="/x_e04051f9">156A路</a><a href="/x_f01f45f5">1121路</a><a href="/x_50565920">1120路</a><a href="/x_d0f70b9d">141路</a><a href="/x_eb6b28a4">143路</a><a href="/x_66338752">1130路</a><a href="/x_ee8b2d1e">1122路</a><a href="/x_943e6e75">1157路</a><a href="/x_e8600a13">1047路</a><a href="/x_05ffaae9">1134路</a><a href="/x_1166e3f7">1160路</a><a href="/x_5a8d1d70">1161路</a><a href="/x_8b3812cb">1108路</a><a href="/x_47bbd47b">142路</a><a href="/x_652ba01e">1029路</a><a href="/x_c7d72e14">1046路</a><a href="/x_bead0f17">1038路</a><a href="/x_739f8ab8">1042路</a><a href="/x_da101425">1096路</a><a href="/x_beb36acc">1057路</a><a href="/x_3dc15678">1007路</a><a href="/x_7e5013d2">1065路</a><a href="/x_417595e7">1145路</a><a href="/x_036fe164">1079路</a></div>
	</body>
</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
	}
	doc.Find("#con_site_1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})



	// Output: Header
}

func CrawLineStation() {
	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var lines []models.DbLineModel
	db.Raw("select line_id,line_name from btk_lines").Find(&lines)
	for i:=0; i<len(lines);i++  {
		resp, err := http.Get("http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=One&name="+lines[i].LineName)
		if err != nil {
			// handle error
		}

		// defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err is",err.Error())
		}else{
			fmt.Println("resp is",string(body))
			var lineinfo models.UpdownModel
			err1 := jsonf.Unmarshal(body, &lineinfo)
			if err1!= nil {
				fmt.Println("err1.err is",err1.Error())
			}else{
				db.Exec("update btk_lines set end_earlytime=?,end_latetime=?,end_stop=?,start_earlytime=?,start_latetime=?,start_stop=? where line_id=?",lineinfo.EndEarlytime,lineinfo.EndLatetime,lineinfo.EndStop,lineinfo.StartEarlytime,lineinfo.StartLatetime,lineinfo.StartStop,lines[i].LineId)
				fmt.Println("update btk_lines set end_earlytime=?,end_latetime=?,end_stop=?,start_earlytime=?,start_latetime=?,start_stop=? where line_id=?",lineinfo.EndEarlytime,lineinfo.EndLatetime,lineinfo.EndStop,lineinfo.StartEarlytime,lineinfo.StartLatetime,lineinfo.StartStop,lines[i].LineId)
			}
		}
	}
}

func CrawLineStation1() {
	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var lines []models.DbLineModel
	db.Raw("select line_id,line_name from btk_lines").Find(&lines)
	for i:=0; i<len(lines);i++  {
		resp, err := http.Get("http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=Two&lineid="+lines[i].LineId+"&name="+lines[i].LineName)
		if err != nil {
			// handle error
		}

		// defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err is",err.Error())
		}else{
			fmt.Println("resp is",string(body))
			var lsm models.LineStationModel
			err1 := jsonf.Unmarshal(body, &lsm)
			if err1!= nil {
				fmt.Println("err1.err is","http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=Two&lineid="+lines[i].LineId+"&name="+lines[i].LineName,err1.Error())
			}else{
				//往line_station插入数据
				fmt.Println("lsm.0 is",lsm.LineResults0.Direction)
				fmt.Println("lsm.1 is",lsm.LineResults1.Direction)
				fmt.Println("lsm.LineResults0.Stops.len is",len(lsm.LineResults0.Stops))
				if lsm.LineResults0.Direction == "true"{
					for k:=0;k<len(lsm.LineResults0.Stops) ;k++  {
						fmt.Println("insert into btk_line_station(line_id,line_name,line_updown,st_name,stop_id) values(?,?,?,?,?)",lines[i].LineId,lines[i].LineName,0,lsm.LineResults0.Stops[k].Zdmc,lsm.LineResults0.Stops[k].Id)
						db.Exec("insert into btk_line_station(line_id,line_name,line_updown,st_name,stop_id) values(?,?,?,?,?)",lines[i].LineId,lines[i].LineName,0,lsm.LineResults0.Stops[k].Zdmc,lsm.LineResults0.Stops[k].Id)
					}
				}
				if lsm.LineResults1.Direction == "false"{
					for k:=0;k<len(lsm.LineResults1.Stops) ;k++  {
						db.Exec("insert into btk_line_station(line_id,line_name,line_updown,st_name,stop_id) values(?,?,?,?,?)",lines[i].LineId,lines[i].LineName,1,lsm.LineResults1.Stops[k].Zdmc,lsm.LineResults1.Stops[k].Id)
					}
				}
			}
		}
	}
}

func CrawLineIdUpdownList() {
	fmt.Println("抓取开始")
	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	//http://61.132.47.90:8998/BusService/Require_AllRouteData/?TimeStamp=123
	//先获取所有线路
	resp, err := http.Get("http://www.gembo.cn/app/shbus/all2018.xml")
	if err != nil {
		// handle error
	}

	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	xml := strings.NewReader(string(body))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	var clm models.CrawLineModel

	err1 := jsonf.Unmarshal([]byte(json.String()), &clm)
	if err1 != nil {
		fmt.Println("err1 is", err.Error())
	} else {
		for i := 0; i < len(clm.Lines.Line); i++ {
			resp, err := http.Get("http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=One&name="+clm.Lines.Line[i].Name)
			if err != nil {
				// handle error
			}

			// defer resp.Body.Close()
			bb, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				// handle error
			}
			var updown models.UpdownModel

			err3 := jsonf.Unmarshal(bb, &updown)
			if err3 !=nil {
				fmt.Println("err3.err is",clm.Lines.Line[i].Name,string(bb),err3.Error())
			}else{
				db.Exec("update btk_lines set line_id=? where line_name=?",updown.LineId,clm.Lines.Line[i].Name)
			}
			//if i==0{
			//db.Exec("insert into btk_lines(line_id,line_name) values(?,?)", clm.Lines.Line[i].Value, clm.Lines.Line[i].Name)
			//fmt.Println("insert into btk_lines(line_id,line_name) values(?,?)", clm.Lines.Line[i].Value, clm.Lines.Line[i].Name)
			//}

		}
	}

	fmt.Println("抓取完成!")
}

func CrawLineNameList() {
	fmt.Println("抓取开始")
	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	//http://61.132.47.90:8998/BusService/Require_AllRouteData/?TimeStamp=123
	//先获取所有线路
	resp, err := http.Get("http://www.gembo.cn/app/shbus/all2018.xml")
	if err != nil {
		// handle error
	}

	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	xml := strings.NewReader(string(body))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	var clm models.CrawLineModel

	err1 := jsonf.Unmarshal([]byte(json.String()), &clm)
	if err1 != nil {
		fmt.Println("err1 is", err.Error())
	} else {
		for i := 0; i < len(clm.Lines.Line); i++ {
			//if i==0{
			db.Exec("insert into btk_lines(line_id,line_name) values(?,?)", clm.Lines.Line[i].Value, clm.Lines.Line[i].Name)
			fmt.Println("insert into btk_lines(line_id,line_name) values(?,?)", clm.Lines.Line[i].Value, clm.Lines.Line[i].Name)
			//}

		}
	}

	fmt.Println("抓取完成!")

}
