package cn.jxau.producer.controller;

import cn.jxau.common.response.CommonResponse;
import cn.jxau.common.util.MD5Util;
import cn.jxau.producer.entity.vo.UserVO;
import cn.jxau.producer.excetion.BaseException;
import cn.jxau.producer.excetion.BizErrorEnum;
import cn.jxau.producer.services.UserService;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.io.UnsupportedEncodingException;
import java.security.NoSuchAlgorithmException;

/**
 * @Author l
 * @Date 2021/4/12 10:44
 * @Version 1.0
 */
@RestController
@RequestMapping(value ="/user")
public class UserController {

    @Autowired
    private UserService userService;


    @GetMapping("/register")
    public CommonResponse register( UserVO userVO) throws BaseException {
           userService.registerUser(userVO);
        return CommonResponse.create(null,"fail");
    }

    @PostMapping("/get")
    public CommonResponse getUser(@RequestParam() Integer id) throws BaseException {
        if(id==null){
            throw new BaseException(BizErrorEnum.PARAMETER_VALIDATION_ERROR);
        }
        return  CommonResponse.create(userService.selectUserById(id),"fail");
    }

    @PostMapping("/login")
    public CommonResponse login(@RequestBody String phone,
                                @RequestBody String password) throws BaseException, UnsupportedEncodingException, NoSuchAlgorithmException {
          if(StringUtils.isEmpty(phone) || StringUtils.isEmpty(password)){
              throw new BaseException(BizErrorEnum.PARAMETER_VALIDATION_ERROR);
          }
        UserVO vo = userService.login(phone, MD5Util.EncodeByMd5(password));

        return CommonResponse.create(null,"success");
    }


}
