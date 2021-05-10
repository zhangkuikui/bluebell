package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"fmt"
)

func GetCommunityList()(data []*models.Community,err error)  {
	//查数据

	data,err= mysql.GetCommunityList()
	//aa:=make([]string,0)
	fmt.Println("len(data):",len(data))
	aa:=make([]string,0,len(data))
	for _,s:= range data{
		aa = append(aa, s.Name)
		fmt.Println(s.Name)
		fmt.Println(s.ID)
	}
	fmt.Printf("aa:%#v\n",aa)
	return data,err

}

func GetCommunityDetail(id int64)(*models.CommunityDetail, error)  {

	return mysql.GetCommunityDetailByID(id)
}
