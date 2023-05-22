package models

/*
 * Typo de solicitação
 */
type SolicitacaoType struct {
	Model
	Type         string        `json:"type"`
	Enabled      bool          `json:"enabled" gorm:"default: true;"`
	Solicitacoes []Solicitacao `json:"solicitacoes" gorm:"foreignKey:ID;" `
}


/*
 * Solcitiação tipo select
 */
type ApiSolicitacaoType struct {
	ID   uint   `json:"id"`
	Type string `json:"string"`
}

func (c *SolicitacaoType) ApiSolicitacaoType() *ApiSolicitacaoType {
	return &ApiSolicitacaoType{
		ID:   c.ID,
		Type: c.Type,
	}
}
