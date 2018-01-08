package x_update

import (
	"testing"
	"fmt"
	"encoding/json"
	"os"
)

func TestRun(t *testing.T) {
	Run()
	fmt.Println("End")
}

func BenchmarkRunWithLock(b *testing.B) {
	b.N = 100
	for i:=0; i< b.N;i++{
		RunWithLock()
	}
}

func BenchmarkRunWithoutLock(b *testing.B) {
	b.N = 100
	for i:=0; i< b.N;i++{
		RunWithoutLock()
	}
}

func TestA(t *testing.T) {
	type A struct {
		Age int `json:"age,string"`
	}
	type omit bool
	var a A
	json.Unmarshal([]byte(`{"age":"100"}`),&a)
	fmt.Println(a.Age)

	//json.Marshal(struct {
	//	*A
	//	OmitAge  omit `json:"age,ommitempty"`
	//
	//	Age int `json:"new_age"`
	//}{
	//	&a,
	//	Age: a.Age,
	//})

	json.NewEncoder(os.Stdout).Encode(
		struct {
			*A
			OmitAge  bool `json:"age,omitempty"`
			Age int `json:"new_age"`
		}{
		     A:&a,
			//Age: a.Age,
			Age:a.Age,
		})

}