package validate

import (
	"fmt"
	"log"
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
	log.Println("Error to validate struct : " + structName)
	for i := range validationErrors {
		log.Println(strconv.Itoa(i+1) + " - " + validationErrors[i])
	}

	return fmt.Errorf(fmt.Sprintf("%v", validationErrors))
}
