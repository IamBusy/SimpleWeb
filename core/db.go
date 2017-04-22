package core

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

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
	SetBuilder(builder Builder)
	GetBuilder() Builder
	//
	//Table(table string)
	//Select(columns string)
	//
	//Where(column,operation,value string)
	//Join(table, column1, operation, column2 string)
	//Skip(offset int)
	//Take(limit int)

	Raw(sqlQuery string) (*sql.Rows,error)
	//Get()
	//Insert()
	//Update()
	//Delete()
}

type Builder interface {
	parse() string
	//clear()

	Table(table string)
	//Select(columns string)
	//
	//Where(column,operation,value string)
	//Join(table, column1, operation, column2 string)
	//Skip(offset int)
	//Take(limit int)

	Raw(sqlQuery string)
	//Get()
	//Insert()
	//Update()
	//Delete()

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


func (db *manager) SetBuilder(builder Builder)  {
	db.builder = builder
}

func (db *manager) GetBuilder() Builder  {
	return db.builder
}

func (db *manager) Raw(query string) (*sql.Rows,error) {
	return db.db.Query(query)
}


type builder struct {
	table string
	param map[string][][]string
	crud string
	limit int
	offset int
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
	//for whereItems,exist := range (builder.param["where"]) {
	//	if exist {
	//		res = res + "AND " + whereItems[0] +
	//			" " + whereItems[1] + whereItems[2]
	//	}
	//}
	if res != "" {
		res = " WHERE " + res[3:]
	}
	return res
}


func (builder *builder) Table(table string) {
	builder.table = table
}

func (builder *builder) Raw(query string)  {

}

