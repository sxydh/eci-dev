import axios from 'axios';

/* 创建 axios 实例 */
const request = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    timeout: 1000 * 30,
});

/* 请求拦截器 */
request.interceptors.request.use(
    (config) => {
        /* localStorage 同源下是共享的 */
        const token = localStorage.getItem('authToken');
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

/* 响应拦截器 */
request.interceptors.response.use(
    (response) => {
        return response.data;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default request;
