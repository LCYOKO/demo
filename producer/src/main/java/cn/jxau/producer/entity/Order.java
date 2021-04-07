package cn.jxau.producer.entity;

import lombok.Data;

import java.io.Serializable;

/**
 * @Author l
 * @Date 2021/3/30 15:46
 * @Version 1.0
 */
@Data
public class Order implements Serializable {
    private Integer id;
    private String orderId;
}
