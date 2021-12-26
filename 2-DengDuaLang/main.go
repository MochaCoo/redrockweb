package main

import "fmt"

type room struct{
	//id int
	empty int
	board int
}
type targetRoom struct{//用于按顺序遍历某层楼的所有目标房间
	v []room
	start int//当前访问的位置
}
func (tr *targetRoom) next(){//到下一个有楼梯的房间
run:
	for i:=tr.start;i<len(tr.v);i++ {
		if tr.v[i].empty==1{
			tr.start=i+1
			return
		}
	}
	//执行到这里说明tr.start后都没有有楼梯的房间了
	tr.start=0
	goto run
}
func main(){
	var n,m int
	var entry int
	//N层楼 M个房间
	fmt.Scanf("%d %d",&n,&m)
	var t [][]room
	t=make([][]room,n)
	for i:=0;i<n;i++{
		t[i]=make([]room,m)
	}
	for i:=0;i<n;i++{
		for j:=0;j<m;j++ {
			fmt.Scanf("%d %d", &(t[i][j].empty), &(t[i][j].board))
		}
	}

	fmt.Scanf("%d",&entry)


	var troom targetRoom
	troom.start=t[0][entry].board
	var passwd int


	//.v=t[0]
	passwd+=t[0][entry].board
	for i:=1;i<n;i++{
		troom.v=t[i]
		passwd+=t[i][entry].board
		fmt.Printf("vv:%d %d %d\n",i,entry,t[i][entry].board)
		for j:=0;j<t[i][entry].board;j++{
			troom.next()
		}
		entry=troom.start
	}

	fmt.Printf("result:%d",passwd)
}