package cn.jxau.producer.services.impl;

import cn.jxau.producer.thrift.Hello;
import org.apache.thrift.TException;

/**
 * @Author l
 * @Date 2021/3/29 21:59
 * @Version 1.0
 */
public class HelloServiceImpl implements Hello.Iface {
    @Override
    public String hi() throws TException {
        return "123";
    }
}
