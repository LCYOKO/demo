package cn.jxau.producer.entity.vo;

import lombok.Data;

import javax.validation.constraints.Max;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 10:23 下午
 * @Version 1.0
 */
@Data
public class UserVO {
  @Max(value = 10,message = "id 最大就是10")
  private Integer id;
  private String name;
  private Byte gender;
  private Integer age;
  private String telPhone;
  private String registerMode;
  private String thirdPartyId;
  private String encrptPassword;

}
