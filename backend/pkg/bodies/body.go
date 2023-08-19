package bodies

type BodyQuery struct{
  Search      string  `json:"Search"`
  MessageId   string  `json:"MessageId"`
  Date        string  `json:"Date"`
  From        string  `json:"From"`
  To          string  `json:"To"`  
  CC          string  `json:"CC"`
  Subject     string  `json:"Subject"`
  Body        string  `json:"Body"`
}
