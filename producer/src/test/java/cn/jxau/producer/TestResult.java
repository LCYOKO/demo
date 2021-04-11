package cn.jxau.producer;

import cn.jxau.producer.entity.Result;
import org.junit.Test;

import java.util.ArrayList;

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
}
