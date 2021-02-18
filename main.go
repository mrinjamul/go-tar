package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrinjamul/go-tar/tarwarper"
	flag "github.com/spf13/pflag"
)

var (
	flagHelp    bool
	flagCreate  string
	flagExtract string
	flagTarName string
)

func main() {
	flag.Parse()
	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage(true)
	}

	if flagHelp {
		printUsage(false)
	}
	// Create a tar
	if flagCreate != "" {
		if _, err := os.Stat(flagCreate); os.IsNotExist(err) {
			log.Fatalf("File or direcotroy does not exists: %s", flagCreate)
			os.Exit(1)
		}

		files := []string{flagCreate}
		var tarName string
		if flagTarName != "" {
			tarName = flagTarName
		} else {
			tarName = flagCreate + ".tgz"
		}

		fmt.Println("Creating tar with the name " + tarName)
		err := tarwarper.CreateTar(files, tarName)
		if err != nil {
			log.Println(err)
			fmt.Println("Error: Something went wrong")
		}
		fmt.Printf("%s created successfully\n", flagCreate)
	}
	// Extract a tar
	if flagExtract != "" {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf(path + " ")
		fmt.Println("Extracting " + flagExtract)
		err = tarwarper.ExtarctTar(path, flagExtract)
		if err != nil {
			log.Println(err)
			fmt.Println("Error: Something went wrong")
		} else {
			fmt.Printf("%s extracted successfully\n", flagExtract)
		}
	}
}

func init() {
	flag.BoolVarP(&flagHelp, "help", "h", false, "help message")
	flag.StringVarP(&flagCreate, "create", "c", "", "Create Tar")
	flag.StringVarP(&flagExtract, "extract", "x", "", "Extract Tar")
	flag.StringVarP(&flagTarName, "output", "o", "", "Output name")
}

func printUsage(flagExit bool) {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	if flagExit {
		os.Exit(1)
	}
	os.Exit(0)
}
