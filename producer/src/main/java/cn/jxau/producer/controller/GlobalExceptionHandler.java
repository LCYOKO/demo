package cn.jxau.producer.controller;

import cn.jxau.common.response.CommonResponse;
import cn.jxau.producer.excetion.BaseException;
import cn.jxau.producer.excetion.BizErrorEnum;
import org.springframework.web.bind.ServletRequestBindingException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.NoHandlerFoundException;

import java.util.HashMap;
import java.util.Map;

/**
 * @Author l
 * @Date 2021/4/12 14:57
 * @Version 1.0
 */
@RestControllerAdvice
public class GlobalExceptionHandler {
    @ExceptionHandler(Exception.class)
    public CommonResponse resolve(Exception ex){
       Map<String,Object> respData= new HashMap<>();
       if(ex instanceof BaseException){
           BaseException baseException = (BaseException) ex;
           respData.put("errorCode",baseException.getErrorCode());
           respData.put("errorMsg",baseException.getErrorMsg());
       }
       else if(ex instanceof ServletRequestBindingException){
            respData.put("errorCode", BizErrorEnum.UNKNOWN_ERROR.getCode());
            respData.put("errorMsg", "访问路径不存在");
       }
       else if(ex instanceof NoHandlerFoundException){
           respData.put("errorCode",BizErrorEnum.UNKNOWN_ERROR.getCode());
           respData.put("errorMsg","方法未绑定");
       }
       else{
           respData.put("errorCode",BizErrorEnum.UNKNOWN_ERROR.getCode());
           respData.put("errorMsg",BizErrorEnum.UNKNOWN_ERROR.getMsg());
       }
       ex.printStackTrace();
       return CommonResponse.create(respData,"fail");
    }
}
