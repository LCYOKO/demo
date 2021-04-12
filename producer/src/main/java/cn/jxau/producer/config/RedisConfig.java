package cn.jxau.producer.config;

import cn.jxau.producer.entity.Order;
import cn.jxau.producer.entity.User;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.module.SimpleModule;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;

import java.util.Date;

/**
 * @Author l
 * @Date 2021/3/30 15:27
 * @Version 1.0
 */
@Configuration
public class RedisConfig {

    @Bean
    public RedisTemplate redisTemplate(RedisConnectionFactory factory){
        RedisTemplate redisTemplate = new RedisTemplate();
        redisTemplate.setConnectionFactory(factory);

        //首先解决key的序列化方式
        StringRedisSerializer stringRedisSerializer = new StringRedisSerializer();
        redisTemplate.setKeySerializer(stringRedisSerializer);

        //解决value的序列化方式
        Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new Jackson2JsonRedisSerializer(Object.class);

        ObjectMapper objectMapper =  new ObjectMapper();
        SimpleModule simpleModule = new SimpleModule();
        //simpleModule.addSerializer(Date.class,new DateTimeSe);
        //simpleModule.addDeserializer(Date.class,new JodaDateTimeJsonDeserializer());
        objectMapper.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
        objectMapper.registerModule(simpleModule);

        jackson2JsonRedisSerializer.setObjectMapper(objectMapper);

        redisTemplate.setValueSerializer(jackson2JsonRedisSerializer);

        return redisTemplate;
    }

   @Bean
   public RedisTemplate<String, Order> orderRedisTemplate(RedisConnectionFactory factory){
       RedisTemplate<String, Order> orderRedisTemplate= new RedisTemplate<>();
       orderRedisTemplate.setValueSerializer(new Jackson2JsonRedisSerializer<>(Order.class));
       orderRedisTemplate.setKeySerializer(new StringRedisSerializer());
       orderRedisTemplate.setConnectionFactory(factory);
       return orderRedisTemplate;
   }
}
