package main

import f "fmt"

func main() {
	arr:=[5]float64{1,2,3,4,5}
	x:=arr[0:4]
	y:=arr[:]
	z:=arr[1:3]

	f.Println(x,y,z)

	s:=make([]int,5,10)
	sl1:=[]int{1,2,3}
	sl2:=append(sl1,4,5,6,7,8,9,0)
	f.Println(s)
	copy(s,sl2)
	f.Println(sl1,sl2,s)

	m := make(map[string]int)
	m["key"] = 10
	f.Println(m["key"],"Len of m=",len(m))

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	f.Println(elements["Li"])

	if name1, result := elements["Ae"]; result {
		f.Println(name1, result)
	} else {
		f.Println("Not found",result)
	}

}