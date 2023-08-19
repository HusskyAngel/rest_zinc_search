package bodies

import "github.com/HusskyAngel/BackendEmailSearcher/pkg/email"

type ZSBody struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct{Mail email.Email `json:"_source"`} `json:"hits"`
	}`json:"hits"`
}
