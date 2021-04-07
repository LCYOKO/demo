package cn.jxau.producer.entity;

import lombok.Data;

import java.sql.ResultSet;

/**
 * @Author l
 * @Date 2021/4/7 15:55
 * @Version 1.0
 */
@Data
public class Result<E> {
  private String msg;
  private Integer code;
  private E data;
  public static <T> Result<T> fail(String msg){
      return new Result();
  }

  public static <T> Result<T> fail(String msg,T data){
          return    new Result<T>().setData(data).setCode(500).setMsg(msg);
  }

  public static <T> Result<T> of(T data){
      return    new Result<T>().setData(data).setCode(500).setMsg("500错误");
  }

  public static <T> Result<T> of(String msg,T data){
      return new Result<T>().setMsg(msg).setData(data).setCode(200);
  }

  public  Result<E> setMsg(String msg){
      this.msg=msg;
      return this;
  }

  public Result<E> setCode(int code){
      this.code=code;
      return this;
  }

  public Result<E> setData(E data){
      this.data=data;
      return this;
  }


  public E getData(){
      return data;
  }

}
