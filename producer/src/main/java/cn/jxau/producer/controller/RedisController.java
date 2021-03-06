package cn.jxau.producer.controller;

import cn.jxau.producer.entity.Order;
import cn.jxau.common.response.Result;
import cn.jxau.producer.entity.User;
import cn.jxau.producer.services.FoodService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


/**
 * @Author l
 * @Date 2021/3/29 16:54
 * @Version 1.0
 */
@RestController
@RequestMapping("/redis")
public class RedisController {
    @Autowired
    private RedisTemplate<String, User> redisTemplate;
    @Autowired
    private RedisTemplate<String, Order> orderRedisTemplate;
    //@Autowired
    //private FoodService foodService;

    @GetMapping("/get")
    public User getVal(String id) {
        ValueOperations<String, User> ops = redisTemplate.opsForValue();
        return ops.get(id);
    }


    @GetMapping("/order/get")
    public Order getOrder(String id) {
        ValueOperations<String, Order> ops = orderRedisTemplate.opsForValue();
        return ops.get(id);
    }

    @GetMapping("/order/set")
    public String setOrder(Integer id) {
        ValueOperations<String, Order> ops = orderRedisTemplate.opsForValue();
        ops.set(id.toString(), buildOrder(id));
        return "ok";
    }


    @GetMapping("/set")
    public String setVal(Integer id) {
        ValueOperations<String, User> ops = redisTemplate.opsForValue();
        ops.set(id.toString(), buildUser(id));
        return "ok";
    }

    @GetMapping("/nginx")
    public String testNginx() {
        return "8083";
    }


    //@GetMapping("/get1")
    //public Result decFood(Integer id, Integer amount) {
    //    int update = foodService.decFoodById(id, amount);
    //    if(update==0){
    //        return Result.fail("扣减库存失败");
    //    }
    //    return Result.of("扣减库存成功");
    //}

    ;

    private Order buildOrder(Integer id) {
        Order order = new Order();
        order.setId(id);
        order.setOrderId("asd1111");
        return order;
    }


    private User buildUser(Integer id) {
        User user = new User();
        user.setId(id);
        user.setName("haha");
        return user;
    }
}
