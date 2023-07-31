package optionutils

import (
	"sync"

	"github.com/HusskyAngel/ZincIndexer/pkg/stack"
)

//stack of paths
var (
  pstack *stack.Stack 
  once   sync.Once  
)

//get a singleton stack of paths 
func PStack() *stack.Stack{
  once.Do(func(){
    pstack=&stack.Stack{}
  })
  return pstack
}
