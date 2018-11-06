package main

import ("fmt"
		"encoding/json"
		)

type attacker struct {
	attackpower int
	dmgbonus int
}

type sword struct {
	attacker
	twohanded bool
}

type gun struct {
	attacker
	bulletsremaining int
}

type weapon interface {
	Wield() bool
}

func wielder(w weapon) bool {
	fmt.Println("Wielding...")
	return w.Wield()
}

func (s sword) Wield() bool {
	fmt.Println("You've wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You've wielded a gun!")
	return true
}

type person struct {
	First string
	Last string
}


func Multiply(a, b int, reply *int) {
 	*reply = a * b
}

func Min(a ...int) int {
 	if len(a)==0 {
 		return 0
	}
 	min := a[0]
 	for _, v := range a {
 		if v < min {
 			min = v
 		}
 	}
 	return min
}


func  fp(a *[3]int) {
	fmt.Println(a) 
}

type File struct {
	Fd int
	Name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}


func add1(a int, b int) int {
	return a + b
}

func add2(a string, b string) string {
	return a + b
}

type TwoInt struct {
	a int
	b int
}

type TwoString struct {
	a string
	b string
}

func (ti *TwoInt) Add() int {
	return ti.a + ti.b
}


func (ts *TwoString	) Add() string {
	return ts.a + ts.b
}

func (ts TwoString) Join() {
	ts.a = "a"
	ts.b = "b"
	return
}

func main() {
	sword1 := sword{attacker: attacker{attackpower: 1, dmgbonus: 5}, twohanded: true}
	gun1 := gun{attacker: attacker{attackpower: 10, dmgbonus: 20}, bulletsremaining: 11}
	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)
	//sword1.Wield()
	//gun1.Wield()
	wielder(sword1)
	wielder(gun1)

    j1 := `{"First": "James0", "Last": "Bond0"}`
    j2 := `[{"First": "James0", "Last": "Bond0"}, {"First": "James1", "Last": "Bond1"}]`

    xp1 := person{}
    xp2 := []person{}

    err := json.Unmarshal([]byte(j1), &xp1)
	if err != nil{
		fmt.Println(err)	
	}
	fmt.Printf("go data: %+v\n", xp1)

    err = json.Unmarshal([]byte(j2), &xp2)
	if err != nil{
		fmt.Println(err)	
	}
	fmt.Printf("go data: %+v\n", xp2)

	var i1 int = 1
	var p1 *int
	p1 = &i1
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("*p1", *p1)

	for k, v := range xp2{
		fmt.Println(k, v)
	}

	n := 0
	reply := &n
	Multiply(1, 2, reply)



	fmt.Println(*reply)

	arr1 := []int{7,9,3,5,1}
	fmt.Printf("The minimum in the array arr is: %d \n", Min(arr1...))

    arr2 := []string{"abc", "bcd", "cde"}
    for _, v := range arr2{
    	fmt.Println(v)
    }

    var ar [3]int
    fp(&ar)

    var arr3 [6]int
    for i :=0; i< len(arr3); i++ {
    	arr3[i] = i
    }
    var slice1 []int = arr3[2:5] 
    for i := 0; i<len(slice1); i++ {
    	fmt.Println(slice1[i])
    }

    var mapLit map[string]int
    mapLit = map[string]int{"one": 1, "two": 2}

    mapCreated := make(map[string]float32)
    mapAssigned := mapLit
    fmt.Println(mapAssigned["one"])
    mapCreated["key1"] = 4.5
    for k, v := range mapCreated{
    	fmt.Println(k, v)
    }

    var isPresent bool
    _, isPresent = mapLit["two"]
    if isPresent{
    	fmt.Println(mapLit["two"])
    }

    var p2 *person = new(person)
    p2.First = "James2"
    p2.Last = "Bond2"
    fmt.Println(*p2)

    f := NewFile(10, "tmp1.txt")
    fmt.Println(f)

    fmt.Println(add1(1,2))
    fmt.Println(add2("abc","def"))

    var ti1 *TwoInt	= new(TwoInt)
    ti1.a = 3
    ti1.b = 5
    fmt.Println(ti1.Add())
    ti2 := TwoInt{3, 4}
	fmt.Println(ti2.Add()) // Pointer and value methods can both be called on pointer or non-pointer values

	ts1 := TwoString{"aaa", "bbb"}
	fmt.Println(ts1.Add())
	ts2 := TwoString{"ccc", "ddd"}
	fmt.Println((&ts2).Add())
	ts2.Join()
	fmt.Println(ts2.a, ts2.b)

}

