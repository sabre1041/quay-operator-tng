package redis

import (
	"errors"

	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
)

// RedisFieldGroup represents the RedisFieldGroup config fields
type RedisFieldGroup struct {
	BuildlogsRedis  *BuildlogsRedisStruct  `default:"" validate:"" json:"BUILDLOGS_REDIS" yaml:"BUILDLOGS_REDIS"`
	UserEventsRedis *UserEventsRedisStruct `default:"" validate:"" json:"USER_EVENTS_REDIS" yaml:"USER_EVENTS_REDIS"`
}

// UserEventsRedisStruct represents the UserEventsRedisStruct config fields
type UserEventsRedisStruct struct {
	Password string `default:"" validate:"" json:"password" yaml:"password"`
	Port     int    `default:"" validate:"" json:"port" yaml:"port"`
	Host     string `default:"" validate:"" json:"host" yaml:"host"`
}

// BuildlogsRedisStruct represents the BuildlogsRedisStruct config fields
type BuildlogsRedisStruct struct {
	Password string `default:"" validate:"" json:"password" yaml:"password"`
	Port     int    `default:"" validate:"" json:"port" yaml:"port"`
	Host     string `default:"" validate:"" json:"host" yaml:"host"`
}

// NewRedisFieldGroup creates a new RedisFieldGroup
func NewRedisFieldGroup(fullConfig map[string]interface{}) (*RedisFieldGroup, error) {
	newRedisFieldGroup := &RedisFieldGroup{}
	defaults.Set(newRedisFieldGroup)

	if value, ok := fullConfig["BUILDLOGS_REDIS"]; ok {
		var err error
		value := shared.FixInterface(value.(map[interface{}]interface{}))
		newRedisFieldGroup.BuildlogsRedis, err = NewBuildlogsRedisStruct(value)
		if err != nil {
			return newRedisFieldGroup, err
		}
	}
	if value, ok := fullConfig["USER_EVENTS_REDIS"]; ok {
		var err error
		value := shared.FixInterface(value.(map[interface{}]interface{}))
		newRedisFieldGroup.UserEventsRedis, err = NewUserEventsRedisStruct(value)
		if err != nil {
			return newRedisFieldGroup, err
		}
	}

	return newRedisFieldGroup, nil
}

// NewUserEventsRedisStruct creates a new UserEventsRedisStruct
func NewUserEventsRedisStruct(fullConfig map[string]interface{}) (*UserEventsRedisStruct, error) {
	newUserEventsRedisStruct := &UserEventsRedisStruct{}
	defaults.Set(newUserEventsRedisStruct)

	if value, ok := fullConfig["password"]; ok {
		newUserEventsRedisStruct.Password, ok = value.(string)
		if !ok {
			return newUserEventsRedisStruct, errors.New("password must be of type string")
		}
	}
	if value, ok := fullConfig["port"]; ok {
		newUserEventsRedisStruct.Port, ok = value.(int)
		if !ok {
			return newUserEventsRedisStruct, errors.New("port must be of type int")
		}
	}
	if value, ok := fullConfig["host"]; ok {
		newUserEventsRedisStruct.Host, ok = value.(string)
		if !ok {
			return newUserEventsRedisStruct, errors.New("host must be of type string")
		}
	}

	return newUserEventsRedisStruct, nil
}

// NewBuildlogsRedisStruct creates a new BuildlogsRedisStruct
func NewBuildlogsRedisStruct(fullConfig map[string]interface{}) (*BuildlogsRedisStruct, error) {
	newBuildlogsRedisStruct := &BuildlogsRedisStruct{}
	defaults.Set(newBuildlogsRedisStruct)

	if value, ok := fullConfig["password"]; ok {
		newBuildlogsRedisStruct.Password, ok = value.(string)
		if !ok {
			return newBuildlogsRedisStruct, errors.New("password must be of type string")
		}
	}
	if value, ok := fullConfig["port"]; ok {
		newBuildlogsRedisStruct.Port, ok = value.(int)
		if !ok {
			return newBuildlogsRedisStruct, errors.New("port must be of type int")
		}
	}
	if value, ok := fullConfig["host"]; ok {
		newBuildlogsRedisStruct.Host, ok = value.(string)
		if !ok {
			return newBuildlogsRedisStruct, errors.New("host must be of type string")
		}
	}

	return newBuildlogsRedisStruct, nil
}
