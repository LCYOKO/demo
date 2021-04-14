package cn.jxau.producer.akka;


import akka.actor.*;
import akka.dispatch.OnComplete;
import akka.japi.pf.ReceiveBuilder;
import akka.pattern.Patterns;
import akka.util.Timeout;
import scala.Option;
import scala.concurrent.Future;
import scala.concurrent.duration.Duration;
import scala.concurrent.impl.FutureConvertersImpl;

import java.io.IOException;
import java.sql.SQLException;
import java.util.concurrent.TimeUnit;

/**
 * @Author l
 * @Date 2021/4/12 17:42
 * @Version 1.0
 */
public class ActorDemo extends AbstractActor {
    private SupervisorStrategy strategy = new OneForOneStrategy(3, java.time.Duration.ofMinutes(1), t -> {
        if (t instanceof IOException) {
            System.out.println("========自动恢复=======");
            return SupervisorStrategy.resume(); // 恢复
        } else if (t instanceof IndexOutOfBoundsException) {
            System.out.println("=========重启==========");
            return SupervisorStrategy.restart(); // 重启
        } else if (t instanceof SQLException) {
            System.out.println("==========停止=========");
            return SupervisorStrategy.stop();  // 停止
        } else {
            System.out.println("==========上报=========");
            return SupervisorStrategy.escalate(); // 上报
        }
    });

   @Override
   public SupervisorStrategy supervisorStrategy(){
        return strategy;
    }

    static class WorkerActor extends AbstractActor{
       private static int stateCount;
        @Override
        public Receive createReceive() {
            return receiveBuilder().matchAny(msg -> {
                //模拟计算任务
                //if ("add".equals(msg)) {
                //    stateCount++;
                //    System.out.println(stateCount);
                //} else if (msg.equals("get")) {
                //    System.out.println(stateCount);
                //} else if (msg instanceof Exception) {
                //    throw (Exception) msg;
                //} else {
                //    unhandled(msg);
                //}
                System.out.println(msg);
            }).build();
        }

        @Override
        public void preStart() throws Exception {
            System.out.println("启动前preStart");
            super.preStart();
        }

        @Override
        public void postStop() throws Exception {
            System.out.println("停止后postStop");
            super.postStop();
        }

        @Override
        public void preRestart(Throwable reason, Option<Object> message)
                throws Exception {
            System.out.println("重启前preRestart");
            super.preRestart(reason, message);
        }

        @Override
        public void postRestart(Throwable reason) throws Exception {
            super.postRestart(reason);
            System.out.println("重启后postRestart");
        }

    }


    @Override
    public Receive createReceive() {
        return receiveBuilder().matchEquals("stopChild",(msg)->{

        }).match(String.class,(msg)->{

        }).match(Terminated.class,(t)->{
           // System.out.println(t.getActor()+"  has stop");
        }).build();
    }

    @Override
    public void preStart() throws Exception {
       super.preStart();
        ActorRef workerActor = getContext().actorOf(Props.create(WorkerActor.class), "workerActor");
        getContext().watch(workerActor);
    }

    @Override
    public void postStop(){
           System.out.println("after stop");
    }



    public static void main(String[] args) {
            //testLifeCycle();
        ActorSystem system = ActorSystem.create("sys");
        ActorRef supervisorActor = system.actorOf(Props.create(ActorDemo.class), "supervisorActor");
        ActorSelection workerActor = system.actorSelection("akka://sys/user/supervisorActor/workerActor");
        workerActor.tell("add",ActorRef.noSender());
       // workerActor.tell(new IndexOutOfBoundsException(),ActorRef.noSender());
      //  system.terminate();
    }







    private void testBasicMethod(){
        ActorSystem system = ActorSystem.create("sys");
        ActorRef askActor = system.actorOf(Props.create(ActorDemo.class), "demo1");
        ActorRef demo2 = system.actorOf(Props.create(ActorDemo.class), "demo2");
        ActorSelection actorSelection = system.actorSelection("/user/demo*");
        //   actorSelection.tell("hello select",ActorRef.noSender());
        Future<ActorRef> actorRefFuture = actorSelection.resolveOne(new Timeout(Duration.create(1, TimeUnit.SECONDS)));
        actorRefFuture.onComplete(new OnComplete<ActorRef>() {
            @Override
            public void onComplete(Throwable failure, ActorRef success) throws Throwable {
                System.out.println(success);
            }
        },system.dispatcher());
//        System.out.println(actorSelec
//        tion.);
//        askActor.tell("213",ActorRef.noSender());
//        Timeout timeout = new Timeout(Duration.create(2, TimeUnit.SECONDS));
//        Future<Object> f = Patterns.ask(askActor, "Akka Ask", timeout);
//        f.onComplete(new OnComplete<Object>() {
//                         @Override
//                         public void onComplete(Throwable failure, Object success) throws Throwable {
//                                    System.out.println("haha"+failure.getMessage());
//                         }
//                     }
//                , system.getDispatcher());
    }

    private static void testLifeCycle(){
        ActorSystem system = ActorSystem.create("sys");
        ActorRef demo1 = system.actorOf(Props.create(ActorDemo.class), "demo1");
        demo1.tell(Kill.getInstance(),ActorRef.noSender());
    }


    private static void testStrategy(){

      //  workerActor.tell(new IndexOutOfBoundsException(), supervisorActor);


    }

}
