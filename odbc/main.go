package main

import (
	"log"
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"database/sql"
	"database/sql/driver"
	_ "github.com/alexbrainman/odbc"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/charmap"
	"flag"
)

var (
	server = flag.String("S", "127.0.0.1", "ASE Server address")
	port = flag.Int("P", 5000, "ASE Server port")
	database = flag.String("d", "test", "Database name")
	uid = flag.String("u", "sa", "User name")
	password = flag.String("p", "", "Password")
	query = flag.String("q", "select Name from Users where IdUser = 7", "Sample SQL")
)

func main() {

	var err error

	flag.Parse()

	dsn := fmt.Sprintf("DRIVER=FreeTDS;" +
		"SERVER=%s;" +
		"PORT=%d;" +
		"UID=%s;" +
		"PWD=%s;" +
		"DATABASE=%s;" +
		// "Charset=UTF-8;" +
		"TDS_Version=5.0",
			*server, *port,
			*uid, *password,
			*database)
	log.Print(dsn)
	db, err := sql.Open("odbc", dsn)
	check(err)
	defer db.Close()


	// Get my sysprocess info
	sysprocess, err := selectSysprocesses(db)
	check(err)
	fmt.Println(sysprocess)

	// Execute sample query
	// row := db.QueryRow("select Name from Users where IdUser = ?", 7) TODO not support numeric type
	row := db.QueryRow(*query)
	var NameInCP1251 string
	err = row.Scan(&NameInCP1251)
	check(err)

	// Decode Name from CP1251 to UTF8
	NameUtf8, err := ioutil.ReadAll(
		transform.NewReader(
			strings.NewReader(NameInCP1251), charmap.Windows1251.NewDecoder()))
	check(err)

	fmt.Println(string(NameUtf8))

}

type SysProcess struct {
	spid int16
	kpid int
	enginenum int
	status string
	suid int
	hostname string
	program_name string
	hostprocess string
	cmd string
	cpu int
	physical_io int
	memusage int
	blocked int16
	dbid int16
	uid int
	gid int
	tran_name sql.NullString
	time_blocked sql.NullInt64
	network_pktsz sql.NullInt64
	fid sql.NullInt64 // smallint
	execlass sql.NullString
	priority sql.NullString
	affinity sql.NullString
	id sql.NullInt64
	stmtnum sql.NullInt64
	linenum sql.NullInt64
	origsuid sql.NullInt64
	block_xloid sql.NullInt64
	clientname sql.NullString
	clienthostname sql.NullString
	clientapplname sql.NullString
	sys_id sql.NullInt64
	ses_id sql.NullInt64
	loggedindatetime NullTime // Datetime
	ipaddr sql.NullString
}

type NullTime struct {
	Time time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt *NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func selectSysprocesses(db *sql.DB) (*SysProcess, error) {
	row := db.QueryRow("select * from  master..sysprocesses sp where sp.spid = @@spid")
	var r SysProcess
	err := row.Scan(
		&r.spid,
		&r.kpid,
		&r.enginenum,
		&r.status,
		&r.suid,
		&r.hostname,
		&r.program_name,
		&r.hostprocess,
		&r.cmd,
		&r.cpu,
		&r.physical_io,
		&r.memusage,
		&r.blocked,
		&r.dbid,
		&r.uid,
		&r.gid,
		&r.tran_name,
		&r.time_blocked,
		&r.network_pktsz,
		&r.fid,
		&r.execlass,
		&r.priority,
		&r.affinity,
		&r.id,
		&r.stmtnum,
		&r.linenum,
		&r.origsuid,
		&r.block_xloid,
		&r.clientname,
		&r.clienthostname,
		&r.clientapplname,
		&r.sys_id,
		&r.ses_id,
		&r.loggedindatetime,
		&r.ipaddr)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*

sudo apt-get install unixodbc unixodbc-dev freetds-dev tdsodbc

; /etc/odbcinst.ini
[FreeTDS]
Description=FreeTDS
Driver=/usr/lib/x86_64-linux-gnu/odbc/libtdsodbc.so
Driver64=
Setup=/usr/lib/x86_64-linux-gnu/odbc/libtdsS.so
Setup64=
UsageCount=
CPTimeout=
CPTimeToLive=
DisableGetFunctions=
DontDLCLose=
ExFetchMapping=
Threading=
FakeUnicode=yes
IconvEncoding=cp1251
Trace=
TraceFile=
TraceLibrary=
FileUsage=1
CPResuse=

 */