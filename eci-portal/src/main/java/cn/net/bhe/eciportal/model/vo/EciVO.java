package cn.net.bhe.eciportal.model.vo;

import lombok.Data;
import lombok.experimental.Accessors;

import java.util.List;

@Data
@Accessors(chain = true)
public class EciVO {

    private String id;
    private String name;
    private String replicaName;
    private String ready;
    private String status;
    private String restarts;
    private String age;
    private String ip;
    private String node;
    List<ContainerVO> containers;

}
