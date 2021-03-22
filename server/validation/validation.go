package validation

import (
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/helpers"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type Scenario string

type Rules string

type Field string

type ValidationMap map[Scenario]ValidationRules

type ValidationRules map[Field]Rules

type FieldError struct {
	Tag     string
	Field   Field
	Message string
	Value   string
	Param   string
	Name    string
}

type ValidationErrors map[Field][]FieldError

func (s ValidationErrors) Error() (result string) {
	for _, errs := range s {
		for _, e := range errs {
			result += e.Message + ";\n"
		}
	}
	return
}

func (s ValidationErrors) collectError(err error, model interface{}) {
	for _, e := range err.(validator.ValidationErrors) {
		field := Field(e.Field())
		if _, ok := s[field]; !ok {
			s[field] = make([]FieldError, 0)
		}
		validationError := FieldError{
			Name:  getType(model),
			Tag:   e.Tag(),
			Field: field,
			Value: fmt.Sprintf("%v", e.Value()),
			Param: e.Param(),
		}
		validationError.setErrorMessage()
		s[field] = append(s[field], validationError)
	}
}

func (s *FieldError) setErrorMessage() {
	switch s.Tag {
	case constants.RequiredTag:
		s.Message = fmt.Sprintf(constants.RequiredErrorMsg, s.Name, s.Field)
	case constants.MinTag:
		s.Message = fmt.Sprintf(constants.MinValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.MaxTag:
		s.Message = fmt.Sprintf(constants.MaxValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.EmailTag:
		s.Message = fmt.Sprintf(constants.EmailErrorMsg, s.Name)
	}
}

func ValidateByScenario(scenario Scenario, m interface{}, validationMap ValidationMap) ValidationErrors {
	errs := make([]error, 0)
	validate := validator.New()
	data := structToMap(m)
	for fieldName, validation := range validationMap[scenario] {
		field, ok := data[fieldName]
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		err := validate.Var(field, string(validation))
		if err != nil {
			errs = append(errs, err)
		}
	}
	return formatErrors(m, errs)
}

func formatErrors(model interface{}, errs []error) ValidationErrors {
	if len(errs) > 0 {
		result := make(ValidationErrors)
		for _, err := range errs {
			if err != nil {
				result.collectError(err, model)
			}
		}
		return result
	}
	return nil
}

func formatError(model interface{}, err error) ValidationErrors {
	if err != nil {
		result := make(ValidationErrors)
		result.collectError(err, model)
		return result
	}
	return nil
}

func getType(s interface{}) string {
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

// need pass ptr or interface instead of struct - otherwise func panics
func structToMap(input interface{}) map[Field]interface{} {
	r := make(map[Field]interface{})
	s := reflect.ValueOf(input).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r[Field(typeOfT.Field(i).Name)] = f.Interface()
	}
	return r
}

func Validate(model interface{}) error {
	v := validator.New()
	err := v.Struct(model)
	return formatError(model, err)
}
