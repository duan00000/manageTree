package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

// 定义结构体, 继承ChainCode接口
type FarmB struct {
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
func (t* FarmB)Init(stub shim.ChaincodeStubInterface) peer.Response {
	return t.init(stub)
}

func (t* FarmB)init(stub shim.ChaincodeStubInterface) peer.Response {
	// 数据初始化
	infos := []FarmBInfo{
		FarmBInfo{Id:"BF-001", Name:"中部橡胶林区1", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:69, Ordinate:52, Open: "可割胶", FromId:"",Temperature: 15, Humidity: 0.5, Phenology: "抽芽期", Speed: 2.3,Yellow: 0.15},
		FarmBInfo{Id:"BF-002", Name:"中部橡胶林区2", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:31, Ordinate:63, Open: "可割胶", FromId:"",Temperature: 20, Humidity: 0.6, Phenology: "抽芽期", Speed: 6.1,Yellow: 0.30},
		FarmBInfo{Id:"BF-003", Name:"中部橡胶林区3", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:56, Ordinate:32, Open: "可割胶", FromId:"",Temperature: 25, Humidity: 0.8, Phenology: "抽芽期", Speed: 5.1,Yellow: 0.20},
		FarmBInfo{Id:"BF-004", Name:"中部橡胶林区4", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:43, Ordinate:95, Open: "可割胶", FromId:"",Temperature: 28, Humidity: 0.2, Phenology: "抽芽期", Speed: 6.3,Yellow: 0.49},
		FarmBInfo{Id:"BF-005", Name:"中部橡胶林区5", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:26, Ordinate:75, Open: "可割胶", FromId:"",Temperature: 26, Humidity: 0.3, Phenology: "展叶期", Speed: 1.0,Yellow: 0.50},
		FarmBInfo{Id:"BF-006", Name:"中部橡胶林区6", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:93, Ordinate:16, Open: "", FromId:"",Temperature: 16, Humidity: 0.6, Phenology: "展叶期", Speed: 0.0,Yellow: 0.49},
		FarmBInfo{Id:"BF-007", Name:"中部橡胶林区7", Date:"2018-12-11", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:159, Ordinate:60, Open: "", FromId:"",Temperature: 17, Humidity: 0.2, Phenology: "展叶期", Speed: 1.2,Yellow: 0.30},
		FarmBInfo{Id:"BF-008", Name:"中部橡胶林区8", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:193, Ordinate:36, Open: "", FromId:"",Temperature: 18, Humidity: 0.3, Phenology: "展叶期", Speed: 0.8,Yellow: 0.60},
		FarmBInfo{Id:"BF-009", Name:"中部橡胶林区9", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:10, Ordinate:56, Open: "", FromId:"",Temperature: 19, Humidity: 0.6, Phenology: "变色期", Speed: 0.4,Yellow: 0.30},
		FarmBInfo{Id:"BF-0010", Name:"中部橡胶林区10", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:100, Ordinate:10, Open: "", FromId:"",Temperature: 20, Humidity: 0.9, Phenology: "变色期", Speed: 3.2,Yellow: 0.20},
		FarmBInfo{Id:"BF-011", Name:"中部橡胶林区11", Date:"2018-12-11", Quality:"差", Healthy:"有死皮", Soil:"", Abscissa:200, Ordinate:90, Open: "", FromId:"",Temperature: 21, Humidity: 0.2, Phenology: "变色期", Speed: 3.6,Yellow: 0.15},
		FarmBInfo{Id:"BF-012", Name:"中部橡胶林区12", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:30, Ordinate:50, Open: "", FromId:"AF-012",Temperature: 22, Humidity: 0.4, Phenology: "变色期", Speed: 3.8,Yellow: 0.15},
		FarmBInfo{Id:"BF-013", Name:"中部橡胶林区13", Date:"2018-12-11", Quality:"良", Healthy:"有死皮", Soil:"", Abscissa:20, Ordinate:66, Open: "", FromId:"AF-013",Temperature: 23, Humidity: 0.6, Phenology: "稳定期", Speed: 4.2,Yellow: 0.15},
		FarmBInfo{Id:"BF-014", Name:"中部橡胶林区14", Date:"2018-12-11", Quality:"差", Healthy:"健康", Soil:"", Abscissa:43, Ordinate:89, Open: "", FromId:"AF-014",Temperature: 24, Humidity: 0.7, Phenology: "稳定期", Speed: 4.6,Yellow: 0.15},
		FarmBInfo{Id:"BF-015", Name:"中部橡胶林区15", Date:"2018-12-11", Quality:"差", Healthy:"有死皮", Soil:"", Abscissa:36, Ordinate:12, Open: "", FromId:"AF-015",Temperature: 25, Humidity: 0.2, Phenology: "稳定期", Speed: 6.3,Yellow: 0.15},
		FarmBInfo{Id:"BF-016", Name:"中部橡胶林区16", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:100, Ordinate:32, Open: "", FromId:"AF-016",Temperature: 27, Humidity: 0.3, Phenology: "稳定期", Speed: 5.8,Yellow: 0.15},
		FarmBInfo{Id:"BF-017", Name:"中部橡胶林区17", Date:"2018-12-11", Quality:"优", Healthy:"健康", Soil:"", Abscissa:12, Ordinate:36, Open: "", FromId:"AF-017",Temperature: 15, Humidity: 0.2, Phenology: "抽芽期", Speed: 2.6,Yellow: 0.15},
		FarmBInfo{Id:"BF-018", Name:"中部橡胶林区18", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:12, Ordinate:96, Open: "", FromId:"AF-018",Temperature: 16, Humidity: 0.1, Phenology: "抽芽期", Speed: 2.9,Yellow: 0.15},
		FarmBInfo{Id:"BF-019", Name:"中部橡胶林区19", Date:"2018-12-11", Quality:"良", Healthy:"健康", Soil:"", Abscissa:35, Ordinate:36, Open: "", FromId:"AF-019",Temperature: 17, Humidity: 0.3, Phenology: "变色期", Speed: 5.3,Yellow: 0.15},
		FarmBInfo{Id:"BF-020", Name:"中部橡胶林区20", Date:"2018-12-11", Quality:"优", Healthy:"有死皮", Soil:"", Abscissa:75, Ordinate:15, Open: "", FromId:"AF-020",Temperature: 20, Humidity: 0.2, Phenology: "变色期", Speed: 7.0,Yellow: 0.15},
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

func(t* FarmB)Invoke(stub shim.ChaincodeStubInterface) peer.Response {
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
	}else if funcName == "judge" {
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
	var temp FarmBInfo

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



func(t* FarmB)query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	keyID := args[0]
	value, error := stub.GetState(keyID)
	if error != nil {
		return shim.Error("GetState fail...")
	}
	return shim.Success(value)
}

func(t* FarmB)judge(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}
	var result string
	var midItem FarmBInfo
	json.Unmarshal(text.Payload, &midItem)


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
	name = midItem.Name

	temperature = midItem.Temperature
	humidity  = midItem.Humidity
	speed = midItem.Speed
	phenology = midItem.Phenology

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
	midItem.Method = methods
	jsonText, error := json.Marshal(midItem)
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
//判定是否可割胶以及湿润度
func(t* FarmB)isopen(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	text := t.query(stub, args)
	if text.Status != shim.OK {
		return shim.Error("query error ......")
	}

	var midItem FarmBInfo
	json.Unmarshal(text.Payload, &midItem)


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
	name = midItem.Name

	temperature = midItem.Temperature
	humidity  = midItem.Humidity
	speed = midItem.Speed
	healthy = midItem.Healthy
	yellow = midItem.Yellow

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
		midItem.Open = "可割胶"
		jsonText, error := json.Marshal(midItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	} else{
		midItem.Open = "不可割胶"
		jsonText, error := json.Marshal(midItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}

	if humidity >= 0 && humidity <0.25 {
		midItem.Soil = "干旱"
		jsonText, error := json.Marshal(midItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}
	}else if humidity >=0.25 && humidity < 0.50{
		midItem.Soil = "半干旱"
		jsonText, error := json.Marshal(midItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.50 && humidity < 0.75{
		midItem.Soil = "半湿润"
		jsonText, error := json.Marshal(midItem)
		if error != nil {
			return shim.Error("json.Marshal(myList) fail...")
		}
		error0 := stub.PutState(id, []byte(jsonText))
		if error0 != nil {
			return shim.Error("PutState fail...")
		}

	}else if humidity >=0.75 && humidity <= 1{
		midItem.Soil = "湿润"
		jsonText, error := json.Marshal(midItem)
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

// 根据keyID查询历史记录
func(t* FarmB)gethistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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

func main() {
	error := shim.Start(new(FarmB))
	if error != nil {
		println("程序启动失败...")
		return
	}
	fmt.Println("程序启动成功...")
}