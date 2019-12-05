module github.com/step/sauron_go

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/spf13/viper v1.5.0
	github.com/step/angmar v0.0.0-20191127113211-fbeaab94f9b7
	github.com/step/saurontypes v0.0.0-20191127114135-1c7b69a4e64f
	github.com/step/uruk v0.0.0-20191127114036-eb84283fad8d
)

replace github.com/step/uruk => ../uruk/

replace github.com/step/angmar => ../angmar/

replace github.com/step/saurontypes => ../saurontypes/
