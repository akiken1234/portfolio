module github.com/akimotokensaku/portfolio/go

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/jinzhu/gorm v1.9.16
	local.packages/controller v0.0.0-00010101000000-000000000000
	local.packages/db v0.0.0-00010101000000-000000000000
	local.packages/model v0.0.0-00010101000000-000000000000
)

replace (
	local.packages/controller => ./controller
	local.packages/db => ./db
	local.packages/model => ./model
)
