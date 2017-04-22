package core

type Config interface {
	Set(key, value string)
	SetByMap(kvs map[string]string)
	SetByMapWithPrefix(prefix string,kvs map[string]string)
	Get(key,def string) string
}

type config struct {
	configuration map[string]string
}

func NewConfig() Config {
	conf := &config{}
	conf.configuration = make(map[string]string,20)
	return conf
}

func (config *config) Set(key, value string)  {
	config.configuration[key] = value
}

func (config *config) SetByMap(kvs map[string]string)  {
	for k,v := range kvs {
		config.configuration[k] = v
	}
}
func (config *config) SetByMapWithPrefix(prefix string,kvs map[string]string)  {
	for k,v := range kvs {
		config.configuration[prefix+"."+k] = v
	}
}


func (config *config) Get(key,def string) (string) {
	value,exist := config.configuration[key]
	if exist {
		return value
	} else {
		return def
	}
}