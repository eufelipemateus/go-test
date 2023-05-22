package controllers

import (
	"context"
	"fmt"
	"net/http"

	config "github.com/eufelipemateus/test-example/configs"
	"github.com/eufelipemateus/test-example/database/query"
	"github.com/eufelipemateus/test-example/models"
	"github.com/eufelipemateus/test-example/utils"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type solicitacaoType struct {
	Name              string `form:"name"  json:"name"  validate:"required" `
	Email             string `form:"email" json:"email"  validate:"required,email" `
	Telefone          string `form:"telefone" json:"telefone"  validate:"required" `
	SolicitacaoTypeID uint   `form:"solicitacao"  json:"solicitacao"  validate:"required"`
	Produto           string `form:"produto" json:"produto"`
	Descricao         string `form:"descricao" json:"descricao"`
}

/**	
 * Retornar tipo de solicitação do DB
 */
func getSolicitacaoType(id uint) string {
	ctx := context.Background()
	solicitacao := query.SolicitacaoType
	data, _ := solicitacao.WithContext(ctx).Where(solicitacao.ID.Eq(id)).First()
	return data.Type
}

/**	
 * Email em formato de texto
 */
func parseText(solicitacao solicitacaoType) string {
	texto := "Nome:  %s\r\n  Email: %s\r\n  Telefone: %s\r\n Solicitação: %s %s\r\n Descricação: %s \r\n"
	return fmt.Sprintf(texto, solicitacao.Name, solicitacao.Email, solicitacao.Telefone, getSolicitacaoType(solicitacao.SolicitacaoTypeID), solicitacao.Produto, solicitacao.Descricao)
}

/**	
 * Email em formato de html
 */
func parseTextHTML(solicitacao solicitacaoType) string {
	texto := "<p><strong>Nome:</strong>  %s</p></br>\r\n  <p><strong>Email:</strong> %s</p></br>\r\n  <p><strong>Telefone:</strong> %s</p></br>\r\n  <p><strong>Solicitação:</strong> %s %s</p></br>\r\n  <p><strong>Descricação:</strong> %s<p></br>\r\n"
	return fmt.Sprintf(texto, solicitacao.Name, solicitacao.Email, solicitacao.Telefone, getSolicitacaoType(solicitacao.SolicitacaoTypeID), solicitacao.Produto, solicitacao.Descricao)
}

/**	
 * Registrar solicitação de suporte
 */
func registerSolicitacao(st *solicitacaoType) {
	ctx := context.Background()

	s := query.Solicitacao

	sol := models.Solicitacao{
		Name: st.Name,
		Email: st.Email,
		Produto: st.Produto,
		Descricao: st.Descricao,
	}

	s.WithContext(ctx).Create(&sol)
}


func SsendMail(c *gin.Context) {

	solicitacao := new(solicitacaoType)

	if err := c.ShouldBind(&solicitacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messag1": err.Error()})
		return
	}

	errors := utils.ValidateStruct(*solicitacao)
	if errors != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"message2": errors[0].Value})

		c.HTML(http.StatusOK, "contact.html", gin.H{
			"message": errors[0].Value,
		})
		return
	}

	from := mail.NewEmail("User", config.GetEmail().From)
	subject := "Novo Contato recebido"
	to := mail.NewEmail("User", config.GetEmail().To)
	plainTextContent := parseText(*solicitacao)
	htmlContent := parseTextHTML(*solicitacao)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.GetSengridApiKey())
	_, err := client.Send(message)

	if err != nil {
		println(err)
		//log.Error(err.Error())
		return
	} else {
		/*c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Email enviado com Sucesso!",
		})*/

		if solicitacao.SolicitacaoTypeID == 1 {
			registerSolicitacao(solicitacao)
		}

		solicitacao = &solicitacaoType{}

		c.HTML(http.StatusOK, "contact.html", gin.H{
			"message":           "Email enviado com sucesso!",
			"solicitacoesTypes": getTypes(),
			"solicitacao":       solicitacao,
		})
		return
	}

}
