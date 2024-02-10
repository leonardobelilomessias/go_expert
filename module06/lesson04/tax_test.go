// working with benchmarking

// use commnad bellow to run bechmark
//go test -bench=. 

//to ignore test jusr run benchmark code below
//go test -bench=.  -run=^#

//to ignore test just run benchmark code below with 10 elements screen
//go test -bench=.  -run=^# -count=10
package tax
 
import "testing"
func TestCalculateTax(t *testing.T){
	amount :=500
	expected :=5.0
	result :=CalculateTax(float64(amount))
	if result != expected{
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
func TestCalculateTaxBatch(t *testing.T){
	
	type calcTax struct{
		amount, expect float64
	}

	table := []calcTax{
		{500,5.0},
		{1000.0,10.0},
		{15000.0,10.0},

	}
	for _, item:= range table{
		result:= CalculateTax(item.amount)
		if result != item.expect{
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}


func BenchmarkCalculateTax(b *testing.B){

	for i:= 0 ; i< b.N; i++{

		CalculateTax(500.0)
	}
}
func BenchmarkCalculateTax2(b *testing.B){

	for i:= 0 ; i< b.N; i++{

		CalculateTax2(500.0)
	}
}

