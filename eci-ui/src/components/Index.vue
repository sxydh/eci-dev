<template>
  <v-app>
    <!-- 标题栏 -->
    <v-app-bar>
      <v-app-bar-nav-icon @click="openDrawer = !openDrawer"/>
      <v-toolbar-title>Application</v-toolbar-title>
      <v-spacer/>
      <v-btn :icon="true" @click="handleLogout">
        <v-icon>mdi-export</v-icon>
      </v-btn>
    </v-app-bar>

    <!-- 导航栏 -->
    <v-navigation-drawer v-model="openDrawer">
      <v-list>
        <v-list-item
            v-for="(item, index) in menuItems"
            :key="index"
            link
        >
          <template v-slot:prepend>
            <v-icon :color="item.iconColor">{{ item.icon }}</v-icon>
          </template>
          <v-list-item-title @click.prevent="handleMenuItemClick(item)">{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <!-- 内容栏 -->
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import {ref} from 'vue';
import {useRouter} from 'vue-router';

/* 变量 */
// 菜单抽屉状态
const openDrawer = ref(true);
// 菜单列表
const menuItems = ref([
  {title: '概览', icon: 'mdi-home', iconColor: '#FF6F00', to: '/'},
  {title: '实例管理', icon: 'mdi-cloud', iconColor: '#3498DB', to: '/eciList'},
  {title: '网络配置', icon: 'mdi-network', iconColor: '#2ECC71', to: '/'},
  {title: '存储管理', icon: 'mdi-database', iconColor: '#34495E', to: '/'},
  {title: '镜像管理', icon: 'mdi-cube-send', iconColor: '#9B59B6', to: '/'},
  {title: '安全配置', icon: 'mdi-shield-account', iconColor: '#E74C3C', to: '/'},
  {title: '监控面板', icon: 'mdi-monitor', iconColor: '#F1C40F', to: '/'},
  {title: '账户管理', icon: 'mdi-account-settings', iconColor: '#3498DB', to: '/'},
  {title: '开发帮助', icon: 'mdi-code-tags', iconColor: '#7F8C8D', to: '/'},
]);
// 路由
const router = useRouter();

/* 函数 */
const handleMenuItemClick = (item: any) => {
  router.push({
    path: item.to as string,
  });
};
const handleLogout = () => {
  localStorage.removeItem('authToken');
  router.push({path: '/login'});
};
</script>