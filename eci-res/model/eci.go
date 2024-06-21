package model

type Eci struct {
	/* json:"id" 指定了当 Eci 结构体实例被编码成 JSON 时， Id 字段将映射到 JSON 对象中的 id 键。 */
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	ReplicaName string      `json:"replicaName"`
	Ready       string      `json:"ready"`
	Status      string      `json:"status"`
	Restarts    string      `json:"restarts"`
	Age         string      `json:"age"`
	Ip          string      `json:"ip"`
	Node        string      `json:"node"`
	Containers  []Container `json:"containers"`
}

type Container struct {
	Name            string   `json:"name"`
	Image           string   `json:"image"`
	ImagePullPolicy string   `json:"imagePullPolicy"`
	Command         string   `json:"command"`
	Ready           string   `json:"ready"`
	Restarts        string   `json:"restarts"`
	ResourceRequest Resource `json:"resourceRequest"`
	ResourceLimit   Resource `json:"resourceLimit"`
}

type Resource struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}
