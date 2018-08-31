package main

import (

	"fmt"
)

type SortInterface interface{
	sort()
}

type Sortor struct {
	name string
}
//交换元素顺序
func swap(a []int , i int, j int ){
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

/*
//选择排序的方法
func (sortor Sortor) sort (arr []int){

	for i:= 0; i< len(arr) ;i++  {
		min := i
		for j:= i+1; j<len(arr);j++{
			if arr[j]<arr[min]{
				min = j
			}
		}
		swap(arr,i,min)
	}
}
*/

//冒泡排序
func (sortor Sortor) sort(arr []int){
	for i:= 0;i<len(arr)-1 ;i++  {
		for j := i+1;j<len(arr)-i-1;j++{
			if arr[j] > arr[j+1]{
				swap(arr, j+1 ,j)
			}
			break

		}
	}
}


func main(){
	arr := []int{1,3,4,2,7,4,5,3}
	//learnsort := Sortor{name: " selctsort: from min to max"}
	learnsort := Sortor{name:"bubblesort: from min to max"}
	learnsort.sort(arr)
	fmt.Println(learnsort.name, arr)
}
