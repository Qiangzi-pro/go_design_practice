package dependency_injection

type ApplicationContext interface {
	getBean(beanId string) interface{}
}

type BeanConfigParser interface { //todo need to more compatible
	//parseStream(in io.Reader) []BeanDefinition
	parse(configContent string) []BeanDefinition //todo need to be *[] ？？
}
