package core

type Config struct {
	configuration map[string]string
}

func (config Config) Set(key, value string)  {
	config.configuration[key] = value
}

func (config Config) SetByMap(kvs map[string]string)  {
	for k,v := range kvs {
		config.configuration[k] = v
	}
}

func (config Config) Get(key,def string) (string,bool) {
	value,exist := config.configuration[key]
	if exist {
		return value,true
	} else {
		return def,false
	}
}