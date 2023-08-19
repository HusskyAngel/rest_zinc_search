package option1

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/HusskyAngel/ZincIndexer/config"
	"github.com/HusskyAngel/ZincIndexer/internal/optionutils"
	"github.com/HusskyAngel/ZincIndexer/pkg/mailreader"
	"github.com/HusskyAngel/ZincIndexer/pkg/pathsgetter"
	"github.com/HusskyAngel/ZincIndexer/pkg/post"
)

var (
  //Number of workers
	numWorkers   = runtime.NumCPU() - 1
  //Wait for the program to finish goroutines
  wg sync.WaitGroup 
  //channel for done path walk 
  doneChannel= make(chan bool)
)

//Push onto the stack
func doOption1(filePath string ) {
  optionutils.SafeAccessStack().Lock()
  optionutils.SafeAccessStack().Stack.Push(filePath)
  optionutils.SafeAccessStack().Unlock()
}

func getAndPop() string{
  defer optionutils.SafeAccessStack().Unlock()

  optionutils.SafeAccessStack().Lock()
  filePath:= optionutils.SafeAccessStack().Stack.GetTop()
  optionutils.SafeAccessStack().Stack.Pop()
  return fmt.Sprintf("%v",filePath)
}
//Get path from top of the stack, read the email and post the email 
func worker(){
  defer wg.Done()

  var done bool=false
  for{
    select{
    case inf:=<-doneChannel:
        done=inf
      default:
        if !optionutils.SafeAccessStack().Stack.IsEmpty(){
          filePath:=getAndPop()
          mail,errM:=mailreader.Reader(filePath)
          log.Println(filePath)
          if errM!= nil{
            log.Println("Error: ", filePath, errM)  
          }else{
            errP:=post.PostEmail(mail)
            if errP!=nil{
              log.Println("Error: ",errP)
            }
          }
        }else if optionutils.SafeAccessStack().Stack.IsEmpty() && done{
          return 
        }else{
          continue 
        } 
    }
  }
} 

//Option 1 
func Option1() {
  enronData:=config.GetConfig().EnronDataPath
  for i:=0;i<numWorkers;i++{
    wg.Add(1)
    go worker()
  }
  pathsgetter.PathsGetter(enronData,doOption1)
  doneChannel<-true
  wg.Wait()

}
