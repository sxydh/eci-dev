package cn.net.bhe.eciportal.controller;

import cn.net.bhe.eciportal.config.ContextAware;
import cn.net.bhe.eciportal.config.ResProp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.util.UriComponentsBuilder;

import javax.websocket.*;
import javax.websocket.server.ServerEndpoint;
import java.io.IOException;
import java.net.URI;
import java.util.List;
import java.util.Map;

/**
 * 每个会话会实例化一个 ContainerServerEndpoint
 */
@Slf4j
@ServerEndpoint("/term")
public class ContainerServerEndpoint {

    private Session session;
    private ContainerClientEndpoint exchangeEndpoint;

    @OnOpen
    public void onOpen(Session session) {
        log.debug("onOpen: getRequestURI={}, getRequestParameterMap={}", session.getRequestURI(), session.getRequestParameterMap());
        this.session = session;
        /* 本地服务作为客户端，建立与 res 服务端的连接。 */
        WebSocketContainer webSocketContainer = ContainerProvider.getWebSocketContainer();
        exchangeEndpoint = new ContainerClientEndpoint(this);
        try {
            ResProp resProp = ContextAware.getBean(ResProp.class);
            Map<String, List<String>> requestParameterMap = session.getRequestParameterMap();
            URI uri = UriComponentsBuilder.fromUriString(resProp.getTerm())
                    .queryParam("type", requestParameterMap.get("type"))
                    .queryParam("name", requestParameterMap.get("name"))
                    .queryParam("eciName", requestParameterMap.get("eciName"))
                    .build()
                    .toUri();
            // noinspection resource
            webSocketContainer.connectToServer(exchangeEndpoint, uri);
        } catch (Exception e) {
            log.error(e.getLocalizedMessage(), e);
        }
    }

    /**
     * 本地服务作为服务端，收到浏览器客户端消息。
     */
    @OnMessage
    public void onMessage(String message) {
        log.debug("onMessage: message={}", message);
        /* 本地服务作为客户端，将消息转发 res 服务端。 */
        exchangeEndpoint.exchange(message);
    }

    @OnClose
    public void onClose(CloseReason closeReason) throws IOException {
        log.debug("onClose: getCloseCode={}, getReasonPhrase={}", closeReason.getCloseCode(), closeReason.getReasonPhrase());
        exchangeEndpoint.exchangeClose(closeReason);
    }

    /**
     * 本地服务作为服务端，将消息转发浏览器客户端。
     */
    public void exchange(String message) {
        session.getAsyncRemote().sendText(message);
    }

    public void exchangeClose(CloseReason closeReason) throws IOException {
        this.session.close(closeReason);
    }

}
