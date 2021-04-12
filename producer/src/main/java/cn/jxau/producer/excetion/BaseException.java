package cn.jxau.producer.excetion;

import cn.jxau.producer.entity.User;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 11:44 下午
 * @Version 1.0
 */
public class BaseException extends Exception {
    private BaseError baseError;
    public BaseException(BaseError baseError){
        super();
        this.baseError=baseError;
    }

    public BaseException(BaseError baseError,String msg){
        super();
        this.baseError=baseError;
        this.baseError.setMsg(msg);
    }

    public Integer getErrorCode(){
        return baseError.getCode();
    }

    public String getErrorMsg(){
        return baseError.getMsg();
    }


}
