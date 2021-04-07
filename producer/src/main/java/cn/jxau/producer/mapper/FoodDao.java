package cn.jxau.producer.mapper;

import cn.jxau.producer.entity.Food;
import org.apache.ibatis.annotations.Param;

public interface FoodDao {
    int deleteByPrimaryKey(Integer id);

    int insert(Food record);

    int insertSelective(Food record);

    Food selectByPrimaryKey(Integer id);

    int updateByPrimaryKeySelective(Food record);

    int updateByPrimaryKey(Food record);

    int decFoodById(@Param("id") Integer foodId, @Param("amount")Integer amount);
}