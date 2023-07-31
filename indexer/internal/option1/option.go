package option1

import (
	"github.com/HusskyAngel/ZincIndexer/config"
	"github.com/HusskyAngel/ZincIndexer/internal/optionutils"
	"github.com/HusskyAngel/ZincIndexer/pkg/pathsgetter"
)


//push onto the stack
func doOption1(filePath string ) {
  optionutils.PStack().Push(filePath)
}

func orchestator(){
  for{
    select{


    }
  }  
}


func Option1() {
  enronData:=config.GetConfig().EnronDataPath
  pathsgetter.PathsGetter(enronData,doOption1)

}
