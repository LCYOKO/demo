package cn.jxau.producer.entity;

import java.io.Serializable;
import lombok.Data;

/**
 * user_info
 * @author 
 */
@Data
public class UserInfo implements Serializable {
    private Integer id;

    private String name;

    /**
     * //1代表男性，2代表女性
     */
    private Byte gender;

    private Integer age;

    private String telphone;

    /**
     * //byphone,bywechat,byalipay
     */
    private String registerMode;

    private String thirdPartyId;

    private static final long serialVersionUID = 1L;
}