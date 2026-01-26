import { apiClient } from "@/lib/apiClient";
import { UserResponse } from "@/type/user";

export const setupEndpoints = {
  // 名前を設定する
  updateName: (name: string) =>
    apiClient.patch<UserResponse>("/users/setup/name", { name }),

  // 画像設定
  updateImage: (formData: FormData) =>
    apiClient.patch<UserResponse>("/users/setup/image", formData, {
      headers: {
        // ファイル送信にはこのヘッダーが必要
        "Content-Type": "multipart/form-data", 
      },
    }),
};