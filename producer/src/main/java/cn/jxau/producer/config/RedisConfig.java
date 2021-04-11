package cn.jxau.producer.config;

import cn.jxau.producer.entity.Order;
import cn.jxau.producer.entity.User;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;

/**
 * @Author l
 * @Date 2021/3/30 15:27
 * @Version 1.0
 */
@Configuration
public class RedisConfig {

   @Bean
   public RedisTemplate<String, User> redisTemplate(RedisConnectionFactory factory){
       RedisTemplate<String, User> template = new RedisTemplate<>();
       template.setValueSerializer(new Jackson2JsonRedisSerializer<>(User.class));
       template.setKeySerializer(new StringRedisSerializer());
       template.setConnectionFactory(factory);
       return template;
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
