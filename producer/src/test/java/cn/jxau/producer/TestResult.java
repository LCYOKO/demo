package cn.jxau.producer;

import cn.jxau.common.response.Result;
import org.junit.Test;

import java.util.Arrays;
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
        int[] nums={1,2,3,4};
        String collect = Arrays.stream(nums).mapToObj(e -> "" + e).sorted((o1, o2) -> {
            String num1 = o1 + o2;
            String num2 = o2 + o1;
            return num2.compareTo(num1);
        }).collect(Collectors.joining());
        System.out.println(collect);
    }

}
