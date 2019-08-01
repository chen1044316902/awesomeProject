package main

import "fmt"


//go 的struct是值类型
type user struct {
	name string
	email string
}


//使用值接受者
func (u user) notify(){
	fmt.Printf("Sending User Email To %s<%s>\n",u.name,u.email)
}

//使用指针接受者
func (u *user) changeEmail(email string){
	u.email = email
}


func main(){

	//user类型的值可以调用
	bill := user{"Bill","bill@email.com"}
	bill.notify()


	//指向user类型值的指针也可以调用
	lisa := &user{"Lisa","lisa@mailcom"}
	lisa.notify()

	//user 类型的值也可以调用
	bill.changeEmail("bill@newdomian.com")
	bill.notify()


}
