package cn.jxau.producer.config;

import cn.jxau.producer.listener.MessageDelegate;
import com.google.common.cache.CacheStats;
import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.connection.CorrelationData;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.rabbit.listener.SimpleMessageListenerContainer;
import org.springframework.amqp.rabbit.listener.adapter.MessageListenerAdapter;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;


/**
 * @Author l
 * @Date 2021/4/2 22:40
 * @Version 1.0
 */
@Configuration
public class RabbitMqConfig {
    @Value("${rabbitmq.exchange-name}")
    private String ORDER_EXCHANGE_NAME;

    @Value("${rabbitmq.queue-name}")
    private String ORDER_QUEUE_NAME;


    @Bean("oderTopicExchange")
    public Exchange topicExchange(){
        Exchange exchange = ExchangeBuilder.topicExchange(ORDER_EXCHANGE_NAME).durable(true).build();
        return exchange;
    }

    @Bean
    public Exchange exchange(){
        return  ExchangeBuilder.directExchange("test_direct").durable(true).build();
    }

    @Bean("orderQueue")
    public Queue orderQueue(){
        return QueueBuilder.durable(ORDER_QUEUE_NAME).build();
    }

    @Bean
    public Binding binding(@Qualifier("orderQueue") Queue orderQueue,
                           @Qualifier("oderTopicExchange") Exchange exchange){

        return BindingBuilder.bind(orderQueue).to(exchange).with("order.#").noargs();
    }

    @Bean
    public RabbitTemplate rabbitTemplate(ConnectionFactory factory){
        RabbitTemplate template = new RabbitTemplate(factory);
        final RabbitTemplate.ConfirmCallback confirmCallback = new RabbitTemplate.ConfirmCallback() {
            @Override
            public void confirm(CorrelationData correlationData, boolean ack, String cause) {
                System.err.println("correlationData: " + correlationData);
                System.err.println("ack: " + ack);
                if(!ack){
                    System.err.println(cause);
                }
            }
        };

        //回调函数: return返回
        final RabbitTemplate.ReturnCallback returnCallback = new RabbitTemplate.ReturnCallback() {
            @Override
            public void returnedMessage(org.springframework.amqp.core.Message message, int replyCode, String replyText,
                                        String exchange, String routingKey) {
                System.err.println("return exchange: " + exchange + ", routingKey: "
                        + routingKey + ", replyCode: " + replyCode + ", replyText: " + replyText);
            }
        };
        template.setConfirmCallback(confirmCallback);
        template.setReturnCallback(returnCallback);
        return template;
    }

}
