package main

/*
/ We use < > to specify Parameter type
class Test<T>
{
    // An object of type T is declared
    T obj;
    Test(T obj) {  this.obj = obj;  }  // constructor
    public T getObject()  { return this.obj; }
}

// Driver class to test above
class Main
{
    public static void main (String[] args)
    {
        // instance of Integer type
        Test <Integer> iObj = new Test<Integer>(15);
        System.out.println(iObj.getObject());

        // instance of String type
        Test <String> sObj =
                          new Test<String>("GeeksForGeeks");
        System.out.println(sObj.getObject());
    }
}
 */
import (
	"fmt"
	"reflect"
)

type Test interface{}
type Test2 interface{}

type AnyT struct {
	objs  Test
	objs2 Test2
}

func NewTest(obj Test, objs2 Test2) AnyT {
	a := AnyT{}
	a.objs = obj
	a.objs2 = objs2
	return a
}

func (t AnyT) GetObject() Test {
	return t.objs
}

func (t AnyT) GetId() Itype {
	return t.objs
}

func main() {
	var ae = Test("kj")
	var ae2 = Test2(2)
	var a = NewTest(ae, ae2)
	fmt.Println(reflect.TypeOf(a.objs).Name())
	fmt.Println(reflect.TypeOf(a.objs2).Name())
}

type Itype interface{}

type IAccount interface {
	GetId() Itype
}
