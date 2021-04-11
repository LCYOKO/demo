package cn.jxau.common.response;

/**
 * @Author l
 * @Date 2021/4/11 11:42
 * @Version 1.0
 */
public class CommonResponse {
    private String status;

    //若status=success,则data内返回前端需要的json数据
    //若status=fail，则data内使用通用的错误码格式
    private Object data;

    //定义一个通用的创建方法
    public static CommonResponse create(Object result){
        return CommonResponse.create(result,"success");
    }

    public static CommonResponse create(Object result,String status){
        CommonResponse type = new CommonResponse();
        type.setStatus(status);
        type.setData(result);
        return type;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
