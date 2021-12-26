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
	is bool//当前start指向的房间是否有楼梯
}

func (tr *targetRoom) next(){//到下一个有楼梯的房间
	if tr.v[tr.start].empty==1{//处理一开始就是有楼梯的情况
		if tr.is==true{
			tr.is=false
			tr.start++
		}else{
			tr.is=true
			return
		}
	}

run:
	for i:=tr.start;i<len(tr.v);i++ {
		if tr.v[i].empty==1{
			tr.is=true
			tr.start=i
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
	var passwd int

	troom.start=entry

	for i:=0;i<n;i++{

		passwd+=t[i][entry].board
		troom.v=t[i]

		//fmt.Printf("\nv:%d %d %d\n",i,entry,t[i][entry].board)
		for j:=0;j<t[i][entry].board;j++{
			troom.next()
			//fmt.Printf("搜索结果 %d ",troom.start)
		}

		entry=troom.start
	}

	fmt.Printf("%d",passwd)
}