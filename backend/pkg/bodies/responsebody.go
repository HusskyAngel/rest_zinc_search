package bodies

import "github.com/HusskyAngel/BackendEmailSearcher/pkg/email"

type ResponseBody struct{
  NumberResults int  `json:"number_results"`
  Results [] struct{Mail email.Email `json:"_source"`} `json:"results"`
   
}
