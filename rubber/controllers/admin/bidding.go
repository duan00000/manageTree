/**
  @Author ZXQ-Administrator
  @Date 2021-12-02
  @node:
**/
package admin

import "fmt"

//1）可避免输配电容量成本
//可避免输配电建设费用 B G，c 可通过可避免输配
//电容量单位成本 cG 及实际避免输配电容量 ΔP 来确
//定［26］，表达式如下：
//B G，c = cG ΔP （2）
//由于输配电过程中存在电能损耗，导致用户参
//与需求响应反映在电网侧的负荷削减量与用户的响
//应功率并不相等。实际避免输配电容量 ΔP 为：
//ΔP = A 1/ a
//1 - α （3）
//式中：A1为第 1 时段用户的总响应功率；α 为电网输
//配电损失系数
//为了简化问题，假定在 k 时段，各用户的响应功
//率保持不变，则用户的总响应功率如下：
//A k = N∑i=1 x k i /t
//（4）
//式中：t为每个峰荷时段的持续时长
type BiddingPriceController struct {
	BaseController
}

//λ为现阶段需求响应周期对应的输配电价 la

//电网公司可避免输配电容量单位成本Cg = 100元/(kW·a);电网输配电损失系数 α=6%；市场平均电价 λa=0.85 元/（kW·h;输配电价λt = 0.15 元/（kW·h）
func (c *BiddingPriceController) Get() {
	var λa, α, cG, λt, Cgedr, t float64 //λ
	var yG = [20]float64{18.54, 14.85, 8.12, 6.18, 5.01, 3.52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(yG)

	cG = 100
	t = 1
	λa, α, λt = 0.85, 0.06, 0.15
	K := cG / (1 - α) //106.382979
	// 求Bgc 可避免输配电容量成本
	A := [...]float64{10531382.99, 10460918.42, 10374002.12, 10324546.45, 10318086.37, 10282741.68, 10239042.73, 10211158.09, 10170824.05,
		10157273.06, 10120250.64, 10117391.91, 10089080.16, 10057365.03, 10046868.36, 10042367.87, 10022226.88, 10012479.78, 10010756.73, 10008416.9}
	Bgc := K * A[0]
	// 求Bge 可避免电量成本
	cSum := c.Sum(A, 20)
	Bge := (λa*α - λt) * cSum * t
	// 求Cge 减少的输配电收益 + 需求响应补偿成本Cgdr =Cgedr
	ifyG := K/t + (λa*α - λt)
	for k := 1; k <= 20; k++ {
		if 0 < yG[1] && yG[1] < ifyG {
			Cgedr = yG[1] * cSum * t
		} else {
		}
	}

	//电网公司的需求响应收益函数Fg
	Fg := Bgc + Bge - Cgedr
	fmt.Println("电网公司的需求响应收益函数Fg",Fg)
	//用户 i 在 k 时段内的需求响应收益函数如下Fci
	//λk r =
	//1 元/（kW·h）；用户 1 到用户 3 的响应成本系数分
	//别 为 a1 = a2 = a3 =0.000 1 元 /（kW·h）2 ，b1 =
	//1 元/（kW·h），b2=2 元/（kW·h），b3=3 元/（kW·h）。
	//为了简化算例，在此设定同一用户在不同时段的响
	//应能力相同（实际上各用户在不同响应时段的最大
	//响应能力并不相同），分别为 H 1k=250 000 kW·h，
	//H 2k=450 000 kW·h，H 3k=600 000 kW·h。
	var x1, x2, x3, ki, bt, Fc1, Fc2, Fc3 float64
	x1, x2, x3 = 200000, 400000, 550000
	ki, bt = 1, 1
	Fc1, Fc2, Fc3 = 0.0, 0.0, 0.0
	var Fc [20] float64
	a := [3]float64{0.0001, 0.0001, 0.0001}
	b := [3]float64{1, 2, 3}
	//H := []float64{250000, 450000, 600000}
	for k := 1; k <= 20; k++ {
		if 0 < yG[1] && yG[1] < ifyG {
			Fc1 = (-1)*a[0]*(x1*x1) + (yG[1]+ki*bt-b[0])*x1
			Fc2 = (-1)*a[1]*(x2*x2) + (yG[1]+ki*bt-b[1])*x2
			Fc3 = (-1)*a[2]*(x3*x3) + (yG[1]+ki*bt-b[2])*x3
			Fc[k]=Fc1+Fc2+Fc3
		} else {
		}
	}
}
