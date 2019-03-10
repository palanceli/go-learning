package goldb

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql" // 如果不import改行，连接数据库时会提示找不到driverName
	"github.com/golang/glog"
)

// 参考资料
// http://www.xorm.io/docs/
// https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/55996 # 两个内容一样，只是这个版本支持搜索

func TestMain(m *testing.M) {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}
func TestSQLDropDatabase(t *testing.T) {
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/?charset=utf8"

	db, err := sql.Open("mysql", dbSourceName)
	if err != nil {
		glog.Fatalf("Failed to connect [%s], err= %v ", dbSourceName, err)
	}
	defer db.Close()

	dbName := "test_db"
	sqlString := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	glog.Infof("exec [%s]...", sqlString)
	_, err = db.Query(sqlString)
	if err != nil {
		glog.Fatalf("Failted to exec [%s], err = %v", sqlString, err)
	}
}
func TestSQLCreateDatabase(t *testing.T) {
	// 参考https://studygolang.com/articles/5402
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/?charset=utf8"

	db, err := sql.Open("mysql", dbSourceName)
	if err != nil {
		glog.Fatalf("Failed to connect [%s], err= %v ", dbSourceName, err)
	}
	defer db.Close()

	dbName := "test_db"
	sqlString := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	glog.Infof("exec [%s]...", sqlString)
	_, err = db.Query(sqlString)
	if err != nil {
		glog.Fatalf("Failted to exec [%s], err = %v", sqlString, err)
	}

	sqlString = fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8 COLLATE utf8_general_ci", dbName)
	glog.Infof("exec [%s]...", sqlString)
	_, err = db.Query(sqlString)
	if err != nil {
		glog.Fatalf("Failted to exec [%s], err = %v", sqlString, err)
	}
}

type WindCoinPrice struct {
	ID      int       `xorm:"'id' pk autoincr"`
	Type    string    `xorm:"coin_type" json:"coin_type"`
	Buy     float64   `xorm:"buy" json:"buy"`
	Sell    float64   `xorm:"sell" json:"sell"`
	Middle  float64   `xorm:"middle" json:"middle"`
	Weight  float64   `xorm:"weight" json:"weight"`
	Charge  float64   `xorm:"charge" json:"charge"`
	Update  int64     `xorm:"update_time index" json:"update_time"`
	Ts      time.Time `xorm:"ts" json:"ts"`
	Usdtcny float64   `xorm:"usdtcny" json:"usdtcny"`
	// TimeGapMs float64   `xorm:"usdtcny" json:"time_gap_ms"`
}

// 使用Xorm建表
func TestXormCreateTable(t *testing.T) {
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dbSourceName)
	glog.Infof("Connecting [%s]", dbSourceName)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true) // 显示SQL的执行，便与调试
	defer engine.Close()
	// 创建表
	err = engine.Sync2(new(WindCoinPrice))
	if err != nil {
		panic(err)
	}
}

// 使用Xorm获取数据库信息
func TestXormDBInfo(t *testing.T) {
	// 参考https://github.com/go-xorm/xorm/blob/master/engine.go
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dbSourceName)
	glog.Infof("Connecting [%s]", dbSourceName)
	if err != nil {
		panic(err)
	}
	// engine.ShowSQL(true) // 显示SQL的执行，便与调试
	defer engine.Close()
	tableInfo, err := engine.DBMetas()

	for _, table := range tableInfo {
		colSeq, cols, err := engine.Dialect().GetColumns(table.Name)
		if err != nil {
			panic(err)
		}
		// glog.Infof("%v", cols)
		glog.Infof("表名：%s", table.Name)
		for _, name := range colSeq {
			glog.Infof("%-12s %v", name, cols[name])
		}
	}
}

// 使用Xorm添加数据
func TestXormInsert(t *testing.T) {
	// 参考https://blog.csdn.net/stpeace/article/details/83114319
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dbSourceName)
	glog.Infof("Connecting [%s]", dbSourceName)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true) // 显示SQL的执行，便与调试
	defer engine.Close()
	item := new(WindCoinPrice)
	item.Buy = 286.12010609442405
	item.Charge = 0.02
	item.Middle = 280.50990793570986
	item.Sell = 280.50990793570986
	item.Ts = time.Now()
	item.Type = "eth"
	item.Update = 1535105050
	item.Weight = 0.02
	affected, err := engine.Insert(item)
	if err != nil {
		panic(err)
	}
	glog.Infoln(affected)
}

// 使用Xorm查询数据
func TestXormQuery(t *testing.T) {
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dbSourceName)
	glog.Infof("Connecting [%s]", dbSourceName)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true) // 显示SQL的执行，便与调试
	defer engine.Close()
	item := new(WindCoinPrice)
	result, err := engine.Get(item)
	if err != nil {
		panic(err)
	}
	glog.Infoln(result)
	glog.Infof("%v", item)
}

// 从admin.db下拷贝所有wind_xxx_coin_price表格中的一条记录到本地
func TestWindImgSyncOneItem(t *testing.T) {
	// 发现我连不上admin.db，不知道连接串该怎么写，暂时放弃
	srcDBSourceName := "admin.db/sql.php?db=pricing_engine&table=wind_eth_coin_price&pos=0"
	glog.Infoln(srcDBSourceName)
}

// 在本地pricing_engine/wind_xxx_coin_price表中创建假数据
func TestCreateFakedata(t *testing.T) {
	dbSourceName := "root:nice@tcp(127.0.0.1:3306)/pricing_engine?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dbSourceName)
	glog.Infof("Connecting [%s]", dbSourceName)
	if err != nil {
		panic(err)
	}
	// engine.ShowSQL(true) // 显示SQL的执行，便与调试
	defer engine.Close()
	tableInfo, err := engine.DBMetas()

	for _, table := range tableInfo {
		if table.Name[:5] == "wind_" &&
			table.Name[len(table.Name)-len("_coin_price"):] == "_coin_price" {
			coinType := table.Name[len("wind_") : len(table.Name)-len("_coin_price")]
			glog.Infoln(coinType)
			sqlString := fmt.Sprintf("INSERT INTO `%s` (`coin_type`,`buy`,`sell`,`middle`,`weight`,`charge`,`update_time`,`ts`,`usdtcny`) VALUES (?, ?, ?, ?, ?, ?, ?, \"2019-03-11 00:00:00\", ?)", table.Name)
			glog.Infof("Exec [%s]", sqlString)
			res, err := engine.Exec(sqlString, coinType, 286.12010609442405, 280.50990793570986, 280.50990793570986, 0.02, 0.02, 1535105050, 0)
			if err != nil {
				panic(err)
			}
			glog.Infoln(res)
		}
	}

}
