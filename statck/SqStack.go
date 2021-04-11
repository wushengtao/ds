package statck

import "errors"

//栈的顺序结构
type SqStack struct{
	top int
	size int
	elems []interface{}
}

func NewSqStack() *SqStack{
	stack:=new(SqStack)
	return stack
}

func (stack *SqStack) InitStack() {
	stack.top=0
	stack.elems=make([]interface{},0)
}

func (stack *SqStack) DestroyStack(){
}

func (stack *SqStack) ClearStack(){
	stack.top=0
	stack.size=0
	stack.elems=make([]interface{},0)
}

func (stack *SqStack) StackEmpty() bool{
	return stack.size==0
}

func (stack *SqStack) GetTop() interface{}{
	if stack.StackEmpty(){
		return nil
	}
	return stack.elems[stack.size-1]
}

func (stack *SqStack) Push(elem interface{}){
	stack.size++
	stack.elems=append(stack.elems,elem)
}

func (stack *SqStack) Pop() interface{}{
	if stack.StackEmpty(){
		return errors.New("stack is empty")
	}
	stack.size--
	elem:=stack.elems[stack.size]
	stack.elems=stack.elems[:stack.size]
	return elem
}

func (stack *SqStack) StackLength() int{
	return stack.size
}



