import { apiClient } from "@/lib/apiClient";
import { UserResponse } from "@/type/user";

export const setupEndpoints = {
  // 名前を設定する
  updateName: (name: string) =>
    apiClient.patch<UserResponse>("/api/v1/users/setup/name", { name }),

  // 画像設定
  updateImage: (imageUrl: string) =>
    apiClient.patch<UserResponse>("/api/v1/users/setup/image", { profile_image_url: imageUrl }),
};