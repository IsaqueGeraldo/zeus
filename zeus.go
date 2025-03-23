package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/IsaqueGeraldo/odin"
	"github.com/spf13/cobra"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var source string

func init() {
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "📂 Path to the environment file")
}

var rootCmd = &cobra.Command{
	Use:   "zeus",
	Short: "⚡ Environment variable manager using a database",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if source == "" {
			source = os.Getenv("ZEUS_SOURCE_DIR")
		}

		odin.Bootstrap(source)
	},
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "📝 Set an environment variable",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := sanitizeKey(args[0])
		value := args[1]
		err := odin.Setenv(key, value)
		if err != nil {
			fmt.Printf("❌ Error setting variable: %v\n", err)
		} else {
			fmt.Printf("✅ Variable '%s' set to '%s'\n", key, value)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "🔍 Get the value of an environment variable",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := sanitizeKey(args[0])
		value := odin.Getenv(key)
		if value == "" {
			fmt.Printf("⚠️ Variable '%s' not found\n", key)
		} else {
			fmt.Printf("🔍 %s=%s\n", key, value)
		}
	},
}

var unsetCmd = &cobra.Command{
	Use:   "unset [key]",
	Short: "🗑️ Remove an environment variable",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := sanitizeKey(args[0])
		err := odin.Unsetenv(key)
		if err != nil {
			fmt.Printf("❌ Error removing variable: %v\n", err)
		} else {
			fmt.Printf("🗑️ Variable '%s' removed\n", key)
		}
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "🧹 Clear all environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		odin.Clearenv()
		fmt.Println("🧹 All environment variables have been removed.")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "📜 List all environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		envs := odin.Environ()
		if len(envs) == 0 {
			fmt.Println("⚠️ No environment variables defined.")
		} else {
			fmt.Println("📜 Environment variables list:")
			for _, env := range envs {
				fmt.Println(env)
			}
		}
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "📖 Show help for commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "⌨️ Generate shell completion script",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenBashCompletion(os.Stdout)
	},
}

func main() {
	rootCmd.AddCommand(setCmd, getCmd, unsetCmd, clearCmd, listCmd, helpCmd, completionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sanitizeKey(key string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, key)
	if err != nil {
		return key
	}

	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	sanitized := re.ReplaceAllString(result, "_")

	return strings.ToUpper(sanitized)
}
