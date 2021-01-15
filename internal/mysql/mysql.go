package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type GMysqlClient struct {
	Engine *xorm.Engine
}

type GMysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string

	ConnMin         int
	ConnMax         int
	ConnMaxLifetime time.Duration
}

func New(config *GMysqlConfig) (*GMysqlClient, error) {
	db, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.ConnMin)
	db.SetMaxOpenConns(config.ConnMax)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMapper(core.SnakeMapper{})

	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// db.SetDefaultCacher(cacher)

	return &GMysqlClient{
		Engine: db,
	}, nil
}

type GMysqlStats struct {
	DDStats sql.DBStats
}

func (r *GMysqlClient) Status() *GMysqlStats {
	return &GMysqlStats{
		DDStats: r.Engine.DB().Stats(),
	}
}

type GMysqlModel struct {
	conn  *GMysqlClient
	table string
}

func NewGMysqlModel(conn *GMysqlClient, table string) *GMysqlModel {
	return &GMysqlModel{
		conn:  conn,
		table: table,
	}
}

func (m *GMysqlModel) Insert(data interface{}) (int64, error) {
	return m.conn.Engine.Table(m.table).Insert(data)
}

func (m *GMysqlModel) InsertOne(data interface{}) (int64, error) {
	return m.conn.Engine.Table(m.table).InsertOne(data)
}

func (m *GMysqlModel) InsertMulti(rowsSlicePtr interface{}) (int64, error) {
	return m.conn.Engine.Table(m.table).InsertMulti(rowsSlicePtr)
}

func (m *GMysqlModel) Update(cond, fields map[string]interface{}) (int64, error) {
	affected, err := m.conn.Engine.Table(m.table).Update(fields, cond)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (m *GMysqlModel) Get(cond map[string]interface{}, result interface{}) (err error) {
	var has bool
	has, err = m.conn.Engine.Table(m.table).Where(cond).Get(result)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}

	return nil
}

func (m *GMysqlModel) Count(cond map[string]interface{}) (count int64, err error) {
	if len(cond) > 0 {
		return m.conn.Engine.Table(m.table).Where(cond).Count()
	} else {
		return m.conn.Engine.Table(m.table).Count()
	}
}

func (m *GMysqlModel) Find(record interface{}, cond ...interface{}) error {
	return m.conn.Engine.Table(m.table).Find(record, cond)
}

func (m *GMysqlModel) FindAndCount(record interface{}, cond ...interface{}) (int64, error) {
	return m.conn.Engine.Table(m.table).FindAndCount(record, cond)
}
