package cn.jxau.producer.thrift;

import com.twitter.util.Function;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import scala.runtime.BoxedUnit;

/**
 * @Author l
 * @Date 2021/3/29 15:24
 * @Version 1.0
 */
public class ThriftClient {
    public static void main(String[] args) {
     RedisTemplate redisTemplate = new RedisTemplate();
        //Hello.FutureIface client = Thrift.client().newIface("localhost:8080", Hello.FutureIface.class);
        //Future<String> response = client.hi().onSuccess(new Function<String, BoxedUnit>() {
        //    @Override
        //    public BoxedUnit apply(String response) {
        //        System.out.println("Received response: " + response);
        //        return null;
        //    }
        //});
    }
}
