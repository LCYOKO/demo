package cn.jxau.producer.entity;

import java.io.Serializable;
import lombok.Data;

/**
 * food
 * @author 
 */
@Data
public class Food implements Serializable {
    private Integer id;

    private String name;

    private Integer cid;

    private String description;

    private Integer num;

    private Integer sale;

    private String imgurl;

    private Integer state;

    private Double price;

    private Integer version;

    private static final long serialVersionUID = 1L;
}