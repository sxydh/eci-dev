<template>
  <v-card>
    <v-card-actions>
      <v-btn :icon="true" @click="handleBack">
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>
      <v-card-title>{{ eciName }}</v-card-title>
    </v-card-actions>

    <v-card variant="plain">
      <v-row>
        <v-col :cols="2">
          <v-card-item subtitle="CPU">{{ eciCpu }}</v-card-item>
        </v-col>
        <v-col :cols="2">
          <v-card-item subtitle="Memory">{{ eciMemory }}</v-card-item>
        </v-col>
      </v-row>
    </v-card>

    <v-data-table :headers="headers" :items="containerList">
      <template v-slot:item.name="{item}">
        <a href="#" @click.prevent="handleOpenTerminal(item)">{{ item.name }}</a>
      </template>
    </v-data-table>

    <!-- 终端对话框 -->
    <v-dialog v-model="openTerminal" transition="dialog-bottom-transition" fullscreen>
      <v-card>
        <v-toolbar>
          <v-btn icon="mdi-arrow-left" @click="handleCloseTerminal()"/>
          <v-toolbar-title>{{ container?.name}}</v-toolbar-title>
        </v-toolbar>
        <div ref="terminalRef" style="width: 100%; height: 100%;"/>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script setup lang="ts">
import {nextTick, onMounted, ref} from 'vue';
import {Container} from '@/model/Container';
import {useRoute, useRouter} from 'vue-router';
import {Terminal} from '@xterm/xterm';
import {FitAddon} from '@xterm/addon-fit';
import {AttachAddon} from '@xterm/addon-attach';
import '@xterm/xterm/css/xterm.css';

/* 变量 */
// 表格头
const headers = [
  {title: 'Name', value: 'name'},
  {title: 'Image', value: 'image'},
  {title: 'ImagePullPolicy', value: 'imagePullPolicy'},
  {title: 'Command', value: 'command'},
  {title: 'Ready', value: 'ready'},
  {title: 'Restarts', value: 'restarts'},
];
// Container 列表
const containerList = ref<Container[]>([]);
// 路由
const route = useRoute();
const router = useRouter();
// 是否打开终端
const openTerminal = ref(false);
const terminalRef = ref(null);
// 当前容器
const container = ref<Container>();
// 当前 ECI
const eciName = ref<string>();
const eciReplicaName = ref<string>();
const eciCpu = ref<string>();
const eciMemory = ref<string>();
// 当前 WebSocket
const webSocket = ref<WebSocket>();

/* 生命周期回调 */
// 组件渲染到 DOM 时回调
onMounted(() => {
  queryContainerList();
});

/* 函数 */
// 获取 Container 列表
const queryContainerList = () => {
  eciName.value = route.query.name as string;
  eciReplicaName.value = route.query.replicaName as string;
  let containers = localStorage.getItem(`${eciName.value}.containers`);
  if (containers) {
    containerList.value = JSON.parse(containers) as Container[];
    eciCpu.value = containerList.value[0]?.resourceRequest?.cpu;
    eciMemory.value = containerList.value[0]?.resourceRequest?.memory;
  }
};
// 返回上一页面
const handleBack = () => {
  router.back();
};
// 打开 Container 终端
const handleOpenTerminal = (item: Container) => {
  container.value = item;
  openTerminal.value = true;
  nextTick(() => {
    if (terminalRef.value) {
      const terminal = new Terminal({
        cursorBlink: true,
      });
      const fitAddon = new FitAddon();
      terminal.loadAddon(fitAddon);

      webSocket.value = new WebSocket(`ws://${import.meta.env.VITE_API_SERVER}/term?type=container&name=${container.value?.name}&eciName=${eciReplicaName.value}`);
      const attachAddon = new AttachAddon(webSocket.value);
      terminal.loadAddon(attachAddon);

      terminal.open(terminalRef.value);
      fitAddon.fit();
      terminal.writeln(`Welcome ${new Date().toDateString()}`);
      terminal.focus();
    }
  });
};
// 关闭 Container 终端
const handleCloseTerminal = () => {
  openTerminal.value = false;
  webSocket.value?.close();
};
</script>