package cn.jxau.producer;

import cn.jxau.producer.entity.Order;
import cn.jxau.producer.mq.RabbitSender;
import cn.jxau.producer.services.impl.HelloServiceImpl;
import cn.jxau.producer.thrift.Hello;
import org.apache.thrift.TProcessor;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TSimpleServer;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TTransportException;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

/**
 * @Author l
 * @Date 2021/3/29 13:04
 * @Version 1.0
 */
@RunWith(org.springframework.test.context.junit4.SpringRunner.class)
@SpringBootTest
public class TestFinagle {

    @Autowired
    private RabbitTemplate rabbitTemplate;

    @Test
    public void testFinagle() throws TTransportException {
        TProcessor tprocessor = new Hello.Processor<Hello.Iface>(
                new HelloServiceImpl());
        TServerSocket serverTransport = new TServerSocket(8080);
        TSimpleServer.Args tArgs = new TSimpleServer.Args(serverTransport);
        tArgs.processor(tprocessor);
        tArgs.protocolFactory(new TBinaryProtocol.Factory());

        TServer server = new TSimpleServer(tArgs);

        server.serve();

    }

    @Autowired
    private RabbitSender sender;
    @Test
    public void sendMsg() throws Exception {
        Order order = new Order();
        order.setId(1);
        order.setId(123);
        sender.sendOrder(order);
    }


    //@Autowired
    //private FoodDao foodDao;
    //@Test
    //public void decFoodAmount(){
    //    System.out.println(foodDao.selectByPrimaryKey(1));
    //}






}
