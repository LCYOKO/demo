package cn.jxau.common.validator;

import com.google.common.base.Strings;
import org.apache.commons.lang3.StringUtils;

import java.util.HashMap;
import java.util.Map;

/**
 * @Author liuchiyun
 * @Date 2021/4/11 11:32 下午
 * @Version 1.0
 */
public class ValidationResult {
    private boolean hasErrors;
    private Map<String, String> erros;
    private Character DEFAULT_DELIMITER=',';
    public ValidationResult() {
        hasErrors = false;
        erros = new HashMap<>();
    }

    public void setHasErrors(boolean hasErrors) {
        this.hasErrors = hasErrors;
    }

    public boolean isHasErrors() {
        return hasErrors;
    }

    public void setErros(Map<String, String> erros) {
        this.erros = erros;
    }

    public Map<String, String> getErros() {
        return erros;
    }

    public String getErrorsMsg(){
        return StringUtils.join(erros.values().toArray(),DEFAULT_DELIMITER);
    }
}
