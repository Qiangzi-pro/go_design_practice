package dependency_injection

import (
	"testing"
)

type RateLimiter struct {
}

func TestDemo(t *testing.T) {
	var applicationContext ApplicationContext = NewClassPathXmlApplicationContext("beans.xml")
	var rateLimiter RateLimiter = applicationContext.getBean("rateLimiter").(RateLimiter)
	t.Log(rateLimiter)
}
