package mreader

import (
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

/*
Mime Email Reader
*/

//read email file from path and set in Email structure
func MReader(path string) Email{
  
  content,errIO:=ioutil.ReadFile(path)
  if errIO!= nil{
    log.Fatal("Error loading the file")
  }
  fileContent:=string(content)

  msg,errE:=mail.ReadMessage(strings.NewReader(fileContent))
  if errE != nil{
    log.Fatal("Error loading email")
  } 

  var email Email

  email.messageId=msg.Header.Get("Message-ID")
  email.date=msg.Header.Get("Date")
  email.from=msg.Header.Get("From")
  email.to=msg.Header.Get("To")
  email.subject=msg.Header.Get("Subject")

  body,errB:=ioutil.ReadAll(msg.Body)
  if errB!=nil{
    log.Panic("Error loading body")
  }else{
    email.body=string(body)
  }

  return email
}



