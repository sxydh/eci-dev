<template>
  <v-card>
    <v-card-actions>
      <v-row align="center" justify="start">
        <v-col :cols="2">
          <v-text-field
              v-model="eciFilter"
              prepend-icon="mdi-magnify"
              variant="solo"
              density="comfortable"
              hide-details
              clearable
          />
        </v-col>
        <v-col>
          <v-btn variant="tonal" density="comfortable" icon="mdi-reload" @click.prevent="queryEciList"/>
          <v-btn variant="tonal" density="comfortable" icon="mdi-plus" @click.prevent="handleOpenAdd"/>
        </v-col>
      </v-row>
    </v-card-actions>

    <v-data-table :headers="headers" :items="eciList">
      <template v-slot:item.name="{item}">
        <a href="#" @click.prevent="routeToContainerList(item)">{{ item.name }}</a>
      </template>
      <template v-slot:item.operation="{item}">
        <v-row>
          <v-col>
            <v-btn :icon="true" density="comfortable" @click.prevent="handleDelete(item.name)">
              <v-icon :color="'#ff3a3a'">mdi-delete</v-icon>
            </v-btn>
          </v-col>
        </v-row>
      </template>
    </v-data-table>

    <v-dialog
        v-model="openAdd"
        max-width="600"
        transition="dialog-bottom-transition"
    >
      <v-card>
        <v-card-title>
          <v-row align="center">
            <v-col>New Instance</v-col>
            <v-col class="d-flex justify-end">
              <v-btn icon="mdi-close" @click="handleCloseAdd"/>
            </v-col>
          </v-row>
        </v-card-title>
        <v-card-item>
          <v-form @submit.prevent="handleAdd">
            <v-text-field
                v-model="eciForm.name"
                label="Instance name"
                :rules="[value => !!value]"
            />
            <v-select
                v-model="eciForm.image"
                label="Select image"
                :items="['tomcat:latest']"
                :rules="[value => !!value]"
            />
            <v-select
                v-model="eciForm.cpu"
                label="Cpu size"
                :items="['100m', '200m', '500m', '1']"
                :rules="[value => !!value]"
            />
            <v-select
                v-model="eciForm.memory"
                label="Memory size"
                :items="['128Mi', '256Mi', '512Mi']"
                :rules="[value => !!value]"
            />
            <v-btn type="submit" block>Submit</v-btn>
          </v-form>
        </v-card-item>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import {addEci, deleteEci, getEciList} from '@/api/eci.api';
import {Eci} from '@/model/Eci';
import {EciDTO} from '@/model/EciDTO';
import {useRouter} from 'vue-router';

/* 变量 */
// 表格头
const headers = [
  {title: 'Name', value: 'name'},
  {title: 'Operation', value: 'operation'},
  {title: 'Ready', value: 'ready'},
  {title: 'Status', value: 'status'},
  {title: 'Restarts', value: 'restarts'},
  {title: 'Age', value: 'age'},
  {title: 'IP', value: 'ip'},
  {title: 'Node', value: 'node'},
];
// ECI 列表
const _eciList = ref<Eci[]>([]);
const eciList = ref<Eci[]>([]);
// ECI 过滤
const eciFilter = ref<string>();
// ECI 表单
const eciForm = ref<EciDTO>({
  name: '',
  image: '',
  cpu: '',
  memory: ''
});
// 路由
const router = useRouter();
// 打开新增对话框
const openAdd = ref<boolean>();

/* 生命周期回调 */
// 组件渲染到 DOM 时回调
onMounted(() => {
  queryEciList();
});

/* 响应式监听 */
watch(eciFilter, () => {
  filterEciList();
});

/* 函数 */
// 获取 ECI 列表
const queryEciList = () => {
  getEciList().then(res => {
    _eciList.value = res.data;
    filterEciList();
  });
};
// 过滤 ECI 列表
const filterEciList = () => {
  eciList.value = _eciList.value.filter((ele: Eci) => {
    return ele.name.indexOf(eciFilter.value || '') >= 0;
  });
};
// 跳转 Container 列表
const routeToContainerList = (item: Eci) => {
  localStorage.setItem(`${item.name}.containers`, JSON.stringify(item.containers));
  router.push({
    path: '/containerList',
    query: {
      name: item.name,
      replicaName: item.replicaName
    },
  });
};
// 打开新增对话框
const handleOpenAdd = () => {
  openAdd.value = true;
  eciForm.value.name = Date.now().toString();
};
// 关闭新增对话框
const handleCloseAdd = () => {
  openAdd.value = false;
  clearEciForm();
};
// 新增
const handleAdd = async (event: any) => {
  let results = await event;
  if (results.valid) {
    addEci(eciForm.value).then(() => {
      queryEciList();
    }).finally(() => {
      openAdd.value = false;
      clearEciForm();
    });
  }
};
// 删除
const handleDelete = async (name: string) => {
  deleteEci(name).then(() => {
    queryEciList();
  });
};
// 清空表单
const clearEciForm = () => {
  eciForm.value = {
    name: '',
    image: '',
    cpu: '',
    memory: ''
  };
};
</script>