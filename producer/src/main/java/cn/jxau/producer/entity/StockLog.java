package cn.jxau.producer.entity;

import java.io.Serializable;
import lombok.Data;

/**
 * item
 * @author 
 */
@Data
public class StockLog implements Serializable {
    private Integer id;

    private String title;

    private Double price;

    private String description;

    private Integer sales;

    private String imgUrl;

    private static final long serialVersionUID = 1L;
}