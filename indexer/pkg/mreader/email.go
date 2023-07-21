package mreader

type Email struct {
  messageId               string `json:"Message-Id"`
  date                    string `json:"Date"`
	from                    string `json:"From"`
  to                      string `json:"To"`
  cc                      string `json:"CC"`
  subject                 string `json:"Subject"` 
  mimeVersion             string `json:"Mime-Version"`  
  contentType             string `json:"Content-Type"`
  contentTransferEncoding string `json:"Content-Transfer-Encoding"`
  xFrom                   string `json:"X-From"`
  xTo                     string `json:"X-To"`
  xcc                     string `json:"X-cc"`
  xbcc                    string `json:"X-bcc"`
  xFolder                 string `json:"X-Folder"`
  xOrigin                 string `json:"X-Origin"`
  xFileName               string `json:"X-FileName"`
  body                    string `json:"Body"`
}
