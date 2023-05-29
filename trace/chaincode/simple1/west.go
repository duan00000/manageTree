package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"math"
	"strconv"
	"time"
)

// 定义结构体, 继承ChainCode接口
type FarmA struct {
}

// 定义数据结构体
type FarmAInfo struct {
	Id      string
	// 橡胶林名字
	Name    string
	// 生产日期
	Date    string
	// 质量等级
	Quality string
	// 健康状况（健康，良好，差）
	Healthy   string
	// 土壤湿润程度 （湿润、半湿润、半干旱、干旱）
	Soil string
	// 横坐标
	Abscissa int
	//纵坐标
	Ordinate int
	//是否开放
	Open string

	//林区温度
	Temperature float64
	//相对湿度
	Humidity float64
	//第二蓬叶物候
	Phenology string
	//风速
	Speed float64
	//割胶方式
	Method string
	//橡胶树黄叶占比系数
	Yellow float64
	//路径
	Way string


}
// 定义数据结构体
type FarmCInfo struct {
	Id      string
	// 橡胶林名字
	Name    string
	// 生产日期
	Date    string
	// 质量等级
	Quality string
	// 健康状况（健康，良好，差）
	Healthy   string
	// 土壤湿润程度 （湿润、半湿润、半干旱、干旱）
	Soil string
	// 横坐标
	Abscissa int
	//纵坐标
	Ordinate int
	//是否开放
	Open string
	//上一个林场的ID
	FromId   string

	//林区温度
	Temperature float64
	//相对湿度
	Humidity float64
	//第二蓬叶物候
	Phenology string
	//风速
	Speed float64
	//割胶方式
	Method string
	//橡胶树黄叶占比系数
	Yellow float64
}

// 定义数据结构体
type FarmBInfo struct {
	Id      string
	// 橡胶林名字
	Name    string
	// 生产日期
	Date    string
	// 质量等级
	Quality string
	// 健康状况（健康，良好，差）
	Healthy   string
	// 土壤湿润程度 （湿润、半湿润、半干旱、干旱）
	Soil string
	// 横坐标
	Abscissa int
	//纵坐标
	Ordinate int
	//是否开放
	Open string
	//上一个林场的ID
	FromId   string

	//林区温度
	Temperature float64
	//相对湿度
	Humidity float64
	//第二蓬叶物候
	Phenology string
	//风速
	Speed float64
	//割胶方式
	Method string
	//橡胶树黄叶占比系数
	Yellow float64

}


// 方法实现
func (t* FarmA)Init(stub shim.ChaincodeStubInterface) peer.Response {
	return t.init(stub)
}

func (t* FarmA)init(stub shim.ChaincodeStubInterface) peer.Response {
	// 数据初始化
	infos := []FarmAInfo{
		FarmAInfo{Id:"AF-001", Name:"西部橡胶林区1", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:100, Ordinate:60, Open:"可割胶",Temperature: 15, Humidity: 0.5, Phenology: "抽芽期", Speed: 2.3,Yellow: 0.30},
		FarmAInfo{Id:"AF-002", Name:"西部橡胶林区2", Date:"2018-12-12", Quality:"优", Healthy:"健康", Soil:"", Abscissa:70, Ordinate:30, Open:"可割胶", Temperature: 21, Humidity: 0.7, Phenology: "变色期", Speed: 3.5,Yellow: 0.40},
		FarmAInfo{Id:"AF-003", Name:"西部橡胶林区3", Date:"2018-12-13", Quality:"良", Healthy:"健康", Soil:"", Abscissa:80, Ordinate:100, Open: "可割胶", Temperature: 25, Humidity: 0.6, Phenology: "展叶期", Speed: 2.3,Yellow: 0.50},
		FarmAInfo{Id:"AF-004", Name:"西部橡胶林区4", Date:"2018-12-14", Quality:"良", Healthy:"健康", Soil:"", Abscissa:50, Ordinate:30, Open: "可割胶", Temperature: 18, Humidity: 0.7, Phenology: "稳定期", Speed: 4.1,Yellow: 0.20},
		FarmAInfo{Id:"AF-005", Name:"西部橡胶林区5", Date:"2018-12-15", Quality:"差", Healthy:"健康", Soil:"", Abscissa:90, Ordinate:29, Open: "可割胶", Temperature: 25, Humidity: 0.6, Phenology: "稳定期", Speed: 5.5,Yellow: 0.10},
		FarmAInfo{Id:"AF-006", Name:"西部橡胶林区6", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:36, Ordinate:58, Open: "", Temperature: 28, Humidity: 0.1, Phenology: "展叶期", Speed: 5.6,Yellow: 0.10},
		FarmAInfo{Id:"AF-007", Name:"西部橡胶林区7", Date:"2018-12-12", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:10, Ordinate:60, Open: "", Temperature: 26, Humidity: 0.3, Phenology: "变色期", Speed: 6.1,Yellow: 0.20},
		FarmAInfo{Id:"AF-008", Name:"西部橡胶林区8", Date:"2018-12-13", Quality:"优", Healthy:"健康", Soil:"", Abscissa:14, Ordinate:40, Open: "", Temperature: 21, Humidity: 0.4, Phenology: "变色期", Speed: 7.0,Yellow: 0.15},
		FarmAInfo{Id:"AF-009", Name:"西部橡胶林区9", Date:"2018-12-14", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:53, Ordinate:37, Open: "", Temperature: 22, Humidity: 0.94, Phenology: "抽芽期", Speed: 5.3,Yellow: 0.20},
		FarmAInfo{Id:"AF-010", Name:"西部橡胶林区10", Date:"2018-12-15", Quality:"良", Healthy:"健康", Soil:"", Abscissa:14, Ordinate:569, Open: "", Temperature: 23, Humidity: 0.45, Phenology: "抽芽期", Speed: 0.3,Yellow: 0.60},
		FarmAInfo{Id:"AF-011", Name:"西部橡胶林区11", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:10, Ordinate:963, Open: "", Temperature: 16, Humidity: 0.86, Phenology: "展叶期", Speed: 0.6,Yellow: 0.30},
		FarmAInfo{Id:"AF-012", Name:"西部橡胶林区12", Date:"2018-12-12", Quality:"优", Healthy:"健康", Soil:"", Abscissa:159, Ordinate:60, Open: "", Temperature: 17, Humidity: 0.63, Phenology: "展叶期", Speed: 0.9,Yellow: 0.30},
		FarmAInfo{Id:"AF-013", Name:"西部橡胶林区13", Date:"2018-12-13", Quality:"差", Healthy:"健康", Soil:"", Abscissa:35, Ordinate:95, Open: "", Temperature: 18, Humidity: 0.79, Phenology: "稳定期", Speed: 1.3,Yellow: 0.30},
		FarmAInfo{Id:"AF-014", Name:"西部橡胶林区14", Date:"2018-12-14", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:46, Ordinate:160, Open: "", Temperature: 19, Humidity: 0.14, Phenology: "稳定期", Speed: 1.5,Yellow: 0.30},
		FarmAInfo{Id:"AF-015", Name:"西部橡胶林区15", Date:"2018-12-15", Quality:"差", Healthy:"有死皮", Soil:"", Abscissa:93, Ordinate:73, Open: "", Temperature: 24, Humidity: 0.52, Phenology: "变色期", Speed: 1.9,Yellow: 0.30},
		FarmAInfo{Id:"AF-016", Name:"西部橡胶林区16", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:17, Ordinate:59, Open: "", Temperature: 21, Humidity: 0.63, Phenology: "变色期", Speed: 2.4,Yellow: 0.30},
		FarmAInfo{Id:"AF-017", Name:"西部橡胶林区17", Date:"2018-12-12", Quality:"优", Healthy:"健康", Soil:"", Abscissa:69, Ordinate:46, Open: "", Temperature: 22, Humidity: 0.96, Phenology: "定期", Speed: 1.1,Yellow: 0.30},
		FarmAInfo{Id:"AF-019", Name:"西部橡胶林区19", Date:"2018-12-14", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:93, Ordinate:75, Open: "", Temperature: 15, Humidity: 0.42, Phenology: "定期", Speed: 6.9,Yellow: 0.30},
		FarmAInfo{Id:"AF-020", Name:"西部橡胶林区20", Date:"2018-12-15", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:24, Ordinate:63, Open: "", Temperature: 25, Humidity: 0.95, Phenology: "抽芽期", Speed: 4.2,Yellow: 0.30},
	}
	// 遍历, 写入账本中
	i := 0
	for i < len(infos) {
		jsontext, error := json.Marshal(infos[i])
		if error != nil {
			return shim.Error("init error, json marshal fail...")
		}
		// 数据写入账本中
		stub.PutState(infos[i].Id, jsontext)
		i++
	}
	return shim.Success([]byte("init ledger OK!!!"))
}

func(t* FarmA)Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()

	var result string
	_ = result
	var err error
	if funcName == "setvalue" {
		result, err = setvalue(stub, args)
	}else if funcName == "query" {
		return t.query(stub, args)
	}else if funcName == "gethistory" {
		return t.gethistory(stub, args)
	}else if funcName == "plan" {
		result, err = t.plan(stub, args)
	}else if funcName == "judge" {
		return t.judge(stub, args)
	}else if funcName == "isopen" {
		return t.isopen(stub, args)
	}else if funcName == "setbvalue" {
		result, err = t.setbvalue(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("invoke OK"))
}


func setvalue(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 16 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}
	var temp FarmAInfo

	temp.Id = args[0]
	temp.Name = args[1]
	temp.Date = args[2]
	temp.Quality = args[3]
	temp.Healthy = args[4]
	temp.Soil = args[5]
	temp.Abscissa, _ = strconv.Atoi(args[6])
	temp.Ordinate, _ = strconv.Atoi(args[7])
	temp.Open = args[8]

	temp.Temperature, _ = strconv.ParseFloat(args[9],64)
	temp.Humidity, _ = strconv.ParseFloat(args[10],64)
	temp.Phenology = args[11]
	temp.Speed, _ = strconv.ParseFloat(args[12],64)
	temp.Method = args[13]
	temp.Yellow, _ = strconv.ParseFloat(args[14],64)
	temp.Way = args[15]
	jsonText, _ := json.Marshal(temp)
	err := stub.PutState(args[0], jsonText)
	if err != nil {
		return "", fmt.Errorf("Failed to set value: %s", args[0])
	}
	return args[1], nil

}

func(t* FarmA)query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	keyID := args[0]
	value, error := stub.GetState(keyID)
	if error != nil {
		return shim.Error("GetState fail...")
	}
	return shim.Success(value)
}

func(t* FarmA)isopen(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}

	var westItem FarmAInfo
	json.Unmarshal(text.Payload, &westItem)


	var a = 0.0
	var b = 0.0
	var c = 0.0
	var d = 0.0
	var e = 0.0
	var final = 0.0
	//割胶方式添加专用


	_ = a
	_ = b
	_ = c
	_ = d
	_ = e

	var id string
	var name string

	var temperature float64
	var humidity float64
	var speed float64
	var healthy string
	var yellow float64


	_ = name

	id = args[0]
	name = westItem.Name

	temperature = westItem.Temperature
	humidity  = westItem.Humidity
	speed = westItem.Speed
	healthy = westItem.Healthy
	yellow = westItem.Yellow

	if temperature <= 28 && temperature >= 15{
		a = 1
	}
    if humidity <= 0.80 && humidity >=0.20{
		b = 1
	}

	if speed <=7.9  {
		c = 1
	}
	if healthy == "健康"{
		d =  1
	}
	if yellow <= 0.5{
		e =  1
	}

	final = a+b+c+d + e

     if final == 5{
		 westItem.Open = "可割胶"
		 jsonText, error := json.Marshal(westItem)
		 if error != nil {
			 return shim.Error("json.Marshal(myList) fail...")
		 }
		 error0 := stub.PutState(id, []byte(jsonText))
		 if error0 != nil {
			 return shim.Error("PutState fail...")
		 }

	 } else{
		westItem.Open = "不可割胶"
		jsonText, error := json.Marshal(westItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}

	if humidity >= 0 && humidity <0.25 {
		westItem.Soil = "干旱"
		jsonText, error := json.Marshal(westItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}else if humidity >=0.25 && humidity < 0.50{
		westItem.Soil = "半干旱"
		jsonText, error := json.Marshal(westItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.50 && humidity < 0.75{
		westItem.Soil = "半湿润"
		jsonText, error := json.Marshal(westItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.75 && humidity <= 1{
		westItem.Soil = "湿润"
		jsonText, error := json.Marshal(westItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}

	return shim.Success([]byte("SetValue Sucess"))


}


func(t* FarmA)judge(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}
	var result string
	var westItem FarmAInfo
	json.Unmarshal(text.Payload, &westItem)


	var methods = ""
	var a = 0.0
	var b = 0.0
	var c = 0.0
	var d = 0.0
	var final = 0.0
	//割胶方式添加专用


	_ = a
	_ = b
	_ = c
	_ = d

	var id string
	var name string

	var temperature float64
	var humidity float64
	var speed float64
	var phenology string


	_ = name

	id = args[0]
	name = westItem.Name

	temperature = westItem.Temperature
	humidity  = westItem.Humidity
	speed = westItem.Speed
	phenology = westItem.Phenology

	if temperature <= 21 && temperature >= 15{
		a = 0.98-((temperature - 14) * 0.07)
	}
	if temperature >= 22 && temperature <= 28{
		a = (temperature - 14) * 0.07
	}
	b = humidity * 100 * 0.01

	if speed <=7.9 && speed>=0 {
		c = 0.0125 * (speed+0.1) *10
	}
	if phenology == "抽芽期"{
		d = 0.25
	}
	if phenology == "展页期"{
		d = 0.50
	}
	if phenology == "变色期"{
		d = 0.75
	}
	if phenology == "稳定期"{
		d = 1
	}
	final = a + b + c + d
	if final >= 0 && final <= 2.2275{
		methods = "深割"
	}
	if final >2.2275 && final <= 3.98{
		methods = "浅割"
	}


	//将割胶方式添加进jason中
	westItem.Method = methods
	jsonText, error := json.Marshal(westItem)
	if error != nil {
		return shim.Error("json.Marshal(myList) fail...")
	}
	error0 := stub.PutState(id, []byte(jsonText))
	if error0 != nil {
		return shim.Error("PutState fail...")
	}


	result += fmt.Sprintf("%s的最优割胶方式为 %s", name,methods)
	return shim.Success([]byte(result))
}


// 根据keyID查询历史记录
func(t* FarmA)gethistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	keyiter, error := stub.GetHistoryForKey(args[0])
	if error != nil {
		return shim.Error("GetHistoryForKey fail...")
	}
	defer keyiter.Close()
	// 通过迭代器对象遍历结果
	var myList []string
	for keyiter.HasNext() {
		// 获取当前值
		result, error := keyiter.Next()
		if error != nil {
			return shim.Error("keyiter.Next() fail...")
		}
		// 获取需要的信息
		txID := result.TxId
		txValue := result.Value
		txTime := result.Timestamp
		txStatus := result.IsDelete
		tm := time.Unix(txTime.Seconds, 0)
		datastr := tm.Format("2006-01-02 15:04:05")
		all := fmt.Sprintf("%s, %s, %s, %t", txID, txValue, datastr, txStatus)
		myList = append(myList, all)
	}
	// 数据格式化为json
	jsonText, error := json.Marshal(myList)
	if error != nil {
		return shim.Error("json.Marshal(myList) fail...")
	}
	return shim.Success(jsonText)
}


func toChaincodeArgs2(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}


func(t* FarmA)setbvalue(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	text := t.query(stub, args)

	if text.Status != shim.OK {
		errStr := fmt.Sprintf("query error ......")
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}
	id := args[0]
	//改变B区域fromID
	myArgsb := [][]byte{[]byte("query"), []byte("BF-001")}
	responseb := stub.InvokeChaincode("mid303", myArgsb, "mychannel")

	if responseb.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(responseb.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}
	var midItem FarmBInfo

	json.Unmarshal(responseb.Payload, &midItem)


	str0 := fmt.Sprintf("%f",midItem.Temperature)
	str1 := fmt.Sprintf("%f",midItem.Humidity)
	str2 := fmt.Sprintf("%f",midItem.Speed)
	str3 := fmt.Sprintf("%f",midItem.Yellow)

	ss0 := midItem.Name
	ss1 := midItem.Date
	ss2 := midItem.Quality
	ss3 := midItem.Healthy
	ss4 := midItem.Soil
	ss5 := midItem.Abscissa
	sss5 := strconv.Itoa(ss5)
	ss6 := midItem.Ordinate
	sss6 := strconv.Itoa(ss6)
	ss7 := midItem.Phenology
	ss8 := midItem.Method
	sss := midItem.Open

	invokeArgs := toChaincodeArgs2("setvalue", "BF-001", ss0,ss1,ss2,ss3,ss4,sss5,sss6,sss,id,str0,str1,ss7,str2,ss8,str3)

	response := stub.InvokeChaincode("mid303", invokeArgs, "mychannel")

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	fmt.Printf("Invoke chaincode successful. Got response %s", string(response.Payload))
	return string(response.Payload), nil

}

func(t* FarmA)plan(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		errStr := fmt.Sprintf("query error ......")
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	var id = args[0]
	var result string
	var eastcanzhao int = 10000
	var jlmidID string
	var jlmidheng int
	var jlmidzong int
	var eastheng int = 0
	var eastzong int = 0
	var eastsub0 int
	var eastsub1 int
	var eastjieguo0 int
	var eastjieguo1 int
	var tempeast string
	var eastadd int
	var jieguo0east int

	var fromID string
	var fromID1 string
	var nextfromID string

	var midheng int = 0
	var midzong int = 0

	var sub0 int
	var sub1 int

	var westadd int
    var jieguo0 int
	var westjieguo0 int
	var westjieguo1 int

	var canzhao int = 100000
	var jilu int = 10
	var westsub0 int = 10
	var westsub1 int = 10
	var finaladd int = 0
	var note int = 0
	var note1 int = 0

	var straset string
	var strbset string
	var strcset string
	var tempmid string
	var tempwest string
	//用于权重的变量
	//叶黄素占比
	var easthealthy string

	var reeasth float64
	var midY float64
	var eastY float64
	//温度度
	var eastwet string

	var reeastw float64
	var midT float64
	var eastT float64
	//用于改变fromid的数据
	var record1 string
	var record2 string
	//质量
	var eastquality string

	var eastH float64
	var midH float64
	//速度
	var midS float64
	var eastS float64
	//加了权重之后的值
	var quanresult0 float64 = 0
	var quanresult1 float64 = 0
	//最终算出的权重参数显示
	var quancun0 float64 = 0
	var quancun1 float64 = 0
	var quanfinal float64 = 0
	var quanzhuan string
	//防止报错
	_ = record1
	_ = record2

	_ = jlmidheng
	_ = jlmidzong
	_ = fromID1
	_ = jieguo0
	_ = quanzhuan
	_ = quanfinal
	_ = reeasth
	_ = midY
	_ = eastY
	_ = eastT
	_ = eastH
	_ = eastS
	_ = midH
	_ = midS
	_ = reeastw
	_ = midT
	_ = quanresult0
	_ = quanresult1
	_ = eastquality

	_ = eastwet

	_ = easthealthy

	_ = tempwest
	_ = tempmid
	_ = strcset
	_ = straset
	_ = strbset
	_ = jilu
	_ = nextfromID
	_ = westadd
	_ = sub0
	_ = sub1

	//三个农场名字初始化
	strbset = "666"

	// 获取需要的信息




	s1 := "BF-00"
	s0 := "CF-00"


	//西区域的横纵坐标
	var westItem FarmAInfo
	json.Unmarshal(text.Payload, &westItem)
	//判定是否开放
	if westItem.Open == "可割胶"{
		westheng := westItem.Abscissa
		westzong := westItem.Ordinate
		straset = westItem.Name
		for i := 1; i <= 5; i++ {

			s2 := s1 + strconv.Itoa(i)
			fromID = s2

			myArgs := [][]byte{[]byte("query"), []byte(fromID)}
			response := stub.InvokeChaincode("mid369", myArgs, "mychannel")

			if response.Status != shim.OK {
				errStr := fmt.Sprintf(" 111 Failed to invoke chaincode. Got error: %s", string(response.Payload))
				fmt.Printf(errStr)
				return "", fmt.Errorf(errStr)
			}

			var midItem FarmBInfo

			json.Unmarshal(response.Payload, &midItem)
			if midItem.Open == "可割胶"{
				note = 1
				midheng = midItem.Abscissa
				midzong = midItem.Ordinate
				tempmid = midItem.Name
				sub0 = westheng - midheng
				sub1 = westzong - midzong
				westjieguo0 = int(math.Pow(float64(westsub0), 2))
				westjieguo1 = int(math.Pow(float64(westsub1), 2))
				westadd = westjieguo0 + westjieguo1
				//用于权重的变量




				//权重到此结束
				//距离的结果
				jieguo0 = int(math.Sqrt(float64(westadd)))
				//叶黄素占比状况   (原来为remidh)

				if midItem.Yellow <=0.49 && midItem.Yellow >=0.40{
					midY = float64(jieguo0) * (midItem.Yellow-0.49 - 0.01) * 10 * (-1) * 0.03
				}
				if midItem.Yellow <=0.39 && midItem.Yellow >=0.30 {
					midY = float64(jieguo0) * (0.03 + (midItem.Yellow-0.39 - 0.01) * 10 * (-1) * 0.03)
				}
				if midItem.Yellow <=0.29 && midItem.Yellow >=0.20 {
					midY = float64(jieguo0) * (0.06 + (midItem.Yellow-0.29 - 0.01) * 10 * (-1) * 0.06)
				}
				if midItem.Yellow <=0.19 && midItem.Yellow >=0.10 {
					midY = float64(jieguo0) * (0.12 + (midItem.Yellow-0.19 - 0.01) * 10 * (-1) * 0.03)
				}
				if midItem.Yellow <=0.09 && midItem.Yellow >=0.00 {
					midY = float64(jieguo0) * (0.15 + (midItem.Yellow-0.09 - 0.01) * 10 * (-1) * 0.03)
				}
				//温度状况    （原来是remidw）

				if midItem.Temperature <= 22 && midItem.Temperature >= 20 {
					midT = float64(jieguo0)  * (0.001 + (22 - midItem.Temperature +1) * 0.013)
				}
				if midItem.Temperature <= 25 && midItem.Temperature >= 23 {
					midT = float64(jieguo0)  * (0.04 + 0.002 + ( midItem.Temperature - 25 - 1) * (-1) * 0.006)
				}
				if midItem.Temperature <= 22 && midItem.Temperature >= 20 {
					midT = float64(jieguo0)  * (0.06 + ( midItem.Temperature - 19 -1) * (-1) * 0.012)
				}
				if midItem.Temperature <= 22 && midItem.Temperature >= 20 {
					midT = float64(jieguo0)  * (0.12 + 0.002 + (28 - midItem.Temperature +1) * 0.006)
				}


				//湿润度 原来为（remidq）

				if midItem.Humidity <= 80  && midItem.Humidity >= 66{
					midH = float64(jieguo0)  * (0.001 + (midItem.Humidity - 80 - 1) * (-1) * 0.0026)
				}
				if midItem.Humidity <= 65  && midItem.Humidity >= 51{
					midH = float64(jieguo0)  * (0.04 + 0.001 + ((midItem.Humidity - 65 - 1) * (-1) * 0.0026))
				}
				if midItem.Humidity <= 50  && midItem.Humidity >= 36{
					midH = float64(jieguo0)  * (0.08 + 0.001 + ((midItem.Humidity - 50 - 1) * (-1) * 0.0026))
				}
				if midItem.Humidity <= 35  && midItem.Humidity >= 20{
					midH = float64(jieguo0)  * (0.12 + ((midItem.Humidity - 35 - 1) * (-1) * 0.0025))
				}

				//风速 （新添加的）

				if midItem.Speed <= 7.9  && midItem.Speed>= 5.5{
					midS = float64(jieguo0)  * ((midItem.Speed - 7.9 - 0.1) * 10 * (-1) * 0.0016)
				}
				if midItem.Speed <= 5.4  && midItem.Speed >= 3{
					midS = float64(jieguo0)  * (0.04 + ((midItem.Speed - 5.4 - 0.1)  * 10 * (-1) * 0.0016))
				}
				if midItem.Speed <= 2.9  && midItem.Speed >= 0{
					midS = float64(jieguo0)  * (0.08 + 0.0001 + ((midItem.Speed - 2.9 - 0.1)  * 10 * (-1) * 0.0133))
				}

				//权重判断结束-----------------------
				quanresult0 = float64(jieguo0) * 0.72  - midY - midH - midS - midT
				if float64(canzhao) > quanresult0 {
					canzhao = jieguo0
					strbset = tempmid
					quancun0 = quanresult0
					jlmidID = midItem.Id


				}
			}





		}

		//改变B区域fromID
		myArgsb := [][]byte{[]byte("query"), []byte(jlmidID)}
		responseb := stub.InvokeChaincode("mid369", myArgsb, "mychannel")

		if responseb.Status != shim.OK {
			errStr := fmt.Sprintf(" 222 Failed to invoke chaincode. Got error: %s", string(responseb.Payload))
			fmt.Printf(errStr)
			return "", fmt.Errorf(errStr)
		}
		var midItem FarmBInfo

		json.Unmarshal(responseb.Payload, &midItem)


		str0 := fmt.Sprintf("%f",midItem.Temperature)
		str1 := fmt.Sprintf("%f",midItem.Humidity)
		str2 := fmt.Sprintf("%f",midItem.Speed)
		str3 := fmt.Sprintf("%f",midItem.Yellow)

		invokeArgs := toChaincodeArgs2("setvalue", jlmidID, midItem.Name,midItem.Date,midItem.Quality,midItem.Healthy,midItem.Soil,string(midItem.Abscissa),string(midItem.Ordinate),midItem.Open,id,str0,str1,midItem.Phenology,str2,midItem.Method,str3)

		responsemid := stub.InvokeChaincode("mid369", invokeArgs, "mychannel")

		if responsemid.Status != shim.OK {

			errStr := fmt.Sprintf(" 333 Failed to invoke chaincode. Got error: %s", string(responsemid.Payload))
			fmt.Printf(errStr)
			return "", fmt.Errorf(errStr)
		}



		//橡胶林区域B是否存在可割胶节点
		if note == 0{
			result += fmt.Sprintf("橡胶林区B未达到割胶条件---")
		}
		if note == 1{
			for j := 1; j <= 5; j++ {

				//中部区域橡胶林的数据存储
				myArgs := [][]byte{[]byte("query"), []byte(jlmidID)}
				response := stub.InvokeChaincode("mid369", myArgs, "mychannel")

				if response.Status != shim.OK {
					errStr := fmt.Sprintf(" 444 Failed to invoke chaincode. Got error: %s", string(response.Payload))
					fmt.Printf(errStr)
					return "", fmt.Errorf(errStr)
				}

				var midItem FarmBInfo
				json.Unmarshal(response.Payload, &midItem)
				//东部区域橡胶林的数据存储
				s6 := s0 + strconv.Itoa(j)
				myArgs0 := [][]byte{[]byte("query"), []byte(s6)}
				response = stub.InvokeChaincode("east369", myArgs0, "mychannel")

				if response.Status != shim.OK {
					errStr := fmt.Sprintf(" 555 Failed to invoke chaincode. Got error: %s", string(response.Payload))
					fmt.Printf(errStr)
					return "", fmt.Errorf(errStr)
				}

				var eastItem FarmCInfo

				json.Unmarshal(response.Payload, &eastItem)
				//判断东部区域是否可以割胶
				if eastItem.Open == "可割胶"{
					note1 = 1
					eastheng = eastItem.Abscissa
					eastzong = eastItem.Ordinate
					tempeast = eastItem.Name
					eastsub0 = eastheng - midItem.Abscissa
					eastsub1 = eastzong - midItem.Ordinate
					eastjieguo0 = int(math.Pow(float64(eastsub0), 2))
					eastjieguo1 = int(math.Pow(float64(eastsub1), 2))
					eastadd = eastjieguo0 + eastjieguo1


					//权重到此结束
					//距离的结果
					jieguo0east = int(math.Sqrt(float64(eastadd)))
					//叶黄素占比状况   (原来为reeasth)

					if eastItem.Yellow <=0.49 && eastItem.Yellow >=0.40{
						eastY = float64(jieguo0east) * (eastItem.Yellow-0.49 - 0.01) * 10 * (-1) * 0.03
					}
					if eastItem.Yellow <=0.39 && eastItem.Yellow >=0.30 {
						eastY = float64(jieguo0east) * (0.03 + (eastItem.Yellow-0.39 - 0.01) * 10 * (-1) * 0.03)
					}
					if eastItem.Yellow <=0.29 && eastItem.Yellow >=0.20 {
						eastY = float64(jieguo0east) * (0.06 + (eastItem.Yellow-0.29 - 0.01) * 10 * (-1) * 0.06)
					}
					if eastItem.Yellow <=0.19 && eastItem.Yellow >=0.10 {
						eastY = float64(jieguo0east) * (0.12 + (eastItem.Yellow-0.19 - 0.01) * 10 * (-1) * 0.03)
					}
					if eastItem.Yellow <=0.09 && eastItem.Yellow >=0.00 {
						eastY = float64(jieguo0east) * (0.15 + (eastItem.Yellow-0.09 - 0.01) * 10 * (-1) * 0.03)
					}
					//温度状况    （原来是reeastw）

					if eastItem.Temperature <= 22 && eastItem.Temperature >= 20 {
						eastT = float64(jieguo0east)  * (0.001 + (22 - eastItem.Temperature +1) * 0.013)
					}
					if eastItem.Temperature <= 25 && eastItem.Temperature >= 23 {
						eastT = float64(jieguo0east)  * (0.04 + 0.002 + ( eastItem.Temperature - 25 - 1) * (-1) * 0.006)
					}
					if eastItem.Temperature <= 22 && eastItem.Temperature >= 20 {
						eastT = float64(jieguo0east)  * (0.06 + ( eastItem.Temperature - 19 -1) * (-1) * 0.012)
					}
					if eastItem.Temperature <= 22 && eastItem.Temperature >= 20 {
						eastT = float64(jieguo0east)  * (0.12 + 0.002 + (28 - eastItem.Temperature +1) * 0.006)
					}


					//湿润度 原来为（reeastq）

					if eastItem.Humidity <= 80  && eastItem.Humidity >= 66{
						eastH = float64(jieguo0east)  * (0.001 + (eastItem.Humidity - 80 - 1) * (-1) * 0.0026)
					}
					if eastItem.Humidity <= 65  && eastItem.Humidity >= 51{
						eastH = float64(jieguo0east)  * (0.04 + 0.001 +((eastItem.Humidity - 65 - 1) * (-1) * 0.0026))
					}
					if eastItem.Humidity <= 50  && eastItem.Humidity >= 36{
						eastH = float64(jieguo0east)  * (0.08 + + 0.001 + ((eastItem.Humidity - 50 - 1) * (-1) * 0.0026))
					}
					if eastItem.Humidity <= 35  && eastItem.Humidity >= 20{
						eastH = float64(jieguo0east)  * (0.12 + ((eastItem.Humidity - 35 - 1) * (-1) * 0.0025))
					}

					//风速 （新添加的）

					if eastItem.Speed <= 7.9  && eastItem.Speed>= 5.5{
						eastS = float64(jieguo0east)  * ((eastItem.Speed - 7.9 - 0.1) * 10 * (-1) * 0.0016)
					}
					if eastItem.Speed <= 5.4  && eastItem.Speed >= 3{
						eastS = float64(jieguo0east)  * (0.04 + ((eastItem.Speed - 5.4 - 0.1)  * 10 * (-1) * 0.0016))
					}
					if eastItem.Speed <= 2.9  && eastItem.Speed >= 0{
						eastS = float64(jieguo0east)  * (0.08 + 0.0001 + ((eastItem.Speed - 2.9 - 0.1)  * 10 * (-1) * 0.0133))
					}

					//权重判断结束-----------------------
					quanresult1 = float64(jieguo0east) * 0.72 - eastH - eastS - eastT - eastY
					if float64(eastcanzhao) > quanresult1 {
						eastcanzhao = jieguo0
						strcset = tempeast
						quancun1 = quanresult1
						record1 = eastItem.Id
					}
				}else{
					result += fmt.Sprintf("%s---不可割胶---",id)
				}


			}


			//改变C区域fromID
			myArgsc := [][]byte{[]byte("query"), []byte(record1)}
			responsec := stub.InvokeChaincode("east369", myArgsc, "mychannel")

			if responsec.Status != shim.OK {
				errStr := fmt.Sprintf(" aaa Failed to invoke chaincode. Got error: %s", string(responsec.Payload))
				fmt.Printf(errStr)
				return "", fmt.Errorf(errStr)
			}
			var eastItem FarmCInfo

			json.Unmarshal(responsec.Payload, &eastItem)


			str0 := fmt.Sprintf("%f",eastItem.Temperature)
			str1 := fmt.Sprintf("%f",eastItem.Humidity)
			str2 := fmt.Sprintf("%f",eastItem.Speed)
			str3 := fmt.Sprintf("%f",eastItem.Yellow)

			invokeArgs3 := toChaincodeArgs2("setvalue", record1, eastItem.Name,eastItem.Date,eastItem.Quality,eastItem.Healthy,eastItem.Soil,string(eastItem.Abscissa),string(eastItem.Ordinate),eastItem.Open,jlmidID,str0,str1,eastItem.Phenology,str2,eastItem.Method,str3)

			responseeast := stub.InvokeChaincode("east369", invokeArgs3, "mychannel")
			if responseeast.Status != shim.OK {

				errStr := fmt.Sprintf(" ccc Failed to invoke chaincode. Got error: %s", string(responseeast.Payload))
				fmt.Printf(errStr)
				return "", fmt.Errorf(errStr)
			}

		}




		if note == 1 && note1 == 1{
			finaladd = canzhao + eastcanzhao

			zhuanhuan := strconv.Itoa(finaladd)
			quanfinal = quancun0 + quancun1
			quanzhuan = strconv.FormatFloat(quanfinal, 'E', -1, 64)//float64
			_ = zhuanhuan

			result += fmt.Sprintf("起点  --->  %s  --->  %s  --->  %s   最短距离：%s ,  算出的最小权重: %s。", straset, strbset, strcset, zhuanhuan ,quanzhuan)
		}



	}else{
		result += fmt.Sprintf("---橡胶林区A未达到割胶条件---")
	}
	//上链
	var westItem0 FarmAInfo
	json.Unmarshal(text.Payload, &westItem0)
	//将割胶方式添加进jason中
	westItem0.Way = result
	jsonText, error := json.Marshal(westItem0)
	if error != nil {
		return "", error
	}
	error0 := stub.PutState(id, []byte(jsonText))
	if error0 != nil {
		return "",error0
	}





	return result, nil


}

func main() {
	error := shim.Start(new(FarmA))
	if error != nil {
		println("程序启动失败...")
		return
	}
	fmt.Println("程序启动成功...")
}