// Package http API.
// @title bbs
// @version 1.1
// @description bbs论坛
// @termsOfService https://github.com/swaggo/swag

// @contact.name shenhy
// @contact.email shenhy

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

import (
	_ "gocamp/ch4/bbs/app/http/swagger"
)
