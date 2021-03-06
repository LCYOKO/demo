package cn.jxau.producer.excetion;

/**
 * @Author l
 * @Date 2021/4/10 20:39
 * @Version 1.0
 */
public enum  BizErrorEnum  implements BaseError{

    //通用错误类型10001
    PARAMETER_VALIDATION_ERROR(10001, "参数不合法"),
    UNKNOWN_ERROR(10002, "未知错误"),

    //20000开头为用户信息相关错误定义
    USER_NOT_EXIST(20001, "用户不存在"),
    USER_LOGIN_FAIL(20002, "用户手机号或密码不正确"),
    USER_NOT_LOGIN(20003, "用户还未登陆"),
    USER_REGISTER_FAIL(20004,"注册失败"),

    //30000开头为交易信息错误定义
    STOCK_NOT_ENOUGH(30001, "库存不足"),
    MQ_SEND_FAIL(30002, "库存异步消息失败"),
    RATELIMIT(30003, "活动太火爆，请稍后再试");





    private String msg;
    private Integer code;

    private BizErrorEnum(Integer code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    @Override
    public Integer getCode() {
        return code;
    }

    @Override
    public String getMsg() {
        return msg;
    }

    @Override
    public void setMsg(String msg) {
        this.msg=msg;
    }
}
