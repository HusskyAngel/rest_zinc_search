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

)

//Push onto the stack
func doOption1(filePath string ) {
  optionutils.SafeAccessStack().Lock()
  optionutils.SafeAccessStack().Stack.Push(filePath)
  optionutils.SafeAccessStack().Unlock()
}

//Get path from top of the stack, read the email and post the email 
func worker(){
  for{
    if optionutils.SafeAccessStack().Stack.IsEmpty(){
      continue
    }else{
        wg.Add(1)

        optionutils.SafeAccessStack().Lock()
        filePath:= optionutils.SafeAccessStack().Stack.GetTop()
        optionutils.SafeAccessStack().Stack.Pop()
        optionutils.SafeAccessStack().Unlock()

          mail,errM:=mailreader.Reader(fmt.Sprintf("%v",filePath))
          if errM!= nil{
            log.Println("Error: ", filePath, errM)  
          }else{
            errP:=post.PostEmail(mail)
            if errP!=nil{
              log.Println("Error: ",errP)
            }
          }

        wg.Done()
    }
  }
} 

//Option 1 
func Option1() {
  enronData:=config.GetConfig().EnronDataPath
  for i:=0;i<numWorkers;i++{
    go worker() 
  }
  pathsgetter.PathsGetter(enronData,doOption1)
  wg.Wait()

}
