package controller

import (
	"context"
	"eci-res/service"
	"eci-res/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

type EciController struct {
	eciService *service.EciService
}

func NewEciController() *EciController {
	return &EciController{
		eciService: &service.EciService{},
	}
}

func (ctrl *EciController) ListEci(c *gin.Context) {
	eciList, err := ctrl.eciService.QueryEciList()
	if err != nil {
		/* gin.H 可以创建哈希字典 */
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, eciList)
}

func (ctrl *EciController) AddEci(c *gin.Context) {
	err := ctrl.eciService.AddEci(c)
	if err != nil {
		/* gin.H 可以创建哈希字典 */
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (ctrl *EciController) DeleteEci(c *gin.Context) {
	err := ctrl.eciService.DeleteEci(c)
	if err != nil {
		/* gin.H 可以创建哈希字典 */
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (ctrl *EciController) ContainerWsHandler(c *gin.Context) {
	/* 获取请求参数 */
	containerName := c.Query("name")
	podName := c.Query("eciName")

	/* 进入容器终端 */
	option := &corev1.PodExecOptions{
		Container: containerName,
		Command:   []string{"bash"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}
	k8sClient := util.GetK8sClient()
	req := k8sClient.CoreV1().RESTClient().Post().Resource("pods").
		Namespace("namespace-demo").
		Name(podName).
		SubResource("exec").
		VersionedParams(option, scheme.ParameterCodec)
	k8sConfig := util.GetK8sConfig()
	exec, err := remotecommand.NewSPDYExecutor(k8sConfig, "POST", req.URL())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"remotecommand.NewSPDYExecutor": err})
		return
	}

	// 将标准的 HTTP 请求升级为 WebSocket 连接
	upgrader := websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"upgrader.Upgrade": err})
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer wsConn.Close()
	wsShellClient := &WsShellClient{conn: wsConn}

	// 将 WebSocket 输入输出接入到容器终端
	err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdin:  wsShellClient,
		Stdout: wsShellClient,
		Stderr: wsShellClient,
		Tty:    true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executor.StreamWithContext": err})
		return
	}
}

type WsShellClient struct {
	conn *websocket.Conn
}

func (wsc *WsShellClient) Write(p []byte) (n int, err error) {
	err = wsc.conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (wsc *WsShellClient) Read(p []byte) (n int, err error) {
	_, b, err := wsc.conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, b), nil
}
