package main

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"github.com/ratemyclass/rmc-web/logging"
)

func createDefaultConfigFile(destinationPath string, destinationFileName string) (err error) {

	// Open our config file w/ directory destinationPath and filename destinationFileName
	var dest *os.File
	if dest, err = os.OpenFile(filepath.Join(destinationPath, destinationFileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
		return
	}

	// Add dest.Close() to the defer stack, this is the last thing that happens when the function returns
	defer dest.Close()

	// Create a new writer for the dest file, write default arguments to it
	writer := bufio.NewWriter(dest)

	// We have no default config arguments (basically just `touch` the file)
	defaultArgs := []byte("")
	if _, err = writer.Write(defaultArgs); err != nil {
		return
	}

	// Add writer.Flush() to the defer stack, this is the second to last thing that happens when the function returns
	defer writer.Flush()

	return
}

func rmcSetup(conf *rmcConfig) {
	var err error

	parser := flags.NewParser(conf, flags.Default)
	if _, err = parser.ParseArgs(os.Args); err != nil {
		// if there are any argument errors, that's a fatal error
		logging.Fatal(err)
		return
	}

	// set default log level
	logging.SetLogLevel(defaultLogLevel)
	if _, err = os.Stat(conf.RMCHomeDir); err != nil {
		logging.Errorf("Error while creating RMC home dir")
	}
	if os.IsNotExist(err) {
		// If this is the first time we're running the server, create a default config file and home directory

		os.Mkdir(conf.RMCHomeDir, 0700)
		logging.Infof("Creating a new default config file")
		if err = createDefaultConfigFile(filepath.Join(conf.RMCHomeDir), defaultConfigFilename); err != nil {
			logging.Fatalf("Error creating a default config file in %v: %s", conf.RMCHomeDir, err)
			return
		}
	}

	// Set the full path of the config file
	conf.ConfigFile = filepath.Join(conf.RMCHomeDir, conf.ConfigFileName)

	// Make sure there's a config file
	if _, err = os.Stat(conf.ConfigFile); err != nil {
		// If there's no config file found in the home directory where the user specified, create one
		logging.Infof("Creating a new config file at %v", conf.ConfigFile)
		if err = createDefaultConfigFile(filepath.Join(conf.RMCHomeDir), conf.ConfigFileName); err != nil {
			logging.Fatalf("Error creating a config file at %v: %s", conf.ConfigFileName, err)
			return
		}
	}

	// Now we can parse the config file
	if err = flags.NewIniParser(parser).ParseFile(conf.ConfigFile); err != nil {
		var isPathError bool
		if _, isPathError = err.(*os.PathError); !isPathError {
			logging.Fatalf("There was a non-path error while parsing the config file: %s", err)
			return
		}
	}

	// Now we can parse the command line options to ensure they take precedence over any config file options
	if _, err = parser.ParseArgs(os.Args); err != nil {
		logging.Fatalf("Error parsing command line arguments: %s", err)
		return
	}

	var logFile *os.File
	if logFile, err = os.OpenFile(filepath.Join(conf.RMCHomeDir, conf.LogFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		logging.Fatalf("Error while opening log file: %s", err)
		return
	}

	// Defer the closing of logfile
	defer logFile.Close()

	// Set the logging package to use this logfile we've specified
	logging.SetLogFile(logFile)

	// Set the log level to default unless we specify otherwise
	logLevel := defaultLogLevel
	logLevelLength := len(conf.LogLevel)
	if logLevelLength <= 3 {
		logLevel = logLevelLength
	} else {
		logLevel = 3
	}

	// Actually set logging package to use this log level
	logging.SetLogLevel(logLevel)

	return
}
