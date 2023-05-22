package models

/*
 * Solicitações de suporte
 */
type Solicitacao struct {
	Model

	SolcicitaoTypeID uint            `json:"solicitacao_type_id" gorm:"not null"`
	SolicitacaoType  SolicitacaoType `json:"solicitacao_type" gorm:"foreignKey:SolcicitaoTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name             string          `json:"name"`
	Email            string          `json:"email"`
	Produto          string          `json:"produto"`
	Descricao        string          `json:"descricao"`
}
