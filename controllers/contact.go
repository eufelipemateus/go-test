package controllers

import (
	"context"
	"net/http"

	"github.com/eufelipemateus/test-example/database/query"
	"github.com/eufelipemateus/test-example/models"
	"github.com/gin-gonic/gin"
)


/**
 * Retornar tipos de solicitaçẽos 	 
 **/
func getTypes() []*models.ApiSolicitacaoType {
	ctx := context.Background()

	solicitacao := query.SolicitacaoType

	data, _ := solicitacao.WithContext(ctx).Where(solicitacao.Enabled.Is(true)).Find()

	var listTypes []*models.ApiSolicitacaoType

	for i := range data {
		listTypes = append(listTypes, data[i].ApiSolicitacaoType())
	}

	return listTypes
}

/**
 * Exibir  pagina de contato
 */
func GetContact(c *gin.Context) {

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"solicitacoesTypes": getTypes(),
	})

}
