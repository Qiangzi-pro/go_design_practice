package article41_43

var instance *singleton

func GetInstance1() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
