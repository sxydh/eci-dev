package cn.net.bhe.eciportal.model;

import lombok.Data;
import lombok.experimental.Accessors;
import org.springframework.http.HttpStatus;

@Data
@Accessors(chain = true)
public class R<X> {

    private Integer code;
    private String msg;
    private X data;

    private R() {
    }

    public static <T> R<T> ok() {
        return ok(null);
    }

    public static <T> R<T> ok(T data) {
        return new R<T>()
                .setCode(HttpStatus.OK.value())
                .setMsg(HttpStatus.OK.name())
                .setData(data);
    }

}
