package dependency_injection

import (
	"io/ioutil"
	"os"
)

// 组织类并提供执行入口

type ClassPathXmlApplicationContext struct {
	beansFactory     BeansFactory
	beanConfigParser BeanConfigParser
}

func NewClassPathXmlApplicationContext(configLocation string) *ClassPathXmlApplicationContext {
	cpXmlContext := &ClassPathXmlApplicationContext{
		beansFactory:     BeansFactory{},
		beanConfigParser: &XmlBeanConfigParser{},
	}
	cpXmlContext.loadBeanDefinitions(configLocation)
	return cpXmlContext
}

func (context *ClassPathXmlApplicationContext) loadBeanDefinitions(configLocation string) {
	var f *os.File
	f, err := os.Open(configLocation)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ctnBytes, _ := ioutil.ReadAll(f)
	content := string(ctnBytes)
	beanDefinitions := context.beanConfigParser.parse(content)
	context.beansFactory.addBeanDefinitions(beanDefinitions)
}

func (context *ClassPathXmlApplicationContext) getBean(beanId string) interface{} {
	return context.beansFactory.getBean(beanId)
}
