package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deacon",
	Short: "deacon",
	Long:  `Deacon is a tool to generate ACLs and certs for managing Kafka clients and having specific certs for each producer/consumer.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test")
	},
}

var aclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Generate ACLs",
	Long:  `ACL generations`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	
	},
}

func init() {
	rootCmd.AddCommand(aclCmd)

	// Defining a flag for the greet command
	greetCmd.Flags().StringP("name", "n", "", "acl to add")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
