package config

import (
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/support/path"
	s3facades "github.com/goravel/s3/facades"

	"goravel/app/facades"
)

func init() {
	config := facades.Config()
	config.Add("filesystems", map[string]any{
		"default": "s3",

		// Filesystem Disks
		//
		// Here you may configure as many filesystem "disks" as you wish, and you
		// may even configure multiple disks of the same driver. Defaults have
		// been set up for each driver as an example of the required values.
		//
		// Supported Drivers: "local", "custom"
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   path.Storage("app"),
			},
			"public": map[string]any{
				"driver": "local",
				"root":   path.Storage("app/public"),
				"url":    config.Env("APP_URL", "").(string) + "/storage",
			},
			"s3": map[string]any{
				"driver": "custom",
				"key":    config.Env("AWS_ACCESS_KEY_ID"),
				"secret": config.Env("AWS_ACCESS_KEY_SECRET"),
				"region": config.Env("AWS_REGION"),
				"bucket": config.Env("AWS_BUCKET"),
				"url":    config.Env("AWS_URL"),
				"via": func() (filesystem.Driver, error) {
					return s3facades.S3("s3") // The `s3` value is the `disks` key
				},
			},
		},
	})
}
