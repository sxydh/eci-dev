<template>
  <v-app>
    <v-row>
      <v-col :cols="7">
        <v-app>
          <v-img src='@/assets/material design background.jpg' cover/>
        </v-app>
      </v-col>
      <v-col>
        <v-app>
          <v-row align="center" justify="center">
            <v-col cols="6">
              <h1 style="color: #588CF9">Hello,</h1>
              <h1 style="color: #ff4f31; font-size: 50px; margin-bottom: 20px">Welcome!</h1>
              <v-form @submit.prevent="submit">
                <v-text-field v-model="username" label="Username" :rules="[value => !!value]"/>
                <v-text-field v-model="password" label="Password" :rules="[value => !!value]" type="password"/>
                <v-btn type="submit" block>Login</v-btn>
              </v-form>
            </v-col>
          </v-row>
        </v-app>
      </v-col>
    </v-row>
  </v-app>
</template>

<script setup lang="ts">
import {ref} from 'vue';
import {useRouter} from 'vue-router';

/* 变量 */
// 用户名
const username = ref<string>('admin');
// 密码
const password = ref<string>('123456');
// 路由
const router = useRouter();

/* 函数 */
const submit = async (event: any) => {
  let results = await event;
  if (results.valid && username.value === 'admin' && password.value === '123456') {
    localStorage.setItem('authToken', 'xxx');
    await router.push({path: '/'});
  }
};
</script>