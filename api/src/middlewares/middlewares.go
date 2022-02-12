package middlewares

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"log"
	"net/http"
)

// Escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logSistema = models.Log{
			Descricao: r.Method,
			Rota:      r.RequestURI,
		}

		db, erro := database.Conectar()
		if erro != nil {
			responses.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		defer db.Close()

		repositorio := repository.NovoRepositorioDeLogs(db)
		logID, erro := repositorio.Log(logSistema)
		if erro != nil {
			responses.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		log.Printf("%d %s %s %s", logID, r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Verifica se o usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
