package cache

import "golang.org/x/sync/singleflight"

var g = &singleflight.Group{}

//func getDataSingleFlight(key string) (interface{}, error) {
//	v, err, _ := g.Do(key, func() (interface{}, error) {
//		// 查缓存
//		data, err := getDataFromCache(key)
//		if err == nil {
//			return data, nil
//		}
//		if err == errNotFound {
//			// 查DB
//			data, err := getDataFromDB(key)
//			if err == nil {
//				setCache(data) // 设置缓存
//				return data, nil
//			}
//			return nil, err
//		}
//		return nil, err // 缓存出错直接返回，防止灾难传递至DB
//	})
//
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
