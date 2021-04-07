package cn.jxau.producer.mq;

import cn.jxau.producer.entity.Order;

import org.springframework.amqp.AmqpException;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessagePostProcessor;
import org.springframework.amqp.rabbit.connection.CorrelationData;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;

import org.springframework.stereotype.Component;

import java.util.Map;

/**
 * @Author l
 * @Date 2021/4/3 21:42
 * @Version 1.0
 */
@Component
public class RabbitSender {
    @Autowired
    private RabbitTemplate rabbitTemplate;

    public void send(Object message,Map<String,Object> headers) throws Exception {


        //id + 时间戳 全局唯一
        CorrelationData correlationData = new CorrelationData("1234567890");
        rabbitTemplate.convertAndSend("topic.exchange", "order.abc", message, correlationData);
    }

    public void sendOrder(Order order) throws Exception {
        //id + 时间戳 全局唯一
        CorrelationData correlationData = new CorrelationData("0987654321");

        rabbitTemplate.convertAndSend("topic.exchange", "order.123",order, new MessagePostProcessor(){

            @Override
            public Message postProcessMessage(Message message) throws AmqpException {
             //   message.getMessageProperties().setExpiration("100");
                System.err.println(message);
                return message;
            }
        },correlationData);
    }

}
