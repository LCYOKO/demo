package cn.jxau.producer;

import cn.jxau.common.response.Result;
import org.junit.Test;

/**
 * @Author l
 * @Date 2021/4/7 16:12
 * @Version 1.0
 */
public class TestResult {


    @Test
    public  void test01(){
        Result<Object> result = Result.fail("123");
    }
}
