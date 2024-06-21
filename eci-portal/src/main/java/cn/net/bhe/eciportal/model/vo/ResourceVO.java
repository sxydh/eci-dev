package cn.net.bhe.eciportal.model.vo;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class ResourceVO {

    private String cpu;
    private String memory;

}
