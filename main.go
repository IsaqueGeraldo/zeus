package main

import (
	"fmt"
	"os"

	"github.com/IsaqueGeraldo/agni"
	"github.com/IsaqueGeraldo/odin"
	"github.com/spf13/cobra"
)

func init() {
	odin.Bootstrap()
}

func main() {
	var rootCmd = &cobra.Command{Use: "zeus"}

	var setEnv = &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Define an environment variable",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			value := args[1]
			err := odin.Setenv(key, value)
			if err != nil {
				agni.Println("[zeus]: Error setting the environment variable: "+err.Error(), agni.RedText)
			} else {
				agni.Println("[zeus]: Environment variable set successfully!")
			}
		},
	}

	var getEnv = &cobra.Command{
		Use:   "get [key]",
		Short: "Get the value of an environment variable",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			value, err := odin.Getenv(key)
			if err != nil {
				agni.Println("[zeus]: Error getting the value of the environment variable: "+err.Error(), agni.RedText)
			} else {
				agni.Println("Value of environment variable '" + key + "': " + value)
			}
		},
	}

	var rename = &cobra.Command{
		Use:   "rename [oldkey] [newkey]",
		Short: "Rename an environment variable",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			oldKey := args[0]
			newKey := args[1]
			err := odin.RenameKey(oldKey, newKey)
			if err != nil {
				agni.Println("[zeus]: Error renaming the environment variable: "+err.Error(), agni.RedText)
			} else {
				agni.Println("[zeus]: Environment variable renamed successfully!")
			}
		},
	}

	var find = &cobra.Command{
		Use:   "find [key]",
		Short: "Search the value of an environment variable",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			envs, err := odin.Find(key)
			if err != nil {
				agni.Println("[zeus]: Error getting the value of the environment variable: "+err.Error(), agni.YellowText)
			} else {
				agni.Println("[zeus]: Environment variables:")
				for _, env := range envs {
					fmt.Printf("%s=%s\n", env.Key, env.Value)
				}
			}
		},
	}

	var environ = &cobra.Command{
		Use:   "environ",
		Short: "List all environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			envs, err := odin.Environ()
			if err != nil {
				agni.Println("[zeus]: Error listing environment variables: "+err.Error(), agni.RedText)
			} else {
				agni.Println("[zeus]: Environment variables:")
				for _, env := range envs {
					fmt.Printf("%s=%s\n", env.Key, env.Value)
				}
			}
		},
	}

	var unsetEnv = &cobra.Command{
		Use:   "unset [key]",
		Short: "Remove an environment variable",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			err := odin.Unsetenv(key)
			if err != nil {
				agni.Println("[zeus]: Error removing the environment variable: "+err.Error(), agni.RedText)
			} else {
				agni.Println("[zeus]: Environment variable removed successfully!")
			}
		},
	}

	var clearEnv = &cobra.Command{
		Use:   "clearenv",
		Short: "Remove all environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("[zeus]: Are you sure you want to clear all environment variables? (yes/no): ")
			var confirmation string
			fmt.Scanln(&confirmation)

			if confirmation == "yes" {
				err := odin.Clearenv()
				if err != nil {
					agni.Println("[zeus]: Error clearing environment variables: "+err.Error(), agni.RedText)
				} else {
					agni.Println("[zeus]: All environment variables have been removed successfully!", agni.GreenText)
				}
			} else {
				agni.Println("[zeus]: Operation canceled. Environment variables were not removed.", agni.YellowText)
			}
		},
	}

	rootCmd.AddCommand(
		setEnv, getEnv, environ, unsetEnv, clearEnv, find, rename,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
