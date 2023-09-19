package zeus

import (
	"errors"

	"gorm.io/gorm"
)

type Environment struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func Getenv(key string) (string, error) {
	if conn == nil {
		return "", errors.New("database connection is not initialized")
	}

	var env Environment

	if err := conn.Where("key = ?", key).First(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("key not found")
		}
		return "", err
	}

	return env.Value, nil
}

func Setenv(key string, value string) error {
	if conn == nil {
		return errors.New("database connection is not initialized")
	}

	env := Environment{Key: key, Value: value}

	if err := conn.Where("key = ?", key).Assign(env).FirstOrCreate(&env).Error; err != nil {
		return err
	}

	return nil
}

func Environ() ([]Environment, error) {
	if conn == nil {
		return nil, errors.New("database connection is not initialized")
	}

	var env []Environment

	if err := conn.Find(&env).Error; err != nil {
		return nil, err
	}

	return env, nil
}
