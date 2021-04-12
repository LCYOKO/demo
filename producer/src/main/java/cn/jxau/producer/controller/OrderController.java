package cn.jxau.producer.controller;

import cn.jxau.producer.entity.Order;
import cn.jxau.producer.mq.RabbitSender;

import com.sun.org.apache.bcel.internal.generic.NEW;
import org.springframework.amqp.core.MessageProperties;
import org.springframework.amqp.rabbit.connection.CorrelationData;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.MessageHeaders;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

/**
 * @Author l
 * @Date 2021/4/2 22:52
 * @Version 1.0
 */
@RestController
@RequestMapping("/order")
public class OrderController {
    @Autowired
    private RabbitSender rabbitSender;
    //@GetMapping("/send")
    //public String sendMessage(String msg){
    //    Map<String, Object> headers = new HashMap<>();
    //    try {
    //        headers.put("x-message-ttl",1000);
    //        rabbitSender.sendOrder(buildOrder());
    //    }catch (Exception e){
    //
    //    }
    //
    //    return "ok";
    //}
    //
    //private Order buildOrder(){
    //    Order order = new Order();
    //    order.setId(1);
    //    order.setId(123);
    //    return order;
    //}
}
