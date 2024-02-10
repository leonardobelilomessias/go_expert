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


func FuzzCalculateTax(f *testing.F) {
	seed:= []float64{-1,-2,500.0, 500.0,1501,2}
	for _, amount := range seed{
		f.Add(amount)
	} 
	f.Fuzz(func(t *testing.T, amount float64){
		result := CalculateTax(amount)
		if amount<=0 && result !=0{
			t.Errorf("Receiverd %f but expeted 0", result)
		}
		if amount>20000 && result!=20{
			t.Errorf("Receiverd %f but expeted 20", result)

		}
	})
}