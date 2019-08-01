package main

import (
	"fmt"
)

func main(){
	//先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,}
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray")
	for _, v := range myArray{
		fmt.Print(v ,",")
	}
	fmt.Println("\nelements of mySlice : ")
	for _,v := range mySlice{
		fmt.Print(v,",")
	}
	fmt.Println()


	//创建切片

	mySlice1 := make([] int , 5) ;
	for _,v := range mySlice1{
		fmt.Println(v)
		size := len(mySlice1)
		fmt.Println(size)
	}

	mySlice3 := []int{1, 2, 3, 4, 5}

	fmt.Println("mySlice3的数据类型为 =  %T\n ",mySlice3)

	//切片复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元z素到slice1的前3个位置
	fmt.Print(slice1)

}