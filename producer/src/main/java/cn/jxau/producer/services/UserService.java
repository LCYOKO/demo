package cn.jxau.producer.services;

import cn.jxau.producer.entity.vo.UserVO;
import cn.jxau.producer.excetion.BaseException;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 10:27 下午
 * @Version 1.0
 */
public interface UserService {


    void registerUser(UserVO userVO) throws BaseException;
    UserVO  selectUserById(Integer id);
    UserVO  login(String phone,String password);
}
