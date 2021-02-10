package main

type Error interface {
	 Caller() []CallerInfo
	 Wraped() []error
	 Code() int
	 error
	 private()
}

type CallerInfo struct{
	FuncName string
	FileName string
	FileLine int
}
