package cn.jxau.producer.services;

import cn.jxau.common.api.OrderService;
import org.apache.dubbo.config.annotation.Service;

/**
 * @Author l
 * @Date 2021/3/28 17:43
 * @Version 1.0
 */
@Service
public class OrderServiceImp implements OrderService {
    @Override
    public String getOrder(Integer id) {
        return "hahah";
    }
}
