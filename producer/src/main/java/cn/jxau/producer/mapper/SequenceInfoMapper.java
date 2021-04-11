package cn.jxau.producer.mapper;

import cn.jxau.producer.entity.SequenceInfo;

public interface SequenceInfoMapper {
    int deleteByPrimaryKey(Integer id);

    int insert(SequenceInfo record);

    int insertSelective(SequenceInfo record);

    SequenceInfo selectByPrimaryKey(Integer id);

    int updateByPrimaryKeySelective(SequenceInfo record);

    int updateByPrimaryKey(SequenceInfo record);
}