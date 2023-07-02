package util

import (
	"fmt"
	"time"

	"github.com/beego/beego/orm"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type mysqlConnectionConfig struct {
	cfg        *mysql.Config
	Alias      string
	RetryCount int
	RetryDelay time.Duration
	TimeZone   *time.Location
	DebugFlag  bool
}

func ConnectToDatabase(keys ...string) error {
	fmt.Println("Connectin to DB")
	connString := "root:Sejal@1234@tcp(127.0.0.1:3306)/project3"
	fmt.Println("ConnString: ", connString)
	cfg := mysqlConnectionConfig{}
	err := cfg.ParseDSN(connString)
	if err != nil {
		return errors.Wrap(err, "failed to parse conn string into struct")
	}
	ISTLocation, _ := time.LoadLocation("Asia/Kolkata")
	cfg.TimeZone = ISTLocation
	cfg.Alias = "default"
	cfg.DebugFlag = true
	cfg.RetryCount = 5
	cfg.RetryDelay = 500 * time.Millisecond
	err = beegoRegisterDB(cfg)
	if err != nil {
		return err
	}
	return nil
}

func beegoRegisterDB(cfg mysqlConnectionConfig) error {
	var err error
	for breaker := cfg.RetryCount; breaker > 0; breaker-- {
		if breaker < cfg.RetryCount {
			time.Sleep(cfg.RetryDelay)
		}
		// check if database already exist
		_, err := orm.GetDB(cfg.Alias)
		if err != nil {
			err = orm.RegisterDataBase(cfg.Alias, "mysql", cfg.FormatDSN())
		}
		if err != nil {
			fmt.Println("failed to register db")
			continue
		}
		err = MysqlTest(cfg.Alias)
		if err != nil {
			fmt.Println("failed to query")
			continue
		}
		orm.DefaultTimeLoc = cfg.TimeZone
		orm.Debug = cfg.DebugFlag
		break
	}
	if err != nil {
		fmt.Println(fmt.Sprintf("beegoRegisterDB(alias: %s, dsn: %s) failed, error: %v", cfg.Alias, cfg.String(), err))
	}
	return err
}
func (c *mysqlConnectionConfig) ParseDSN(s string) error {
	var err error
	c.cfg, err = mysql.ParseDSN(s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func MysqlTest(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	a, err := o.Raw("SELECT 1").Exec()
	fmt.Println(a)
	return err
}
func (c *mysqlConnectionConfig) String() string {
	t := c.cfg.Clone()
	t.Passwd = "root"
	return t.FormatDSN()
}
func (c *mysqlConnectionConfig) FormatDSN() string {
	return c.cfg.FormatDSN()
}
