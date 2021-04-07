package cn.jxau.consumer;

import org.apache.curator.RetryPolicy;
import org.apache.curator.framework.CuratorFramework;
import org.apache.curator.framework.CuratorFrameworkFactory;
import org.apache.curator.retry.ExponentialBackoffRetry;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.data.Stat;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * @Author l
 * @Date 2021/3/27 21:18
 * @Version 1.0
 */
public class TestLog {
   Logger logger=LoggerFactory.getLogger(TestLog.class);
   CuratorFramework client;
   String ip="192.168.102.130";
   @Before
   public void createClient(){
      RetryPolicy retryPolicy;
      client=CuratorFrameworkFactory.builder()
              .connectionTimeoutMs(10000)
              .connectString(ip)
              .retryPolicy(new ExponentialBackoffRetry(1000,3))
              //.namespace("/testZK")
              .build();
      client.start();
   }

   @After
   public void closeClient(){
       client.close();
   }
    @Test
    public void testLog() throws Exception{
       // 创建永久节点
        final String s = client.create()
                .creatingParentsIfNeeded()
                .withMode(CreateMode.PERSISTENT)
                .forPath("/testZK/child01", "child01".getBytes());
       // 获取节点信息
        Stat stat = new Stat();
        byte[] bytes = client.getData()
                .storingStatIn(stat)
                .forPath("/testZK/child01");
        System.out.println(new String(bytes));

        // 注册wacher
        //final byte[] bytes1 = client.getData()
        //        .usingWatcher()
        //        .forPath(s);
    }
}
