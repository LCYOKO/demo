package cn.jxau.producer.listener;

import cn.jxau.producer.entity.Order;
import com.rabbitmq.client.Channel;
import org.springframework.amqp.rabbit.annotation.*;
import org.springframework.amqp.support.AmqpHeaders;
import org.springframework.messaging.MessageHeaders;
import org.springframework.messaging.handler.annotation.Headers;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.util.Map;

/**
 * @Author l
 * @Date 2021/4/3 22:57
 * @Version 1.0
 */
@Component
public class MessageListener {

   @RabbitListener(bindings = @QueueBinding(
    value=@Queue(value = "${rabbitmq.queue-name}",durable ="true"),
    exchange = @Exchange(value = "${rabbitmq.exchange-name}",
    type = "${rabbitmq.exchange-type}",
    durable = "true",
    autoDelete = "false",
    ignoreDeclarationExceptions ="true"
    ),
    key ="${rabbitmq.topic-name}"
   ))
   @RabbitHandler
   public void handleMessage(@Payload Order order, @Headers Map<String, Object> headers, Channel ch) throws IOException {
        System.out.println(order);
        ch.basicAck((long)headers.get(AmqpHeaders.DELIVERY_TAG),false);
   }
}
