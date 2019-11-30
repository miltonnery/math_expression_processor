package environment

import (
	"github.com/spf13/viper"
	"time"
)

//INTERFACES
type Reader interface {
	GetAllKeys() []string
	Get(key string) interface{}
	GetBool(key string) bool
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetDuration(key string) time.Duration
}

//STRUCTS
type Environment struct {
	viper *viper.Viper
}

//CONSTRUCTOR
func NewEnvironment(v *viper.Viper) *Environment {
	return &Environment{viper: v}
}

//GETTERS
func (v Environment) GetAllKeys() []string {
	return v.viper.AllKeys()
}

func (v Environment) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v Environment) GetBool(key string) bool {
	return v.viper.GetBool(key)
}

func (v Environment) GetString(key string) string {
	return v.viper.GetString(key)
}

func (v Environment) GetViper() *viper.Viper {
	return v.viper
}

func (v Environment) GetInt(key string) int {
	return v.viper.GetInt(key)
}

func (v Environment) GetInt64(key string) int64 {
	return v.viper.GetInt64(key)
}

func (v Environment) GetDuration(key string) time.Duration {
	return v.viper.GetDuration(key)
}

//SETTERS
func (v Environment) SetKey(key string, value interface{}) {
	v.viper.Set(key, value)
}
