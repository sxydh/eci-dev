package cn.net.bhe.eciportal.model.vo;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class ContainerVO {

    private String name;
    private String image;
    private String imagePullPolicy;
    private String command;
    private String ready;
    private String restarts;
    private ResourceVO resourceRequest;
    private ResourceVO resourceLimit;

}
