export const apiClient = {
    // <>でのジェネリクス追加で、呼び出し側で型を指定できるようにする
    async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
        const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
        const url = `${BASE_URL}${endpoint}`;
        const headers = new Headers(options.headers);

        if (!(options.body instanceof FormData)) {
            headers.set("Content-Type", "application/json");
        }

        const config = { ...options, headers };

        try {
            const response = await fetch(url, config);
            if (!response.ok) {
                const errorData = await response.json().catch(() => ({}));
                throw new Error(errorData.message || "API通信エラー");
            }
            return response.status !== 204 ? await response.json() : ({} as T);
        } catch (error) {
            console.error("Network Error:", error);
            throw error;
        }
    },

    get<T>(endpoint: string, options?: RequestInit) {
        return this.request<T>(endpoint, {
            ...options,
            method: "GET"
        });
    },

    post<T>(endpoint: string, body: any, options?: RequestInit) {
        return this.request<T>(endpoint, {
            ...options,
            body: body instanceof FormData ? body : JSON.stringify(body),
            method: "POST",
        });
    },

    patch<T>(endpoint: string, body: any, options?: RequestInit) {
        return this.request<T>(endpoint, {
            ...options,
            body: body instanceof FormData ? body : JSON.stringify(body),
            method: "PATCH",
        });
    },
};