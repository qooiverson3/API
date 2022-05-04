/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"ces-api/pkg/delivery"
	"ces-api/pkg/service"
	"ces-api/pkg/storage"
	"fmt"
	"io"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbConnect := os.Getenv("DB_CONNECT")

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPass,
			dbConnect,
			"CES",
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		logFile, _ := os.Create("logs/gin.log")
		gin.DefaultWriter = io.MultiWriter(logFile)

		httpEngine := gin.Default()
		httpEngine.Use(gzip.Gzip(gzip.BestSpeed))
		// get instance list
		r := storage.NewInstanceRepo(db)
		s := service.NewInstanceService(r)
		h := delivery.NewInstanceHandler(s)

		h.Router(httpEngine)
		httpEngine.Run(":8089")

	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
