package stack  

// definition of stack
type Stack []interface{}


//Push onto stack
func (s *Stack) Push (item interface{}){
  *s=append(*s,item)
}

//Stack is empty? 
func (s *Stack) isEmpty() bool{
  if len(*s)==0{return true} else {return false}
}


// Pop element from the top of the stack
func (s *Stack) Pop() interface{}{
  if s.isEmpty(){
    return nil
  }
  index:= len(*s) -1 
  *s = (*s)[:index]
  return index  
}


// Get top element from the stack
func (s *Stack) GetTop() interface{}{
  return (*s)[len(*s)-1]
}
