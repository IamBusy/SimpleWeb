package core

type Container interface {
	Init()
	RegisterService(name string, service Service)
	Service(name string) Service
	RegisterMiddleWare(name string, middleware MiddleWare)
	Middleware(name string) (*MiddleWare)
}

type container struct {
	//items map[string]func() struct{}
	//instances map[string]struct{}
	services map[string]*Service
	middlewares map[string]*MiddleWare
}

/**
 * Create and allocate memory for Container
 */
func NewContainer() Container  {
	ctn := &container{}
	ctn.services = make(map[string]*Service)
	ctn.middlewares = make(map[string]*MiddleWare)
	return ctn
}





//func (ctn *container) Bind(name string,buildFunc func() struct{})  {
//	ctn.items[name] = buildFunc;
//}
//
//func (ctn *container) Build(name string) (struct{},bool) {
//	instance, exist := ctn.instances[name]
//	if exist {
//		return instance,true
//	} else {
//		if buildFunc,exist := ctn.items[name];exist {
//			ctn.instances[name] = buildFunc()
//			return ctn.instances[name],true
//		} else {
//			return struct {}{},false
//		}
//	}
//}

/**
 * Register Service in the Container
 */
func (ctn *container) RegisterService(name string, service Service)  {
	ctn.services[name] = &service
}

/**
 * Get a Service in the Container
 */
func (ctn *container) Service(name string) (Service) {
	service,exist := ctn.services[name]
	if !exist {
		return nil
	}
	return service
}

/**
 * Register a MiddleWare in the Container
 */
func (ctn *container) RegisterMiddleWare(name string, middleware MiddleWare)  {
	ctn.middlewares[name] = &middleware
}

/**
 * Get a MiddleWare in the Container
 */
func (ctn *container) Middleware(name string) (*MiddleWare) {
	middleware,exist := ctn.middlewares[name]
	if !exist {
		return nil
	}
	return middleware
}
