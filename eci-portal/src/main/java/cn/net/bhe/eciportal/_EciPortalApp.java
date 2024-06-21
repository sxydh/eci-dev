package cn.net.bhe.eciportal;

import cn.net.bhe.eciportal.config.ResProp;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

@SpringBootApplication
@EnableConfigurationProperties({ResProp.class})
public class _EciPortalApp {

    public static void main(String[] args) {
        SpringApplication.run(_EciPortalApp.class);
    }

}