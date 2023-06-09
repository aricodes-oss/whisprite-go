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

	"whisprite/model"
)

func newAlias(db *gorm.DB, opts ...gen.DOOption) alias {
	_alias := alias{}

	_alias.aliasDo.UseDB(db, opts...)
	_alias.aliasDo.UseModel(&model.CommandAlias{})

	tableName := _alias.aliasDo.TableName()
	_alias.ALL = field.NewAsterisk(tableName)
	_alias.ID = field.NewUint(tableName, "id")
	_alias.CreatedAt = field.NewTime(tableName, "created_at")
	_alias.UpdatedAt = field.NewTime(tableName, "updated_at")
	_alias.DeletedAt = field.NewField(tableName, "deleted_at")
	_alias.Name = field.NewString(tableName, "name")
	_alias.Target = field.NewString(tableName, "target")

	_alias.fillFieldMap()

	return _alias
}

type alias struct {
	aliasDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Target    field.String

	fieldMap map[string]field.Expr
}

func (a alias) Table(newTableName string) *alias {
	a.aliasDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alias) As(alias string) *alias {
	a.aliasDo.DO = *(a.aliasDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alias) updateTableName(table string) *alias {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Name = field.NewString(table, "name")
	a.Target = field.NewString(table, "target")

	a.fillFieldMap()

	return a
}

func (a *alias) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alias) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["name"] = a.Name
	a.fieldMap["target"] = a.Target
}

func (a alias) clone(db *gorm.DB) alias {
	a.aliasDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alias) replaceDB(db *gorm.DB) alias {
	a.aliasDo.ReplaceDB(db)
	return a
}

type aliasDo struct{ gen.DO }

type IAliasDo interface {
	gen.SubQuery
	Debug() IAliasDo
	WithContext(ctx context.Context) IAliasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAliasDo
	WriteDB() IAliasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAliasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAliasDo
	Not(conds ...gen.Condition) IAliasDo
	Or(conds ...gen.Condition) IAliasDo
	Select(conds ...field.Expr) IAliasDo
	Where(conds ...gen.Condition) IAliasDo
	Order(conds ...field.Expr) IAliasDo
	Distinct(cols ...field.Expr) IAliasDo
	Omit(cols ...field.Expr) IAliasDo
	Join(table schema.Tabler, on ...field.Expr) IAliasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAliasDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAliasDo
	Group(cols ...field.Expr) IAliasDo
	Having(conds ...gen.Condition) IAliasDo
	Limit(limit int) IAliasDo
	Offset(offset int) IAliasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAliasDo
	Unscoped() IAliasDo
	Create(values ...*model.CommandAlias) error
	CreateInBatches(values []*model.CommandAlias, batchSize int) error
	Save(values ...*model.CommandAlias) error
	First() (*model.CommandAlias, error)
	Take() (*model.CommandAlias, error)
	Last() (*model.CommandAlias, error)
	Find() ([]*model.CommandAlias, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommandAlias, err error)
	FindInBatches(result *[]*model.CommandAlias, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CommandAlias) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAliasDo
	Assign(attrs ...field.AssignExpr) IAliasDo
	Joins(fields ...field.RelationField) IAliasDo
	Preload(fields ...field.RelationField) IAliasDo
	FirstOrInit() (*model.CommandAlias, error)
	FirstOrCreate() (*model.CommandAlias, error)
	FindByPage(offset int, limit int) (result []*model.CommandAlias, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAliasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aliasDo) Debug() IAliasDo {
	return a.withDO(a.DO.Debug())
}

func (a aliasDo) WithContext(ctx context.Context) IAliasDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aliasDo) ReadDB() IAliasDo {
	return a.Clauses(dbresolver.Read)
}

func (a aliasDo) WriteDB() IAliasDo {
	return a.Clauses(dbresolver.Write)
}

func (a aliasDo) Session(config *gorm.Session) IAliasDo {
	return a.withDO(a.DO.Session(config))
}

func (a aliasDo) Clauses(conds ...clause.Expression) IAliasDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aliasDo) Returning(value interface{}, columns ...string) IAliasDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aliasDo) Not(conds ...gen.Condition) IAliasDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aliasDo) Or(conds ...gen.Condition) IAliasDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aliasDo) Select(conds ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aliasDo) Where(conds ...gen.Condition) IAliasDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aliasDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAliasDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a aliasDo) Order(conds ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aliasDo) Distinct(cols ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aliasDo) Omit(cols ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aliasDo) Join(table schema.Tabler, on ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aliasDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAliasDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aliasDo) RightJoin(table schema.Tabler, on ...field.Expr) IAliasDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aliasDo) Group(cols ...field.Expr) IAliasDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aliasDo) Having(conds ...gen.Condition) IAliasDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aliasDo) Limit(limit int) IAliasDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aliasDo) Offset(offset int) IAliasDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aliasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAliasDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aliasDo) Unscoped() IAliasDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aliasDo) Create(values ...*model.CommandAlias) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aliasDo) CreateInBatches(values []*model.CommandAlias, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aliasDo) Save(values ...*model.CommandAlias) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aliasDo) First() (*model.CommandAlias, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandAlias), nil
	}
}

func (a aliasDo) Take() (*model.CommandAlias, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandAlias), nil
	}
}

func (a aliasDo) Last() (*model.CommandAlias, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandAlias), nil
	}
}

func (a aliasDo) Find() ([]*model.CommandAlias, error) {
	result, err := a.DO.Find()
	return result.([]*model.CommandAlias), err
}

func (a aliasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommandAlias, err error) {
	buf := make([]*model.CommandAlias, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aliasDo) FindInBatches(result *[]*model.CommandAlias, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aliasDo) Attrs(attrs ...field.AssignExpr) IAliasDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aliasDo) Assign(attrs ...field.AssignExpr) IAliasDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aliasDo) Joins(fields ...field.RelationField) IAliasDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aliasDo) Preload(fields ...field.RelationField) IAliasDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aliasDo) FirstOrInit() (*model.CommandAlias, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandAlias), nil
	}
}

func (a aliasDo) FirstOrCreate() (*model.CommandAlias, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandAlias), nil
	}
}

func (a aliasDo) FindByPage(offset int, limit int) (result []*model.CommandAlias, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aliasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aliasDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aliasDo) Delete(models ...*model.CommandAlias) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aliasDo) withDO(do gen.Dao) *aliasDo {
	a.DO = *do.(*gen.DO)
	return a
}
