package mreader

import (
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

//read email file from path and set in Email structure
func EReader(path string) Email{
  
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

  email.MessageId=msg.Header.Get("Message-ID")
  email.Date=msg.Header.Get("Date")
  email.From=msg.Header.Get("From")
  email.To=msg.Header.Get("To")
	email.CC=msg.Header.Get("CC")
  email.MimeVersion=msg.Header.Get("Mime-Version")
  email.ContentType=msg.Header.Get("Content-Type")
  email.ContentTransferEncoding=msg.Header.Get ("Content-Transfer-Encoding") 
  email.XFrom=msg.Header.Get("X-From")
  email.XTo=msg.Header.Get("X-To")
  email.Xcc=msg.Header.Get("X-cc")
  email.Xbcc=msg.Header.Get("X-bcc")
  email.XFolder=msg.Header.Get("X-folder")
  email.XOrigin=msg.Header.Get("X-Origin")
  email.XFileName=msg.Header.Get("X-FileName")
  email.Subject=msg.Header.Get("Subject")

  body,errB:=ioutil.ReadAll(msg.Body)
  if errB!=nil{
    log.Panic("Error loading body")
  }else{
    email.Body=string(body)
  }

  return email
}



