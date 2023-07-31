package main

import (
	"log"
	"os"

	"github.com/HusskyAngel/ZincIndexer/config"
	"github.com/HusskyAngel/ZincIndexer/internal/option1"
	"github.com/HusskyAngel/ZincIndexer/internal/option2"
	"github.com/HusskyAngel/ZincIndexer/internal/optionutils"
)


func main() {

  if (config.GetConfig().EnronDataPath =="<your path to enron data>/ "){

    log.Fatal("Your need especify the path to the uncompressed enron data in ./config/config.toml\n"+
              " Example:<your path to enron data>/ ")

  }else if len(os.Args)==1{

    log.Fatal("You need to specify with which method the data will be sent\n"+ 
                " Method 1: uses orchestator (paralellism)\n"+
                " Method 2: doesnt use orchestator (not paralellism)\n"+
                "\n Example: go run . 1 \n")

  }else if os.Args[1]=="1"{

    option1.Option1()

  }else if os.Args[1]=="2"{

    option2.Option2()
  
  }else if os.Args[1]=="3" {

    optionutils.SafeAccessStack().Stack.Push("Hola")
    log.Println(optionutils.SafeAccessStack().Stack.GetTop())

  }else{

    log.Fatal("The option "+ os.Args[1]+ " doesnt exist"+
              "The avaiable options are:\n"+
                " 1: uses orchestator (paralellism)\n"+
                " 2: doesnt use orchestator\n"+
                "\n Example: go run . 1 \n")

  } 

  
}
