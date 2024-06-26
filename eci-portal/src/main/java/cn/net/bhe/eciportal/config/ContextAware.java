package cn.net.bhe.eciportal.config;

import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.stereotype.Component;

@Component
public class ContextAware implements ApplicationContextAware {

    private static ApplicationContext applicationContext;

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        ContextAware.applicationContext = applicationContext;
    }

    public static <T> T getBean(Class<T> clazz) {
        return ContextAware.applicationContext.getBean(clazz);
    }

}
