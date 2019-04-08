package main

import (
	"net/http"
	"os"

	"github.com/ratemyclass/rmc-web/logging"
	"github.com/ratemyclass/rmc-web/rmcserver"

	"github.com/ratemyclass/rmc-web/rmcdb/rmcpostgres"
)

var (
	defaultRMCHomeDir     = os.Getenv("HOME") + "/.ratemyclass/"
	defaultConfigFilename = "rmc.conf"
	defaultLogFilename    = "ratemyclass.log"
	defaultHTTPPort       = uint16(8080)
	defaultHTTPHost       = "localhost"
	defaultCertLocation   = defaultRMCHomeDir + "certificate.pem"
	defaultDBHost         = "localhost"
	defaultDBPort         = uint16(5432)
	defaultDBUsername     = "postgres"
	defaultDBPassword     = ""
	defaultDBName         = "postgres"
	defaultLogLevel       = 0
)

type rmcConfig struct {
	// Home directory, Config options, Log file options
	RMCHomeDir     string `short:"d" long:"dir" description:"Specify home directory fo RateMyClass as an absolute path"`
	ConfigFile     string
	ConfigFileName string `long:"configfilename" description:"Specify config filename within RateMyClass home directory"`
	LogFileName    string `long:"logfilename" description:"Specify log filename within RateMyClass home directory"`

	// HTTP options
	HTTPPort uint16 `short:"p" long:"httpport" description:"Set HTTP port to connect to"`
	HTTPHost string `short:"h" long:"httphost" description:"Set HTTP host to listen on"`

	// SSL options
	CertLocation string `long:"certlocation" description:"Set location of certificate (Optional, for x509 certificates)"`

	// Database options
	DBHost     string `long:"dbhost" description:"Set DB host to connect to (Optional, for SQL)"`
	DBPort     uint16 `long:"dbport" description:"Set DB port to connect to (Optional, for SQL)"`
	DBUsername string `long:"dbusername" description:"Set DB username (Optional, for SQL)"`
	DBPassword string `long:"dbpassword" description:"Set DB password (Optional, for SQL)"`
	DBName     string `long:"dbname" description:"Set DB Name (Optional for PostgreSQL)"`

	// Logging options
	LogLevel []bool `short:"v" description:"Set verbosity level to verbose (-v), very verbose (-vv), or very very verbose (-vvv)"`
}

func main() {
	var err error

	conf := rmcConfig{
		RMCHomeDir:     defaultRMCHomeDir,
		ConfigFileName: defaultConfigFilename,
		LogFileName:    defaultLogFilename,
		HTTPPort:       defaultHTTPPort,
		HTTPHost:       defaultHTTPHost,
		CertLocation:   defaultCertLocation,
		DBUsername:     defaultDBUsername,
		DBPassword:     defaultDBPassword,
		DBHost:         defaultDBHost,
		DBPort:         defaultDBPort,
		DBName:         defaultDBName,
	}

	// Setup conf
	rmcSetup(&conf)

	// Start the database using config file options
	var db *rmcpostgres.RMCPostgres
	if err = db.InitializeDB(conf.DBUsername, conf.DBPassword, conf.DBName, conf.DBHost, conf.DBPort); err != nil {
		logging.Fatalf("Error initializing PostgreSQL DB: %s", err)
		return
	}

	if err = db.SetupClient(); err != nil {
		logging.Fatalf("Error setting up PostgreSQL client: %s", err)
		return
	}

	rmcServer := rmcserver.InitServer(db)

	if err = rmcServer.LoadPEMCertificateFromFile(conf.CertLocation); err != nil {
		logging.Fatalf("Error loading PEM Certificate from file: %s", err)
		return
	}

	// Load the port from its environment variable. If unspecified for whatever reason, default to

	// Setup a static file handler. All routes prefixed with '/static/' will be routed appropriately
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./www/static/templates/index.html")
	})
	http.ListenAndServe(conf.HTTPHost+":"+string(conf.HTTPPort), nil)
}
