package optionutils

import (
	"sync"

	"github.com/HusskyAngel/ZincIndexer/pkg/stack"
)

type sharedStack struct{ 
   sync.RWMutex
   Stack stack.Stack
}


//stack of paths
var (
  saStack *sharedStack
  once   sync.Once  
)

//get a singleton stack of paths 
func SafeAccessStack() *sharedStack{
  once.Do(func(){
    saStack=&sharedStack{}
  })
  return saStack
} 




