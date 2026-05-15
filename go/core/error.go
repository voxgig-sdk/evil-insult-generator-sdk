package core

type EvilInsultGeneratorError struct {
	IsEvilInsultGeneratorError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewEvilInsultGeneratorError(code string, msg string, ctx *Context) *EvilInsultGeneratorError {
	return &EvilInsultGeneratorError{
		IsEvilInsultGeneratorError: true,
		Sdk:              "EvilInsultGenerator",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *EvilInsultGeneratorError) Error() string {
	return e.Msg
}
