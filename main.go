package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/IsaqueGeraldo/agni"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conn *gorm.DB

type Environment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	Bootstrap()
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
			err := Setenv(key, value)
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
			value, err := Getenv(key)
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
			err := RenameKey(oldKey, newKey)
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
			envs, err := Find(key)
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
			envs, err := Environ()
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
			err := Unsetenv(key)
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
				err := Clearenv()
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

func Bootstrap() {
	var err error

	conn, err = gorm.Open(sqlite.Open("zeus.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		agni.Println("[zeus]: "+err.Error(), agni.RedText)
	}

	conn.AutoMigrate(&Environment{})
}

func Getenv(key string) (string, error) {
	if conn == nil {
		return "", errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	var env Environment

	if err := conn.Where("key = ?", key).First(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("key not found")
		}
		return "", err
	}

	return env.Value, nil
}

func Find(key string) ([]Environment, error) {
	if conn == nil {
		return nil, errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	var env []Environment

	if err := conn.Where("key LIKE ?", key).Find(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no records found matching the key")
		}
		return nil, err
	}

	return env, nil
}

func Setenv(key string, value string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	env := Environment{Key: key, Value: value}

	if err := conn.Where("key = ?", key).Assign(env).FirstOrCreate(&env).Error; err != nil {
		return err
	}

	return nil
}

func Unsetenv(key string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	env := Environment{Key: key}

	if err := conn.Where("key = ?", key).Delete(&env).Error; err != nil {
		return err
	}

	return nil
}

func Environ() ([]Environment, error) {
	if conn == nil {
		return nil, errors.New("the database connection is not initialized")
	}

	var env []Environment

	if err := conn.Find(&env).Error; err != nil {
		return nil, err
	}

	return env, nil
}

func Clearenv() error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	return conn.Exec("DELETE FROM environments").Error
}

func RenameKey(oldKey, newKey string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	oldKey = sanitizeKey(oldKey)
	newKey = sanitizeKey(newKey)

	var env Environment
	if err := conn.Where("key = ?", oldKey).First(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("key not found")
		}
		return err
	}

	var existingEnv Environment
	if err := conn.Where("key = ?", newKey).First(&existingEnv).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if existingEnv.ID != 0 {
		return errors.New("new key already exists")
	}

	env.Key = newKey
	if err := conn.Save(&env).Error; err != nil {
		return err
	}

	return nil
}

func sanitizeKey(key string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9_]+")
	cleaned := regex.ReplaceAllString(key, "_")

	cleaned = strings.TrimSuffix(cleaned, "_")

	return strings.ToUpper(cleaned)
}
