package models

import "time"

type Log struct {
	ID        uint64    `json:"id,omitempty"`
	Descricao string    `json:"descricao,omitempty"`
	Usuario   uint64    `json:"usuario,omitempty"`
	Rota      string    `json:"rota,omitempty"`
	CriadoEm  time.Time `json:"criado-em,omitempty"`
}
