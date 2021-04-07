package cn.jxau.producer.listener;

import com.rabbitmq.client.Channel;
import org.springframework.amqp.core.Message;



/**
 * @Author l
 * @Date 2021/4/3 20:00
 * @Version 1.0
 */
public class MessageDelegate {
    public void handleMessage(byte[] messageBody) {
        System.err.println("默认方法, 消息内容:" + new String(messageBody));
    }
}
