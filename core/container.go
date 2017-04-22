package core

type Container interface {
	Bind(name string,buildFunc func(Container) interface{})
	Put(name string, inst interface{})
	Get(name string) (interface{},bool)

	RegisterService(name string, service Service)
	Service(name string) *Service
	RegisterMiddleWare(name string, middleware MiddleWare)
	Middleware(name string) (*MiddleWare)
	RegisterDB(db Manager)
	DB() Manager


	AfterBootstrap(callback func(Container))
	Bootstrap()()
}

type container struct {
	items map[string]func(Container) interface{}
	instances map[string]interface{}
	services map[string]*Service
	middlewares map[string]*MiddleWare
	db Manager
	bootstrapFunc []func(Container)

}

/**
 * Create and allocate memory for Container
 */
func NewContainer() Container  {
	ctn := &container{}
	ctn.services = make(map[string]*Service)
	ctn.middlewares = make(map[string]*MiddleWare)
	ctn.bootstrapFunc = make([]func(Container),0)
	ctn.items = make(map[string]func(Container) interface{})
	ctn.instances = make(map[string] interface{})
	return ctn
}

func (ctn *container) Bind(name string,buildFunc func(Container) interface{})  {
	ctn.items[name] = buildFunc;
}

func (ctn *container) Put(name string, inst interface{})  {
	ctn.instances[name] = inst
}

func (ctn *container) Get(name string) (interface{},bool) {
	instance, exist := ctn.instances[name]
	if exist {
		return instance,true
	} else {
		if buildFunc,exist := ctn.items[name];exist {
			ctn.instances[name] = buildFunc(ctn)
			return ctn.instances[name],true
		} else {
			return nil,false
		}
	}
}

/**
 * Register Service in the Container
 */
func (ctn *container) RegisterService(name string, service Service)  {
	ctn.services[name] = &service
}

/**
 * Get a Service in the Container
 */
func (ctn *container) Service(name string) (*Service) {
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

func (ctn *container) RegisterDB(db Manager)  {
	ctn.db = db
}

func (ctn *container) DB() Manager  {
	return ctn.db
}


/**
 * Add callback function when bootstrapping
 */
func (ctn *container) AfterBootstrap(callback func(Container))  {
	ctn.bootstrapFunc = append(ctn.bootstrapFunc,callback)
}

/**
 * Bootstrap the container
 */
func (ctn *container) Bootstrap()  {
	for i:=0; i<len(ctn.bootstrapFunc); i++ {
		ctn.bootstrapFunc[i](ctn)
	}
}
