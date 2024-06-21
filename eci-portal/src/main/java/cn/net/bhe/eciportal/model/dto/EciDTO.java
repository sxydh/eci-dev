package cn.net.bhe.eciportal.model.dto;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class EciDTO {

    private String name;
    private String image;
    private String cpu;
    private String memory;

}
