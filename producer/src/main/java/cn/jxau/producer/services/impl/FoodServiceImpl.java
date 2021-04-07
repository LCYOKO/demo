package cn.jxau.producer.services.impl;

import cn.jxau.producer.mapper.FoodDao;
import cn.jxau.producer.services.FoodService;
import org.checkerframework.checker.units.qual.A;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @Author l
 * @Date 2021/4/7 15:46
 * @Version 1.0
 */
@Service
public class FoodServiceImpl implements FoodService {
   @Autowired
   private FoodDao foodDao;
    @Override
    public int decFoodById(Integer id, Integer amount) {
        return foodDao.decFoodById(id,amount);
    }
}
