package config

type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)

	GetStringDefault(name, defValue string) (configValue string)
	GetIntDefault(name string, defValue int) (configValue int)
	GetBoolDefault(name string, defValue bool) (configValue bool)
	GetFloatDefault(name string, defValue float64) (configValue float64)

	GetSection(sectionName string) (section Configuration, found bool)
}
