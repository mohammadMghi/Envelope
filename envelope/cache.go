package envelope

import "time"

const(
	DefueltTimeExpetion = 0
)

type Cache struct{
	obj map[string]obj
	timeExpression int64
}


type obj struct{
	item interface{}

}


func (cache *Cache)set(key string ,item interface{} , timeExp time.Duration){
	 

	cache.timeExpression = int64(timeExp)

	cache.obj[key] = obj{
		item: item,
	}
}

func (cache *Cache)get(key string) interface{}{
	if cache[key].timeExpression{}
	
	return cache[key]
}


func (cache *cache)setDefueltTimeExpertion(){
	DefueltTimeExpetion = 
}