package cn.jxau.producer.mapper;

import cn.jxau.producer.entity.ItemStock;

public interface ItemStockMapper {
    int deleteByPrimaryKey(Integer id);

    int insert(ItemStock record);

    int insertSelective(ItemStock record);

    ItemStock selectByPrimaryKey(Integer id);

    int updateByPrimaryKeySelective(ItemStock record);

    int updateByPrimaryKey(ItemStock record);
}