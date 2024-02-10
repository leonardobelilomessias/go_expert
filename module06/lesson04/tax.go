// working with benchmarking
// use commnad bellow to run bechmark
// go test -bench=.
package tax

import "time"
func CalculateTax(amount float64)float64  {
	if amount >= 1000{
		return 10.0
	}
	if amount == 0{
		return 0.0
	}
	return 5.0
}

func CalculateTax2(amount float64)float64  {
	time.Sleep(time.Millisecond)
	if amount >= 1000{
		return 10.0
	}
	if amount == 0{
		return 0.0
	}
	return 5.0
}

// use comand bellow to running test and get feedback terminal
//go test -coverprofile=coverage.out . ,

// use comand bellow to running test and get feedback on browser 
//go tool cover -html=coverage.out" 