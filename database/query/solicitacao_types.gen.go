// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/eufelipemateus/test-example/models"
)

func newSolicitacaoType(db *gorm.DB, opts ...gen.DOOption) solicitacaoType {
	_solicitacaoType := solicitacaoType{}

	_solicitacaoType.solicitacaoTypeDo.UseDB(db, opts...)
	_solicitacaoType.solicitacaoTypeDo.UseModel(&models.SolicitacaoType{})

	tableName := _solicitacaoType.solicitacaoTypeDo.TableName()
	_solicitacaoType.ALL = field.NewAsterisk(tableName)
	_solicitacaoType.ID = field.NewUint(tableName, "id")
	_solicitacaoType.CreatedAt = field.NewTime(tableName, "created_at")
	_solicitacaoType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_solicitacaoType.Type = field.NewString(tableName, "type")
	_solicitacaoType.Enabled = field.NewBool(tableName, "enabled")
	_solicitacaoType.Solicitacoes = solicitacaoTypeHasManySolicitacoes{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Solicitacoes", "models.Solicitacao"),
		SolicitacaoType: struct {
			field.RelationField
			Solicitacoes struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Solicitacoes.SolicitacaoType", "models.SolicitacaoType"),
			Solicitacoes: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Solicitacoes.SolicitacaoType.Solicitacoes", "models.Solicitacao"),
			},
		},
	}

	_solicitacaoType.fillFieldMap()

	return _solicitacaoType
}

type solicitacaoType struct {
	solicitacaoTypeDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Type         field.String
	Enabled      field.Bool
	Solicitacoes solicitacaoTypeHasManySolicitacoes

	fieldMap map[string]field.Expr
}

func (s solicitacaoType) Table(newTableName string) *solicitacaoType {
	s.solicitacaoTypeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s solicitacaoType) As(alias string) *solicitacaoType {
	s.solicitacaoTypeDo.DO = *(s.solicitacaoTypeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *solicitacaoType) updateTableName(table string) *solicitacaoType {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.Type = field.NewString(table, "type")
	s.Enabled = field.NewBool(table, "enabled")

	s.fillFieldMap()

	return s
}

func (s *solicitacaoType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *solicitacaoType) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["type"] = s.Type
	s.fieldMap["enabled"] = s.Enabled

}

func (s solicitacaoType) clone(db *gorm.DB) solicitacaoType {
	s.solicitacaoTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s solicitacaoType) replaceDB(db *gorm.DB) solicitacaoType {
	s.solicitacaoTypeDo.ReplaceDB(db)
	return s
}

type solicitacaoTypeHasManySolicitacoes struct {
	db *gorm.DB

	field.RelationField

	SolicitacaoType struct {
		field.RelationField
		Solicitacoes struct {
			field.RelationField
		}
	}
}

func (a solicitacaoTypeHasManySolicitacoes) Where(conds ...field.Expr) *solicitacaoTypeHasManySolicitacoes {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a solicitacaoTypeHasManySolicitacoes) WithContext(ctx context.Context) *solicitacaoTypeHasManySolicitacoes {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a solicitacaoTypeHasManySolicitacoes) Session(session *gorm.Session) *solicitacaoTypeHasManySolicitacoes {
	a.db = a.db.Session(session)
	return &a
}

func (a solicitacaoTypeHasManySolicitacoes) Model(m *models.SolicitacaoType) *solicitacaoTypeHasManySolicitacoesTx {
	return &solicitacaoTypeHasManySolicitacoesTx{a.db.Model(m).Association(a.Name())}
}

type solicitacaoTypeHasManySolicitacoesTx struct{ tx *gorm.Association }

func (a solicitacaoTypeHasManySolicitacoesTx) Find() (result []*models.Solicitacao, err error) {
	return result, a.tx.Find(&result)
}

func (a solicitacaoTypeHasManySolicitacoesTx) Append(values ...*models.Solicitacao) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a solicitacaoTypeHasManySolicitacoesTx) Replace(values ...*models.Solicitacao) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a solicitacaoTypeHasManySolicitacoesTx) Delete(values ...*models.Solicitacao) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a solicitacaoTypeHasManySolicitacoesTx) Clear() error {
	return a.tx.Clear()
}

func (a solicitacaoTypeHasManySolicitacoesTx) Count() int64 {
	return a.tx.Count()
}

type solicitacaoTypeDo struct{ gen.DO }

type ISolicitacaoTypeDo interface {
	gen.SubQuery
	Debug() ISolicitacaoTypeDo
	WithContext(ctx context.Context) ISolicitacaoTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISolicitacaoTypeDo
	WriteDB() ISolicitacaoTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISolicitacaoTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISolicitacaoTypeDo
	Not(conds ...gen.Condition) ISolicitacaoTypeDo
	Or(conds ...gen.Condition) ISolicitacaoTypeDo
	Select(conds ...field.Expr) ISolicitacaoTypeDo
	Where(conds ...gen.Condition) ISolicitacaoTypeDo
	Order(conds ...field.Expr) ISolicitacaoTypeDo
	Distinct(cols ...field.Expr) ISolicitacaoTypeDo
	Omit(cols ...field.Expr) ISolicitacaoTypeDo
	Join(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo
	Group(cols ...field.Expr) ISolicitacaoTypeDo
	Having(conds ...gen.Condition) ISolicitacaoTypeDo
	Limit(limit int) ISolicitacaoTypeDo
	Offset(offset int) ISolicitacaoTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISolicitacaoTypeDo
	Unscoped() ISolicitacaoTypeDo
	Create(values ...*models.SolicitacaoType) error
	CreateInBatches(values []*models.SolicitacaoType, batchSize int) error
	Save(values ...*models.SolicitacaoType) error
	First() (*models.SolicitacaoType, error)
	Take() (*models.SolicitacaoType, error)
	Last() (*models.SolicitacaoType, error)
	Find() ([]*models.SolicitacaoType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SolicitacaoType, err error)
	FindInBatches(result *[]*models.SolicitacaoType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.SolicitacaoType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISolicitacaoTypeDo
	Assign(attrs ...field.AssignExpr) ISolicitacaoTypeDo
	Joins(fields ...field.RelationField) ISolicitacaoTypeDo
	Preload(fields ...field.RelationField) ISolicitacaoTypeDo
	FirstOrInit() (*models.SolicitacaoType, error)
	FirstOrCreate() (*models.SolicitacaoType, error)
	FindByPage(offset int, limit int) (result []*models.SolicitacaoType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISolicitacaoTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s solicitacaoTypeDo) Debug() ISolicitacaoTypeDo {
	return s.withDO(s.DO.Debug())
}

func (s solicitacaoTypeDo) WithContext(ctx context.Context) ISolicitacaoTypeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s solicitacaoTypeDo) ReadDB() ISolicitacaoTypeDo {
	return s.Clauses(dbresolver.Read)
}

func (s solicitacaoTypeDo) WriteDB() ISolicitacaoTypeDo {
	return s.Clauses(dbresolver.Write)
}

func (s solicitacaoTypeDo) Session(config *gorm.Session) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Session(config))
}

func (s solicitacaoTypeDo) Clauses(conds ...clause.Expression) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s solicitacaoTypeDo) Returning(value interface{}, columns ...string) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s solicitacaoTypeDo) Not(conds ...gen.Condition) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s solicitacaoTypeDo) Or(conds ...gen.Condition) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s solicitacaoTypeDo) Select(conds ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s solicitacaoTypeDo) Where(conds ...gen.Condition) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s solicitacaoTypeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISolicitacaoTypeDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s solicitacaoTypeDo) Order(conds ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s solicitacaoTypeDo) Distinct(cols ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s solicitacaoTypeDo) Omit(cols ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s solicitacaoTypeDo) Join(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s solicitacaoTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s solicitacaoTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s solicitacaoTypeDo) Group(cols ...field.Expr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s solicitacaoTypeDo) Having(conds ...gen.Condition) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s solicitacaoTypeDo) Limit(limit int) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s solicitacaoTypeDo) Offset(offset int) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s solicitacaoTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s solicitacaoTypeDo) Unscoped() ISolicitacaoTypeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s solicitacaoTypeDo) Create(values ...*models.SolicitacaoType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s solicitacaoTypeDo) CreateInBatches(values []*models.SolicitacaoType, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s solicitacaoTypeDo) Save(values ...*models.SolicitacaoType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s solicitacaoTypeDo) First() (*models.SolicitacaoType, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SolicitacaoType), nil
	}
}

func (s solicitacaoTypeDo) Take() (*models.SolicitacaoType, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SolicitacaoType), nil
	}
}

func (s solicitacaoTypeDo) Last() (*models.SolicitacaoType, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SolicitacaoType), nil
	}
}

func (s solicitacaoTypeDo) Find() ([]*models.SolicitacaoType, error) {
	result, err := s.DO.Find()
	return result.([]*models.SolicitacaoType), err
}

func (s solicitacaoTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SolicitacaoType, err error) {
	buf := make([]*models.SolicitacaoType, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s solicitacaoTypeDo) FindInBatches(result *[]*models.SolicitacaoType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s solicitacaoTypeDo) Attrs(attrs ...field.AssignExpr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s solicitacaoTypeDo) Assign(attrs ...field.AssignExpr) ISolicitacaoTypeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s solicitacaoTypeDo) Joins(fields ...field.RelationField) ISolicitacaoTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s solicitacaoTypeDo) Preload(fields ...field.RelationField) ISolicitacaoTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s solicitacaoTypeDo) FirstOrInit() (*models.SolicitacaoType, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SolicitacaoType), nil
	}
}

func (s solicitacaoTypeDo) FirstOrCreate() (*models.SolicitacaoType, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SolicitacaoType), nil
	}
}

func (s solicitacaoTypeDo) FindByPage(offset int, limit int) (result []*models.SolicitacaoType, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s solicitacaoTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s solicitacaoTypeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s solicitacaoTypeDo) Delete(models ...*models.SolicitacaoType) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *solicitacaoTypeDo) withDO(do gen.Dao) *solicitacaoTypeDo {
	s.DO = *do.(*gen.DO)
	return s
}
