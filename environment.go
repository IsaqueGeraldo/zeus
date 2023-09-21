package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Environment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func Unsetenv(key string) error {
	if conn == nil {
		return errors.New("database connection is not initialized")
	}

	env := Environment{Key: key}

	if err := conn.Where("key = ?", key).Delete(&env).Error; err != nil {
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

func Clearenv() error {
	if conn == nil {
		return errors.New("database connection is not initialized")
	}

	return conn.Migrator().DropTable(
		Environment{},
	)
}
