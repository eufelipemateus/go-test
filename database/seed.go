package database

import (
	"github.com/eufelipemateus/test-example/database/query"
	"github.com/eufelipemateus/test-example/models"
	"gorm.io/gorm/clause"
)

/**
 * Semear o Banco de dados
 */

func seed() {

	var list []*models.SolicitacaoType

	list = append(list,
		&models.SolicitacaoType{
			Type: "Suporte",
		})

	list = append(list,
		&models.SolicitacaoType{
			Type: "Sugestão",
		})

	query.Q.SolicitacaoType.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(list, len(list))
}
