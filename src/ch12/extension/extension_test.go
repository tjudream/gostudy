package extension

import (
	"fmt"
	"testing"
)
/**
17 扩展与复用
*/
type Pet struct{

}
func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("Wang!")
}

//func (d *Dog) SpeakTo(host string) {
//	//d.p.SpeakTo(host)
//	d.Speak()
//	fmt.Println(" ", host)
//}

func TestDog(t *testing.T) {
	//dog := new(Dog)
	//var dog Pet = new(Dog) //cannot use new(Dog) (type *Dog) as type Pet in assignment
	var dog *Dog = new(Dog)
	//var p = (*Pet)(dog)//cannot convert dog (type *Dog) to type *Pet
	dog.Speak()
	dog.SpeakTo("Chao")
}