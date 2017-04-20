package core


type Container struct {
	items map[string]func() struct{}
	instances map[string]struct{}
	services map[string]*Service
	middlewares map[string]*MiddleWare
}

/**
 * Allocate memory for Container
 */
func (container *Container) Init ()  {
	container.middlewares = make(map[string]*MiddleWare)
	container.services = make(map[string]*Service)
}


func (container *Container) Bind(name string,buildFunc func() struct{})  {
	container.items[name] = buildFunc;
}

func (container *Container) Build(name string) (struct{},bool) {
	instance, exist := container.instances[name]
	if exist {
		return instance,true
	} else {
		if buildFunc,exist := container.items[name];exist {
			container.instances[name] = buildFunc()
			return container.instances[name],true
		} else {
			return struct {}{},false
		}
	}
}

func (container *Container) RegisterService(name string, service Service)  {
	container.services[name] = &service
}

func (container *Container) Service(name string) (Service) {
	service,exist := container.services[name]
	if !exist {
		return nil
	}
	return service
}

/**
 * Register a MiddleWare in the Container
 */
func (container *Container) RegisterMiddleWare(name string, middleware MiddleWare)  {
	container.middlewares[name] = &middleware
}

/**
 * Get a MiddleWare in the Container
 */
func (container *Container) Middleware(name string) (*MiddleWare) {
	middleware,exist := container.middlewares[name]
	if !exist {
		return nil
	}
	return middleware
}
