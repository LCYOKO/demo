package cn.jxau.common.validator;

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
}
