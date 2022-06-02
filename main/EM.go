package main

import (
	"fmt"
	"math"
)

/*
*   getPI
*  @Description: 获取p(x(i))的值和 p(1|xi)的值
*  @param data 当前条数据
*  @param dataP 𝛉数组
*  @param pi1 𝛑1
*  @param pi2 𝛑2
*  @return float64 p(x)
*  @return float64 p(1|x)
 */
func getPI(data []int, dataP [][]float64, pi1, pi2 float64) (float64, float64) {
	//两个分类分别计算概率的值
	params1, params2 := pi1, pi2
	for index, value := range data {
		//如果value为1 表示选取了当前值
		if value == 1 {
			params1 *= dataP[0][index]
			params2 *= dataP[1][index]
		} else {
			params1 *= 1 - dataP[0][index]
			params2 *= 1 - dataP[1][index]
		}
	}
	return params1 + params2, params1
}

/*
*   getResult
*  @Description:  获取聚类结果
*  @param pxTotal 计算出来的概率
*  @return []int  在第一类的值
*  @return []int  在第二类的值
 */
func getResult(pxTotal [][]float64) ([]int, []int) {
	var pi1 []int
	var pi2 []int
	for index, value := range pxTotal {
		if value[0] > value[1] {
			pi1 = append(pi1, index+1)
		} else {
			pi2 = append(pi2, index+1)
		}
	}
	return pi1, pi2
}

/*
*   compareResult
*  @Description:  判断两次结果是否相同 误差设置为0.01
*  @param pxTotal
*  @param pxTotalOld
*  @param pi1 新的pi1
*  @param pi2 新的pi2
*  @param pi1Old 上一次的pi1
*  @param pi2Old 上一次的pi2
*  @return bool 比较结果
 */
func compareResult(dataP, dataPOld [][]float64, pi1, pi2, pi1Old, pi2Old float64) bool {
	precision := 0.005
	//return reflect.DeepEqual(dataP, dataPOld) && pi1 == pi1Old && pi2 == pi2Old
	if len(dataP) != len(dataPOld) {
		return false
	}
	for i := 0; i < len(dataP); i++ {
		for j := 0; j < len(dataP[0]); j++ {
			if math.Abs(dataP[i][j]-dataPOld[i][j]) > precision {
				return false
			}
		}
	}
	return math.Abs(pi1-pi1Old) <= precision && math.Abs(pi2-pi2Old) <= precision
}

func main() {
	/**
	数据初始化 start
	*/
	//初始化𝛉数组
	dataP := make([][]float64, 2)
	dataP[0] = []float64{0.6, 0.6, 0.3, 0.5, 0.3}
	dataP[1] = []float64{0.4, 0.4, 0.7, 0.5, 0.7}
	// 初始化数据数组
	data := make([][]int, 9)
	data[0] = []int{1, 1, 0, 0, 1}
	data[1] = []int{0, 1, 0, 1, 0}
	data[2] = []int{0, 1, 1, 0, 0}
	data[3] = []int{1, 1, 0, 1, 0}
	data[4] = []int{1, 0, 1, 0, 0}
	data[5] = []int{0, 1, 1, 0, 0}
	data[6] = []int{1, 0, 1, 0, 0}
	data[7] = []int{1, 1, 1, 0, 1}
	data[8] = []int{1, 1, 1, 0, 0}
	//初始化𝛑1  𝛑2
	pi1, pi2, pi1Old, pi2Old := 0.6, 0.4, 0.6, 0.4

	var pi1Res []int
	var pi2Res []int
	var dataPOld [][]float64
	//计算聚类结果
	for l := 0; l < 100; l++ {
		// p数据
		var pxTotal [][]float64
		for _, value := range data {
			px, p1x := getPI(value, dataP, pi1, pi2)
			item := []float64{p1x / px, 1 - p1x/px}
			pxTotal = append(pxTotal, item)
		}
		//fmt.Printf("%+v\n", pxTotal)
		pi1ResNew, pi2ResNew := getResult(pxTotal)

		pi1Res = pi1ResNew
		pi2Res = pi2ResNew

		//计算𝛑1  𝛑2
		tempPi1 := 0.0
		tempPi2 := 0.0
		for _, value := range pxTotal {
			tempPi1 += value[0]
			tempPi2 += value[1]
		}
		//pi1, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tempPi1/float64(len(data))), 64)
		//pi2, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tempPi2/float64(len(data))), 64)
		pi1 = tempPi1 / float64(len(data))
		pi2 = tempPi2 / float64(len(data))
		//更新𝛑1  𝛑2 结束

		//更新𝛉
		pi1Total := 0.0
		pi2Total := 0.0
		//计算分母
		for _, value := range pxTotal {
			pi1Total += value[0]
			pi2Total += value[1]
		}
		for i := 0; i < len(dataP[0]); i++ {
			temp1 := 0.0
			temp2 := 0.0
			for j := 0; j < len(data); j++ {
				if data[j][i] == 1 {
					temp2 += pxTotal[j][1] * float64(data[j][i])
					temp1 += pxTotal[j][0] * float64(data[j][i])
				}
			}
			//dataP[0][i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temp1/pi1Total), 64)
			//dataP[1][i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temp2/pi2Total), 64)
			dataP[0][i] = temp1 / pi1Total
			dataP[1][i] = temp2 / pi2Total
		}

		fmt.Printf("\n第%d轮：\npi1=%f,pi2=%f\n", l, pi1, pi2)
		fmt.Printf("𝛑1 = %+v\t𝛑2 = %+v\n", pi1Res, pi2Res)
		fmt.Printf("%+v\n", dataP)
		fmt.Printf("%+v", dataPOld)
		//更新𝛉数组 完成
		resultEqual := compareResult(dataP, dataPOld, pi1, pi2, pi1Old, pi2Old)
		if resultEqual {
			break
		}
		//更新参数
		pi1Old, pi2Old, dataPOld = pi1, pi2, dataP
	}

}
