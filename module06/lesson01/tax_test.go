package tax
 
import "testing"
func TestCalculateTax(t *testing.T){
	amount :=500
	expected :=6.0
	result :=CalculateTax(float64(amount))
	if result != expected{
		t.Errorf("Expected %f but got %f", expected, result)
	}
}