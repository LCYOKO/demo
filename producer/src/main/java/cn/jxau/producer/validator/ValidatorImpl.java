package cn.jxau.producer.validator;

import cn.jxau.common.validator.ValidationResult;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.stereotype.Component;

import javax.validation.ConstraintViolation;
import javax.validation.Path;
import javax.validation.Validation;
import javax.validation.Validator;
import java.util.Set;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 11:36 下午
 * @Version 1.0
 */
@Component
public class ValidatorImpl implements InitializingBean {
    private Validator validator;

    public ValidationResult validate(Object bean){
        ValidationResult result = new ValidationResult();
        Set<ConstraintViolation<Object>> set = validator.validate(bean);
        if(!set.isEmpty()){
            set.forEach(e->{
                String propertyPath = e.getPropertyPath().toString();
                String message = e.getMessage();
                result.getErros().put(propertyPath,message);
            });
            result.setHasErrors(true);
        }
        return result;
    }
    @Override
    public void afterPropertiesSet() throws Exception {
        this.validator= Validation.buildDefaultValidatorFactory().getValidator();
    }
}
