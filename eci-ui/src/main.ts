/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import {registerPlugins} from '@/plugins';

// Components
import App from './App.vue';

/* createApp是 Vue3 中的一个函数，用于创建一个新的 Vue 应用程序实例。 */
import {createApp} from 'vue';

import router from './router';

const app = createApp(App);

registerPlugins(app);

app
    /* 注册路由组件 */
    .use(router)
    /* 将 app 应用程序挂载（即渲染）到 DOM 中具有 id 为 app 的元素上 */
    .mount('#app');
