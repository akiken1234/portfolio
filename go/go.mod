module github.com/akimotokensaku/portfolio/go

go 1.16

require (
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
	local.packages/db v0.0.0-00010101000000-000000000000
	local.packages/model v0.0.0-00010101000000-000000000000
)

replace (
	local.packages/db => ./db
	local.packages/model => ./model
)
