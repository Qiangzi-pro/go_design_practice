package dependency_injection

import (
	"errors"
)

type BeansFactory struct {
	singletonObject map[string]interface{} //todo consider concurrent situation
	beanDefinitions map[string]BeanDefinition
}

func (factory *BeansFactory) addBeanDefinitions(beanDefinitionList []BeanDefinition) {
	for _, beanDefinition := range beanDefinitionList {
		if _, ok := factory.beanDefinitions[beanDefinition.id]; !ok {
			factory.beanDefinitions[beanDefinition.id] = beanDefinition
		}
	}

	for _, beanDefinition := range beanDefinitionList {
		if beanDefinition.isLazyInit() == false && beanDefinition.isSingleton() {
			factory.createBean(beanDefinition)
		}
	}
}

func (factory *BeansFactory) getBean(beanId string) interface{} {
	beanDefinition, ok := factory.beanDefinitions[beanId]
	if !ok {
		errors.New("Bean is not defined: " + beanId)
	}
	return factory.createBean(beanDefinition)
}

func (factory *BeansFactory) createBean(definition BeanDefinition) interface{} {
	//todo 待反射实现
	return string("mock")
}
