package mailreader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

// set email fields with msg data 
func (email *Email) mailSetter(msg *mail.Message) error{

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
    log.Println("Error loading body")
    return errB
  }else{
    email.Body=string(body)
  }

  return nil
}


//read email file from path and set in Email structure
func Reader(path string) (string,error) {

  content,errIO:=ioutil.ReadFile(path)
  if errIO!= nil{
    log.Println("Error: loading the file ",path)
    return "",errIO
  }
  fileContent:=string(content)

  msg,errE:=mail.ReadMessage(strings.NewReader(fileContent))
  if errE != nil{
    log.Println("Error: loading email ")
    return "",errE
  }

  var email Email
  errS:=email.mailSetter(msg)
  if errS != nil{
    log.Println("Error: setting email structure ")
    return "",errS
  }

  jsonEmail,errJ:=json.Marshal(email)
  if errJ!=nil{
    log.Println("Error: generating struct to json ",path)
    return "",errJ
  }

  return string(jsonEmail),nil
}



