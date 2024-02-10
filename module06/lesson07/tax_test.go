// working with mocks
package tax

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestCalculateTax(t *testing.T){
	tax, err:= CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t,10.,tax)
	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater then 0")
	assert.Equal(t,0.0, tax)
}
func TestCalculateTaxSave(t *testing.T){
	repository:= &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax",0.0).Return((errors.New("error save tax")))
	
	err:= CalculateTaxAndSave(1000.0,repository)
	assert.Nil(t, err)
	err= CalculateTaxAndSave(0.0,repository)
	assert.Error(t , err,"error save tax!")
	repository.AssertExpectations(t)
}