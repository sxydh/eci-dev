package cn.net.bhe.eciportal.controller;

import lombok.extern.slf4j.Slf4j;

import javax.websocket.*;
import java.io.IOException;

/**
 * 每个会话会实例化一个 ContainerClientEndpoint
 */
@Slf4j
@ClientEndpoint
public class ContainerClientEndpoint {

    private Session session;
    private final ContainerServerEndpoint exchangeEndpoint;

    public ContainerClientEndpoint(ContainerServerEndpoint exchangeEndpoint) {
        this.exchangeEndpoint = exchangeEndpoint;
    }

    @OnOpen
    public void onOpen(Session session) {
        log.debug("onOpen: getRequestURI={}, getRequestParameterMap={}", session.getRequestURI(), session.getRequestParameterMap());
        this.session = session;
    }

    /**
     * 本地服务作为客户端，收到 res 服务端消息。
     */
    @OnMessage
    public void onMessage(String message) {
        log.debug("onMessage: message={}", message);
        /* 本地服务作为服务端，将消息转发浏览器客户端。 */
        exchangeEndpoint.exchange(message);
    }

    @OnClose
    public void onClose(CloseReason closeReason) throws IOException {
        log.debug("onClose: getCloseCode={}, getReasonPhrase={}", closeReason.getCloseCode(), closeReason.getReasonPhrase());
        exchangeEndpoint.exchangeClose(closeReason);
    }

    /**
     * 本地服务作为客户端，将消息转发 res 服务端。
     */
    public void exchange(String message) {
        session.getAsyncRemote().sendText(message);
    }

    public void exchangeClose(CloseReason closeReason) throws IOException {
        this.session.close(closeReason);
    }

}
