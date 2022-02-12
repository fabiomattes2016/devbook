package repository

import (
	"api/src/models"
	"database/sql"
)

type logs struct {
	db *sql.DB
}

// Cria um repositorio de usuários
func NovoRepositorioDeLogs(db *sql.DB) *logs {
	return &logs{db}
}

// Registra um log
func (repositorio logs) Log(log models.Log) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO logs (descricao, rota) VALUES (?,?)",
	)
	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	resultado, erro := statement.Exec(log.Descricao, log.Rota)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}
