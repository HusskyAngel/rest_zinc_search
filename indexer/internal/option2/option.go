package option2

import (
	"log"

	"github.com/HusskyAngel/ZincIndexer/config"
	"github.com/HusskyAngel/ZincIndexer/pkg/mailreader"
	"github.com/HusskyAngel/ZincIndexer/pkg/pathsgetter"
	"github.com/HusskyAngel/ZincIndexer/pkg/post"
)

func doOption2(filePath string){

  mail,errM:=mailreader.Reader(filePath)
  if errM!= nil{
    log.Println("Error: ", filePath, errM)  
  }else{
    errP:=post.PostEmail(mail)
    if errP!=nil{
      log.Println("Error: ",errP)
    }
  }

}


func Option2() {

  enronData:=config.GetConfig().EnronDataPath
  pathsgetter.PathsGetter(enronData,doOption2)

}
