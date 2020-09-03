package dependency_injection

type Scope int32

const (
	SINGLETON Scope = 0
	PROTOTYPE Scope = 1
)

// 配置文件解析类
type XmlBeanConfigParser struct {
}

func (parser *XmlBeanConfigParser) parse(configContent string) []BeanDefinition {
	bds := make([]BeanDefinition, 0)
	return bds
}

type BeanDefinition struct {
	id              string
	className       string
	constructorArgs []ConstructorArg
	scope           Scope
	lazyInit        bool
}

func (bd *BeanDefinition) isLazyInit() bool {
	return bd.lazyInit
}

func (bd *BeanDefinition) isSingleton() bool {
	return bd.scope == SINGLETON
}

type ConstructorArg struct {
	isRef bool
	kind  interface{}
	arg   interface{}
}
