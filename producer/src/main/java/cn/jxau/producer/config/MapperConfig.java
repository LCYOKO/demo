package cn.jxau.producer.config;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.context.annotation.Configuration;

/**
 * @Author l
 * @Date 2021/4/7 15:24
 * @Version 1.0
 */
@Configuration
@MapperScan("cn.jxau.producer.mapper")
public class MapperConfig {
}
