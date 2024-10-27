package zeus

import (
	"fmt"    // For formatting and printing functions
	"regexp" // For regular expressions
	"strings"

	"gorm.io/driver/sqlite" // SQLite driver for GORM
	"gorm.io/gorm"          // GORM ORM
)

var conn *gorm.DB // Global variable to hold the database connection

// Environment represents an environment variable with a key and value
type Environment struct {
	Key   string `json:"key" gorm:"primaryKey"` // Unique key for the variable
	Value string `json:"value"`                 // Value associated with the key
}

// Bootstrap initializes the connection to the SQLite database
func Bootstrap() {
	db, err := gorm.Open(sqlite.Open("zeus.db"), &gorm.Config{}) // Opens the connection to the database
	if err != nil {
		fmt.Printf("[zeus]: %v \n", err) // Displays error if the connection fails
		return
	}

	if err := db.AutoMigrate(&Environment{}); err != nil {
		fmt.Printf("[zeus]: %v \n", err) // Displays error if migration fails
		return
	}

	conn = db // Assigns the connection to the global variable

	fmt.Println("[zeus]: database connection established") // Success message
}

// normalizeKey ensures the key is in UPPER_SNAKE_CASE format
func normalizeKey(key string) string {
	// Remove non-alphabetic characters except spaces
	reg := regexp.MustCompile(`[^A-Za-z\s]+`)
	key = reg.ReplaceAllString(key, "")

	// Replace spaces with underscores
	key = strings.ReplaceAll(key, " ", "_")

	// Convert to uppercase
	return strings.ToUpper(key)
}

// Getenv retrieves the environment variable for the given key
func Getenv(key string) (Environment, error) {
	key = normalizeKey(key) // Normalize the key
	var env Environment     // Variable to hold the retrieved environment

	result := conn.First(&env, "key = ?", key) // Finds the first record matching the key
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return Environment{}, fmt.Errorf("key '%s' not found", key) // Returns error if key not found
		}
		return Environment{}, result.Error // Returns any other error encountered
	}

	return env, nil // Returns the retrieved environment variable
}

// Setenv sets a new environment variable with the given key and value
func Setenv(key, value string) error {
	key = normalizeKey(key) // Normalize the key

	var existing Environment
	if result := conn.First(&existing, "key = ?", key); result.Error == nil {
		return fmt.Errorf("key '%s' already exists", key) // Returns error if key already exists
	}

	env := Environment{Key: key, Value: value} // Creates a new environment variable

	result := conn.Save(&env) // Saves the environment variable to the database
	if result.Error != nil {
		return fmt.Errorf("error saving variable: %v", result.Error) // Returns error if save fails
	}

	return nil // Returns nil if successful
}

// Clearenv removes the environment variable for the given key
func Clearenv(key string) error {
	key = normalizeKey(key) // Normalize the key

	result := conn.Delete(&Environment{}, "key = ?", key) // Deletes the environment variable from the database
	if result.Error != nil {
		return fmt.Errorf("error removing variable: %v", result.Error) // Returns error if deletion fails
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("key '%s' not found", key) // Returns error if no rows were affected (key not found)
	}

	return nil // Returns nil if successful
}

// Environ lists all environment variables
func Environ() ([]Environment, error) {
	var envs []Environment // Slice to hold all environment variables

	result := conn.Find(&envs) // Finds all environment variables in the database
	if result.Error != nil {
		return nil, fmt.Errorf("error when listing environment: %v", result.Error) // Returns error if listing fails
	}

	return envs, nil // Returns the list of environment variables
}
