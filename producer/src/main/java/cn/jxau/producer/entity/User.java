package cn.jxau.producer.entity;

import lombok.Data;

import javax.print.Doc;
import javax.print.DocPrintJob;
import java.util.List;

/**
 * @Author liuchiyun
 * @Date 2021/4/8 9:43 下午
 * @Version 1.0
 */
@Data
public class User extends Object {

  private String name;
  private Integer id;
}
