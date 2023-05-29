package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// 定义结构体, 继承ChainCode接口
type FarmC struct {
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

// 方法实现
func (t* FarmC)Init(stub shim.ChaincodeStubInterface) peer.Response {
	return t.init(stub)
}

func (t* FarmC)init(stub shim.ChaincodeStubInterface) peer.Response {
	// 数据初始化
	infos := []FarmCInfo{
		FarmCInfo{Id:"CF-001", Name:"东部橡胶林区1", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:13, Ordinate:32, Open: "可割胶", FromId:"",Temperature: 15, Humidity: 0.50, Phenology: "抽芽期", Speed: 2.2,Yellow: 0.15},
		FarmCInfo{Id:"CF-002", Name:"东部橡胶林区2", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:26, Ordinate:91, Open: "可割胶", FromId:"",Temperature: 20, Humidity: 0.2, Phenology: "抽芽期", Speed: 2.3,Yellow: 0.50},
		FarmCInfo{Id:"CF-003", Name:"东部橡胶林区3", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:17, Ordinate:82, Open: "可割胶", FromId:"",Temperature: 19, Humidity: 0.3, Phenology: "抽芽期", Speed: 0.3,Yellow: 0.30},
		FarmCInfo{Id:"CF-004", Name:"东部橡胶林区4", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:93, Ordinate:36, Open: "可割胶", FromId:"",Temperature: 22, Humidity: 0.4, Phenology: "抽芽期", Speed: 3.1,Yellow: 0.35},
		FarmCInfo{Id:"CF-005", Name:"东部橡胶林区5", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:17, Ordinate:26, Open: "可割胶", FromId:"",Temperature: 21, Humidity: 0.6, Phenology: "抽芽期", Speed: 0.9,Yellow: 0.45},
		FarmCInfo{Id:"CF-006", Name:"东部橡胶林区6", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:34, Ordinate:63, Open: "", FromId:"",Temperature: 23, Humidity: 0.5, Phenology: "抽芽期", Speed: 2.5,Yellow: 0.65},
		FarmCInfo{Id:"CF-007", Name:"东部橡胶林区7", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:100, Ordinate:60, Open: "", FromId:"",Temperature: 16, Humidity: 0.2, Phenology: "抽芽期", Speed: 3.8,Yellow: 0.70},
		FarmCInfo{Id:"CF-008", Name:"东部橡胶林区8", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:100, Ordinate:60, Open: "", FromId:"",Temperature: 17, Humidity: 0.8, Phenology: "抽芽期", Speed: 4.1,Yellow: 0.25},
		FarmCInfo{Id:"CF-009", Name:"东部橡胶林区9", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:100, Ordinate:60, Open: "", FromId:"",Temperature: 18, Humidity: 0.9, Phenology: "抽芽期", Speed: 5.2,Yellow: 0.31},
		FarmCInfo{Id:"CF-010", Name:"东部橡胶林区10", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:100, Ordinate:60, Open: "", FromId:"",Temperature: 19, Humidity: 1.0, Phenology: "抽芽期", Speed: 6.1,Yellow: 0.07},
		FarmCInfo{Id:"CF-011", Name:"东部橡胶林区11", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:100, Ordinate:60, Open: "", FromId:"",Temperature: 22, Humidity: 0.1, Phenology: "抽芽期", Speed: 7.1,Yellow: 0.9},
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

func(t* FarmC)Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	var result string
	_ = result
	var err error
	if funcName == "setvalue" {
		result, err = setvalue(stub, args)
	}else if funcName == "query" {
		return t.query(stub, args)
	}else if funcName == "trace" {
		return t.trace(stub, args)
	} else if funcName == "judge" {
		return t.judge(stub, args)
	}else if funcName == "isopen" {
		return t.isopen(stub, args)
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
	var temp FarmCInfo

	temp.Id = args[0]
	temp.Name = args[1]
	temp.Date = args[2]
	temp.Quality = args[3]
	temp.Healthy = args[4]
	temp.Soil = args[5]
	temp.Abscissa, _ = strconv.Atoi(args[6])
	temp.Ordinate, _ = strconv.Atoi(args[7])
	temp.Open = args[8]
	temp.FromId = args[9]
	temp.Temperature, _ = strconv.ParseFloat(args[10],64)
	temp.Humidity, _ = strconv.ParseFloat(args[11],64)
	temp.Phenology = args[12]
	temp.Speed, _ = strconv.ParseFloat(args[13],64)
	temp.Method = args[14]
	temp.Yellow, _ = strconv.ParseFloat(args[15],64)
	jsonText, _ := json.Marshal(temp)
	err := stub.PutState(args[0], jsonText)
	if err != nil {
		return "", fmt.Errorf("Failed to set value: %s", args[0])
	}
	return args[1], nil
}


func(t* FarmC)query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	keyID := args[0]
	value, error := stub.GetState(keyID)
	if error != nil {
		return shim.Error("GetState fail...")
	}
	return shim.Success(value)
}
//
func(t* FarmC)judge(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}
	var result string
	var eastItem FarmCInfo
	json.Unmarshal(text.Payload, &eastItem)


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
	name = eastItem.Name

	temperature = eastItem.Temperature
	humidity  = eastItem.Humidity
	speed = eastItem.Speed
	phenology = eastItem.Phenology

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
	eastItem.Method = methods
	jsonText, error := json.Marshal(eastItem)
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
func(t* FarmC)trace(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}
	var result string
	var fromID string
	// 获取需要的信息
	var eastItem FarmCInfo
	json.Unmarshal(text.Payload, &eastItem)
	fromID = eastItem.FromId
	result += fmt.Sprintf("东部橡胶林:%s, ID:%s, 树木质量等级:%s, 健康状况:%s, 土壤湿润程度:%s, 是否可割胶:%s, FromID:%s <--- ", eastItem.Name, eastItem.Id,  eastItem.Quality, eastItem.Healthy, eastItem.Soil, eastItem.Open, eastItem.FromId)

	// 搜索中部林场信息
	myArgs := [][]byte{[]byte("query"), []byte(fromID)}
	response := stub.InvokeChaincode("mid369", myArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error("InvokeChaincode error ......"+ string(response.Payload))
	}
	var midItem FarmBInfo
	json.Unmarshal(response.Payload, &midItem)
	fromID = midItem.FromId
	result += fmt.Sprintf("中部橡胶林:%s, ID:%s,  树木质量等级:%s, 健康状况:%s, 土壤湿润程度:%s, 是否可割胶:%s, FromID:%s <--- ", midItem.Name, midItem.Id, midItem.Quality, midItem.Healthy, midItem.Soil, midItem.Open, midItem.FromId)

	// 搜索西部林场信息
	myArgs = [][]byte{[]byte("query"), []byte(fromID)}
	response = stub.InvokeChaincode("west369", myArgs, "mychannel")
	if response.Status != shim.OK {
		return shim.Error("InvokeChaincode error ......")
	}
	var westItem FarmAInfo
	json.Unmarshal(response.Payload, &westItem)
	result += fmt.Sprintf("西部橡胶林:%s, 树木质量等级:%s, 健康状况:%s, 土壤湿润程度:%s, 是否可割胶:%s  。", westItem.Name, westItem.Quality, westItem.Healthy, westItem.Soil, westItem.Open)

	// 数据格式化为json
	//jsonText, error := json.Marshal(myList)
	//if error != nil {
	//	return shim.Error("json.Marshal(myList) fail...")
	//}
	return shim.Success([]byte(result))
}
func(t* FarmC)isopen(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}

	var eastItem FarmCInfo
	json.Unmarshal(text.Payload, &eastItem)


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
	name = eastItem.Name

	temperature = eastItem.Temperature
	humidity  = eastItem.Humidity
	speed = eastItem.Speed
	healthy = eastItem.Healthy
	yellow = eastItem.Yellow

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
		eastItem.Open = "可割胶"
		jsonText, error := json.Marshal(eastItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	} else{
		eastItem.Open = "不可割胶"
		jsonText, error := json.Marshal(eastItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}

	if humidity >= 0 && humidity <0.25 {
		eastItem.Soil = "干旱"
		jsonText, error := json.Marshal(eastItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}else if humidity >=0.25 && humidity < 0.50{
		eastItem.Soil = "半干旱"
		jsonText, error := json.Marshal(eastItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.50 && humidity < 0.75{
		eastItem.Soil = "半湿润"
		jsonText, error := json.Marshal(eastItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.75 && humidity <= 1{
		eastItem.Soil = "湿润"
		jsonText, error := json.Marshal(eastItem)
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


func main() {
	error := shim.Start(new(FarmC))
	if error != nil {
		println("程序启动失败...")
		return
	}
	fmt.Println("程序启动成功...")
}