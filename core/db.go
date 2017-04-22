package core

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	DB_MYSQL string = "mysql"
)

type Connector interface {
	Connect(dbType ,host , port , user, password, db string)
	GetDB() *sql.DB
}

type Manager interface {
	SetConnector(connector Connector)
	GetConnector() Connector
	GetBuilder() Builder
}

type Builder interface {

	setDB(*sql.DB)

	Table(table string)Builder
	Select(columns string)Builder
	Where(column,operation,value string)Builder
	Join(table, column1, operation, column2 string) Builder
	Skip(offset int)Builder
	Take(limit int)Builder

	Raw(sqlQuery string)(*sql.Rows,error)
	Get()(*sql.Rows,error)
	Insert([][2]string)(sql.Result,error)
	Update([][2]string)(sql.Result,error)
	Delete()(sql.Result,error)

}

type connector struct {
	db *sql.DB
}

func NewConnector() Connector  {
	return &connector{}
}

func (connector *connector) Connect(dbType ,host , port , user, password, db string)  {
	dbConnector, err := sql.Open(dbType, user+":"+password+"@tcp("+host+":"+port+")/"+db+"?charset=utf8")
	if err == nil {
		connector.db = dbConnector
	} else {
		err.Error()
	}
}

func (connector *connector) GetDB() *sql.DB {
	return connector.db
}


/**
 * Manager
 */
type manager struct {
	db *sql.DB
	connector Connector
	builder Builder
}

func NewManager() Manager {
	mng := &manager{}
	mng.builder = &builder{}
	return mng
}


func (db *manager) SetConnector(connector Connector)  {
	db.connector = connector
	db.db = connector.GetDB()
}

func (db *manager) GetConnector() Connector {
	return db.connector
}


func (db *manager) GetBuilder() Builder  {
	builder := &builder{
		columns: "*",
		limit: -1,
		offset: -1,
	}
	builder.setDB(db.db)
	builder.param = make(map[string][][]string)
	return builder
}

func (db *manager) Raw(query string) (*sql.Rows,error) {
	return db.db.Query(query)
}


type builder struct {
	db *sql.DB
	table string
	param map[string][][]string
	crud string
	limit int
	offset int
	columns string
}

func (builder *builder) setDB(db *sql.DB)  {
	builder.db = db
}

func (builder *builder) parse() string {
	//whereSql := ""
	switch builder.crud {
	case "SELECT":
	case "UPDAT":
	case "INSERT":
	case "DELETE":
		
	}

	return ""
}

func (builder *builder) parseWhere() string  {
	res := ""
	for i:=0; i< len(builder.param["where"]); i++ {
		whereItems :=  builder.param["where"][i]
		res = res + "AND " + whereItems[0] +
			" " + whereItems[1] + whereItems[2]
	}
	if res != "" {
		res = " WHERE " + res[3:]
	}
	return res
}


func (builder *builder) Table(table string) Builder {
	builder.table = table
	return builder
}

func (builder *builder) Where(column,operation,value string) Builder  {
	builder.param["where"]= append(builder.param["where"],[]string{
		column,
		operation,
		value,
	})
	return  builder
}

func (builder *builder) Join(table, column1, operation, column2 string) Builder  {
	builder.param["join"] = append(builder.param["join"],[]string{
		table,
		column1,
		operation,
		column2,
	})
	return  builder
}

func (builder *builder) Select(columns string) Builder  {
	builder.columns = columns
	return  builder
}

func (builder *builder) Skip(offet int) Builder  {
	builder.offset = offet
	return builder
}

func (builder *builder) Take(limit int) Builder  {
	builder.limit = limit
	return builder
}

func (builder *builder) Get() (*sql.Rows,error) {
	//TODO add support for join
	sqlQuery := "SELECT " + builder.columns + " FROM " +builder.table + builder.parseWhere()
	if builder.limit == -1 && builder.offset != -1{
		sqlQuery += " LIMIT " + string(builder.offset) + ",-1"
	} else if builder.limit != -1 && builder.offset == -1 {
		sqlQuery += " LIMIT " + string(builder.offset)
	} else if builder.limit != -1 && builder.offset != -1 {
		sqlQuery += " LIMIT " +string(builder.offset) + "," +string(builder.limit)
	}

	return builder.db.Query(sqlQuery)
}

func (builder *builder) Insert(data [][2]string) (sql.Result,error) {
	columns := ""
	values := ""
	for i:=0 ;i<len(data) ;i++  {
		columns += ", "+data[i][0]
		values += ", "+data[i][1]
	}

	if columns != "" {
		columns = " ( " + columns[1:] + " ) "
		values = "(" + values[1:] + ")"
	}
	sqlQuery := "INSERT INTO " + builder.table + columns + " VALUES " +values
	return builder.db.Exec(sqlQuery)
}

func (builder *builder) Update(data [][2]string) (sql.Result,error) {
	setSentence := builder.parseSet(data)
	whereSentence := builder.parseWhere()

	sqlQuery := "UPDATE " + builder.table +setSentence + whereSentence
	return builder.db.Exec(sqlQuery)
}

func (builder *builder) Delete() (sql.Result,error) {
	whereSentence := builder.parseWhere()
	sqlQuery := "DELETE from "+ builder.table + whereSentence
	return builder.db.Exec(sqlQuery)
}



func (builder *builder) parseSet(data [][2]string) string {
	setSentence := ""
	for i:=0; i<len(data);i++ {
		setSentence += ", " + data[i][0] + " = " +data[i][1]
	}
	if setSentence != "" {
		setSentence = " SET " + setSentence[1:]
	}
	return setSentence
}

func (builder *builder) parseJoin()  {
	//TODO
}



func (builder *builder) Raw(query string) (*sql.Rows,error) {
	query = strings.TrimSpace(query)
	if strings.ToUpper(query[0:6]) == "SELECT" {
		return builder.db.Query(query)
	} else {
		_,err := builder.db.Exec(query)
		if err == nil {
			return nil,nil
		}
		return nil,err
	}
}

