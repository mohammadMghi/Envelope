package envelope

import "time"

const(
	DefueltTimeExpetion = 0
)

type Cache struct{
	object map[string]objects
	defaultExpiration time.Duration
}


type objects struct{
	item interface{}
	timeExpression int64
}


func NewCacheEnv(defaultExpiration time.Duration)Cache{
	obj := make(map[string]objects)
	return Cache{
		object : obj,
		defaultExpiration :defaultExpiration,
	}
}

func (cache *Cache)set(key string ,item interface{} , timeExp time.Duration){
	var exp int64
	if timeExp == DefueltTimeExpetion{
		timeExp = cache.defaultExpiration
	}

	if timeExp > 0 {
		exp  = time.Now().Add(timeExp).UnixNano()
	}

	cache.object[key] = objects{
		item: item,
		timeExpression : exp,
	}
}

func (cache *Cache)get(key string) (interface{} , bool){

	//check the obj existed in cache
	i , f := cache.object[key]

	//object dose not existed
	if f == false{
		return nil , false
	} 

	if i.timeExpression < 0{
		if !(time.Now().UnixNano() > i.timeExpression) {
			return cache.object[key] , false
		}
	}

	return nil ,false
 
}
 