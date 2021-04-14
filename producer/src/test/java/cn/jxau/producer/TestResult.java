package cn.jxau.producer;

import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;
import akka.dispatch.OnComplete;
import akka.dispatch.OnSuccess;
import akka.pattern.Patterns;
import akka.util.Timeout;
import cn.jxau.common.response.Result;
import cn.jxau.producer.akka.ActorDemo;
import org.junit.Test;
import scala.concurrent.Future;
import scala.concurrent.duration.Duration;

import java.util.Arrays;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

/**
 * @Author l
 * @Date 2021/4/7 16:12
 * @Version 1.0
 */
public class TestResult {


    @Test
    public void test01() {
        Result<Object> result = testDebug(Result.fail("123"));
    }

    private Result testDebug(Result<Object> result) {

        return result;
    }

    @Test
    public void test(){
    }

}
