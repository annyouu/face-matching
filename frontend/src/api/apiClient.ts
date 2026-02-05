import axios from 'axios';

// axiosのインスタンスを作成
export const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8000',
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json',
    },
});

// 送信前の共通処理
apiClient.interceptors.request.use(
    (config) => {
        // Cookieなどからトークンを取り出してヘッダーに入れる
        // 全てのAPIリクエストで自動的にトークンが送られる
        const token = document.cookie
        .split('; ')
        .find((row) => row.startsWith('token='))
        ?.split('=')[1];

        if (token && config.headers) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 受信後の共通処理
apiClient.interceptors.response.use(
    (response) => response.data,
    (error) => {
        if (error.reponse?.status === 401) {
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);