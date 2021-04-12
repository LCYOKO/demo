package cn.jxau.producer.services.impl;

import cn.jxau.common.validator.ValidationResult;
import cn.jxau.producer.entity.UserInfo;
import cn.jxau.producer.entity.vo.UserVO;
import cn.jxau.producer.excetion.BaseException;
import cn.jxau.producer.excetion.BizErrorEnum;
import cn.jxau.producer.mapper.UserInfoMapper;
import cn.jxau.producer.services.UserService;
import cn.jxau.producer.validator.ValidatorImpl;
import com.sun.org.apache.bcel.internal.generic.NEW;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.sound.sampled.Line;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 10:27 下午
 * @Version 1.0
 */
@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private ValidatorImpl validator;
    @Autowired
    private UserInfoMapper userInfoMapper;
    @Override
    @Transactional(rollbackFor = Exception.class)
    public void registerUser(UserVO userVO) throws BaseException {
        if(userVO==null){
            throw new BaseException(BizErrorEnum.PARAMETER_VALIDATION_ERROR);
        }
        ValidationResult validationResult = validator.validate(userVO);
        if(validationResult.isHasErrors()){
            throw new BaseException(BizErrorEnum.PARAMETER_VALIDATION_ERROR,validationResult.getErrorsMsg());
        }
        UserInfo userInfo=userVo2UserInfo(userVO);
        if(userInfoMapper.insertSelective(userInfo)==0){
            throw new BaseException(BizErrorEnum.USER_REGISTER_FAIL);
        }


    }

    @Override
    public UserVO selectUserById(Integer id) {
        UserInfo userInfo = userInfoMapper.selectByPrimaryKey(id);
        return userInfo2UserVo(userInfo);
    }

    @Override
    public UserVO login(String phone, String password) {
        return null;
    }


    private UserInfo  userVo2UserInfo(UserVO vo){
        UserInfo info = new UserInfo();
        BeanUtils.copyProperties(vo,info);
        return info;
    }

    private UserVO userInfo2UserVo(UserInfo info){
        if(info==null){
            return null;
        }
        UserVO userVO = new UserVO();
        BeanUtils.copyProperties(info,userVO);
        return userVO;
    }




}
