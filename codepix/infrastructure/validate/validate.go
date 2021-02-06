package validate

import (
	"fmt"
	"strconv"

	"github.com/diegoclair/go_utils-lib/v2/validstruct"
)

//Struct - to handle struct validation
func Struct(dataSet interface{}, structName string) error {
	err := validstruct.ValidateStruct(dataSet)
	if err != nil {
		return handleValidationStructError(err.Causes(), structName)
	}
	return nil
}

func handleValidationStructError(err interface{}, structName string) error {

	validationErrors := err.([]string)
	fmt.Println("Error to validate struct : " + structName)
	for i := range validationErrors {
		fmt.Println(strconv.Itoa(i+1) + " - " + validationErrors[i])
	}

	return fmt.Errorf(fmt.Sprintf("%v", validationErrors))
}
