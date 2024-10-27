package zeus

import (
	"fmt"
	"os"
	"regexp"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

type Environment struct {
	Key   string `json:"key" gorm:"primaryKey"`
	Value string `json:"value"`
}

func Bootstrap() {
	dsn := "file:zeus.db?_key=" + os.Getenv("ZEUS_PASSWORD")
	var err error

	conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("[zeus]: %v \n", err)
		return
	}

	fmt.Println("[zeus]: database connection established")
}

func validateKey(key string) error {
	re := regexp.MustCompile(`^[A-Z_]+$`)
	if !re.MatchString(key) {
		return fmt.Errorf("key '%s' is not in UPPER_SNAKE_CASE format", key)
	}
	return nil
}

func Getenv(key string) (Environment, error) {
	var env Environment

	result := conn.First(&env, "key = ?", key)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return Environment{}, fmt.Errorf("key '%s' not found", key)
		}
		return Environment{}, result.Error
	}

	return env, nil
}

func Setenv(key, value string) error {
	if err := validateKey(key); err != nil {
		return err
	}

	var existing Environment
	if result := conn.First(&existing, "key = ?", key); result.Error == nil {
		return fmt.Errorf("key '%s' already exists", key)
	}

	env := Environment{Key: key, Value: value}

	result := conn.Save(&env)
	if result.Error != nil {
		return fmt.Errorf("error saving variable: %v", result.Error)
	}

	return nil
}

func Clearenv(key string) error {
	result := conn.Delete(&Environment{}, "key = ?", key)
	if result.Error != nil {
		return fmt.Errorf("error removing variable: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("key '%s' not found", key)
	}

	return nil
}

func Environ() ([]Environment, error) {
	var envs []Environment

	result := conn.Find(&envs)
	if result.Error != nil {
		return nil, fmt.Errorf("error when listing environment: %v", result.Error)
	}

	return envs, nil
}
