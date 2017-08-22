package main

import (
	"fmt"
)

func main() {
	var general interface{}
	//general = 6.6
	//type_cast(general)
	general = 2
	type_cast(general)
}

func type_cast(general interface{}) {

	//var i int = general.(int)类型强制转换
	//fmt.Println(i)
	m := general.(type)

	switch m := general.(type) {
	case int:
		{
			//fmt.Println(m)
			break
		}
	case float32:
		{
			fmt.Println(m)
			break
		}
	case float64:
		{
			fmt.Println(m)
			break
		}
	default:
		fmt.Println("unknown type")
		break
	}
	/*
		switch m {
		case int:

			fmt.Println("the general type is int")
			newInt, ok := general.(int)
			check_convert(ok)
			fmt.Println("newInt 的值本来是", newInt)
			newInt += 2
			fmt.Println("加2后，结果是", newInt)
			newInt -= 6
			fmt.Println("接着减6后，结果是", newInt)
			newInt *= 4
			fmt.Println("然后乘4，结果是", newInt)
			newInt /= 3
			fmt.Println("最后除3，结果是", newInt)
			fmt.Println()
			fmt.Println()

		case float32:

			fmt.Println("the general type is float32")
			newFloat32, ok := general.(float32)
			check_convert(ok)
			fmt.Println("newFloat32 的值本来是", newFloat32)
			newFloat32 += 2.0
			fmt.Println("加2.0后，结果是", newFloat32)
			newFloat32 -= 6.0
			fmt.Println("接着减6.0后，结果是", newFloat32)
			newFloat32 *= 4.0
			fmt.Println("然后乘4.0，结果是", newFloat32)
			newFloat32 /= 3.0
			fmt.Println("最后除3.0，结果是", newFloat32)
			fmt.Println()
			fmt.Println()

		case float64:

			fmt.Println("the general type is float64")
			newFloat64, ok := general.(float64)
			check_convert(ok)
			fmt.Println("newFloat64 的值本来是", newFloat64)
			newFloat64 += 2.0
			fmt.Println("加2.0后，结果是", newFloat64)
			newFloat64 -= 6.0
			fmt.Println("接着减6.0后，结果是", newFloat64)
			newFloat64 *= 4.0
			fmt.Println("然后乘4.0，结果是", newFloat64)
			newFloat64 /= 3.0
			fmt.Println("最后除3.0，结果是", newFloat64)
			fmt.Println()
			fmt.Println()

		default:
			fmt.Println("unknown type")
		}
	*/
}

func check_convert(ok bool) {
	if false == ok {
		panic("type cast failed!")
	}
}

/*
import (
	"fmt"
)

type a interface {
	Test()
}

type b struct {
	K uint32
}

func (self *b) Test() {
	fmt.Printf("k1:%d\n", self.K)
}

//////////////////////////////////////////////
type c struct {
	kc a
}

func (self *c) Set(l a) {
	self.kc = l
}

func (self *c) Get() {
	self.kc.Test()
}

func main() {
	var Yb b
	Yb.K = 2
	var j *b = &Yb
	var Yc c
	Yc.Set(j)
	Yc.Get()
}
*/

/*
import (
	"fmt"
	"test3/common"
)

var msg chan interface{}

func Read() {
	for {
		select {
		case r := <-msg:
			{
				fmt.Print("sss\n")
				switch m := r.(type) {
				case *common.Login:
					var msg *common.Login = m
					fmt.Printf("r(common.Login)uid=%d\n", (msg.Uid))
				case int:
					var msg int = m
					fmt.Printf("r(int) = %d\n", msg)
				case string:
					var msg string = m
					fmt.Printf("r(string) = %s\n", msg)
				default:
					return
				}

			}
		}
	}
}

func Get() *common.Login {
	var a common.Login
	var value int = 1
	a.Uid = &value
	return &a
}
func Write1() {
	//var a common.Login
	//var value int = 1
	//a.Uid = &value
	var a *common.Login = Get()
	msg <- a
}
func Write2() {
	var a string = "aaaa"
	msg <- a
}
func main() {
	msg = make(chan interface{}, 20)
	go Write1()
	go Read()
	go Write2()
	for {

	}
}
*/
/*
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	e3 := l.InsertBefore(3, e4)
	k1 := l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Printf("=============\n")
	l.Remove(k1)
	l.InsertBefore(5, e3)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
1
2
3
4
=============
1
5
3
4
*/

/*
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Cap() int           { return cap(h) }
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	var h IntHeap
	heap.Init(&h)
	heap.Push(&h, 1)
	heap.Push(&h, 2)
	heap.Push(&h, 3)
	heap.Push(&h, 4)
	heap.Push(&h, 5)
	fmt.Printf("len:%d, cap:%d\n", h.Len(), h.Cap())
	for lIt := range h {
		fmt.Print(lIt)
	}
	/*
		for h.Len() > 0 {
			fmt.Printf("%d ", heap.Pop(h))
			fmt.Printf("len:%d, cap:%d\n", h.Len(), h.Cap())
		}
*/
//}

/*
输出：
len:5, cap:8
1 len:4, cap:8
2 len:3, cap:8
3 len:2, cap:8
4 len:1, cap:8
5 len:0, cap:8
*/
/*
import (
	"fmt"
	"time"
)

func main() {
	lTimeNow := time.Now()
	lYear := lTimeNow.Year()
	lMonth := lTimeNow.Month()
	lDay := lTimeNow.Day()

	lHour := lTimeNow.Hour()
	lMin := lTimeNow.Minute()
	lSecond := lTimeNow.Second()

	lYear, lMonth, lDay := lTimeNow.UTC().Date()
	lHour, lMin, lSecond := lTimeNow.UTC().Clock()
	var tmp int
	tmp = lMonth
	fmt.Print(tmp)
	fmt.Printf("lYear:%d, lMonth:%d, lDay:%d, lHour:%d, lMin:%d, lSecond:%d", lYear, lMonth, lDay, lHour, lMin, lSecond)
}
*/

/*
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	a := []byte("aaaaaa")
	var b []byte
	var c []byte
	lTmpBuf := bytes.NewBuffer(b)
	var lenTmp int16 = int16(len(a))
	fmt.Printf("len1:%d\n", lenTmp)
	binary.Write(lTmpBuf, binary.BigEndian, lenTmp)
	d := lTmpBuf.Bytes()
	fmt.Printf("len2:%d\n", len(d))
	for _, tmp := range d {
		c = append(c, tmp)
	}
	fmt.Printf("len3:%d\n", len(c))
	fmt.Print(c)
	for _, tmp := range a {
		c = append(c, tmp)
	}
	fmt.Printf("\nlen4:%d\n", len(c))
	fmt.Print(c)

}*/

/*
package main

import (
	"fmt"
	"reflect"
)

func say(text1 string, text2 string) {
	fmt.Println(text1)
	fmt.Println(text2)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	f_type := reflect.TypeOf(m[name])

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
		fmt.Println(k)
	}
	result = f.Call(in)

	return
}

func main() {
	var funcMap = make(map[string]interface{})
	funcMap["say"] = say
	Call(funcMap, "say", "hello", "world")
}
*/

/*
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Title string
}

func (this *User) Test(id1 string, id2 int32) string {
	fmt.Println(id1)
	fmt.Println(id2)
	return this.Name
}

func main() {
	o := &User{1, "Tom", 12, "nan"}
	vType := reflect.TypeOf(o)
	fmt.Printf("MethodBNum:%d\n", vType.NumMethod())
	for m := 0; m < vType.NumMethod(); m++ {
		mFunc := vType.Method(m)
		mFuncType := mFunc.Type
		function := mFunc.Func
		fmt.Printf("-----ParamNum:%d\n", mFuncType.NumIn())
		k := "hanfeng"
		var k1 int32
		k1 = 1
		retss := function.Call([]reflect.Value{reflect.ValueOf(o), reflect.ValueOf(k), reflect.ValueOf(k1)})
		fmt.Println(retss)

		idType := mFuncType.In(2)

		if !idType.AssignableTo(reflect.TypeOf((*int32)(nil)).Elem()) {
			fmt.Printf("Failed\n")
		} else {
			fmt.Printf("Success\n")
		}
	}
	/*
		k := "hanfeng"
		v := reflect.ValueOf(o)
		m := v.MethodByName("Test")
		rets := m.Call([]reflect.Value{reflect.ValueOf(k)})
		fmt.Println(rets)
*/
//}

/*package main


import (
	"fmt"
	"reflect"
)

type Test1 struct {
}

func main() {
	Register(func(args interface{}) {})
}

type func1 func(interface{})

func Register(f interface{}) {

	TypeTest := reflect.TypeOf(f)
	fmt.Println(TypeTest)


		switch TypeTest{
			case func(interface{}):
			fmt.Println("ddd")
			default:
		}

}
*/

/*
import (
	"fmt"
	"reflect"
)

type Test1 struct {
}

func (a *Test1) func1() {
	fmt.Println("func1")
}
func (a *Test1) func2() {
	fmt.Println("func2")
}

func main() {
	Register()
}

func Register() {
	a := &Test1{}
	b_type := reflect.TypeOf(a)
	b_value := reflect.ValueOf(a)
	fmt.Println(b_type)
	fmt.Println(b_value)

	for m := 0; m < b_type.NumMethod(); m++ {
		fmt.Println(b_type.Method(m))//1
		fmt.Println(b_type.Method(m).Name)//2	函数名
		fmt.Println(b_type.Method(m).Index)//3	函数在集合中的序列号
		fmt.Println(b_type.Method(m).Type)//4	类型
		fmt.Println(b_type.Method(m).Func)//5	函数地址
		fmt.Println("=================")
	}
}

结果：
	*main.Test1
	&{}
	1:{func1 main func(*main.Test1) <func(*main.Test1) Value> 0}
	2:func1
	3:0
	4:func(*main.Test1)
	5:0x401040
	=================
	{func2 main func(*main.Test1) <func(*main.Test1) Value> 1}
	func2
	1
	func(*main.Test1)
	0x401140
	=================
*/
/*
func Register() {
	//结果：
	//func(interface {})
	//*main.Test1
	//&{}
	//main.Test1
	//1:我们把传递使用变量的地址
	//2：通过反射的Indirect(v Value)可以将变量指针解引用，还原成原来的变量使用
	//	*main.Test1 - - > main.Test1变化是解引用的强大功能

	m := &Test1{}
	m_Type := reflect.TypeOf(m) //指针类型，所以是*类型
	m_Value := reflect.ValueOf(m)
	fmt.Println(m_Type)
	fmt.Println(m_Value)
	var m1 = reflect.Indirect(m_Value).Type() //Indirect：将指针进行了解引用，然后在取值（去了指针后的）的类型。
	fmt.Println(m1)
}
*/

/*
import (
	"fmt"
	"reflect"
	"strconv"
)

type Record struct {
	Index int "index"
	Name  string
}

func main() {
	var Test1 Record = Record{}
	test(Test1)
}

type Index map[interface{}]interface{}

func test(st interface{}) {
	records := make([]interface{}, 3)
	TypeRecord := reflect.TypeOf(st)
	indexes := []Index{}
	index := 0

	for n := 1; n < 4; n++ {
		value_test := reflect.New(TypeRecord) //根据类型（反射得到的）创建一个Value的指针，指向的是存放数据地址的指针（实际就是类型指针）
		records[n-1] = value_test.Interface() //获取value保存的指针
		ValueRecord := value_test.Elem()      //获取指针指向的值,结构体每个字段的组成的。

		for i := 0; i < TypeRecord.NumField(); i++ {
			f := TypeRecord.Field(i)
			field := ValueRecord.Field(i)

			if f.Type.Kind() == reflect.Int {
				var v int64
				v, _ = strconv.ParseInt("12", 0, f.Type.Bits())
				field.SetInt(v)
			} else if f.Type.Kind() == reflect.String {
				field.SetString("hanfeng")
			}
			if f.Tag == "index" {
				indexdd := indexes[index]
				index++
				indexdd[field.Interface()] = records[n-1]
			}
		}

		fmt.Println(records[n-1])

	}

}
*/
