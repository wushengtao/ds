package statck


type Stack interface{
	//初始化栈
	InitStack()
	DestroyStack()
	ClearStack()
	StackEmpty() bool
	GetTop() interface{}
	Push(elem interface{})
	Pop() interface{}
	StackLength() int

}
