package database

import (
	"github.com/eufelipemateus/test-example/models"
	"gorm.io/gen"
)

/**
 * Gerar Banco de dados
 */
func GenerateDB() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(DB)

	g.ApplyBasic(
		models.Solicitacao{},
		models.SolicitacaoType{},
	)

	g.Execute()

	DB.AutoMigrate(
		models.Solicitacao{},
		models.SolicitacaoType{},
	)

	seed()
}
