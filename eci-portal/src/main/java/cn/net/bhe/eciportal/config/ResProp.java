package cn.net.bhe.eciportal.config;

import lombok.Data;
import lombok.experimental.Accessors;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Data
@Accessors(chain = true)
@ConfigurationProperties(prefix = "res")
public class ResProp {

    private String term;
    private String apiEciList;
    private String apiEci;

}
