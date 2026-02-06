import { apiClient } from "@/api/apiClient";
import type { UserRegisterInput, UserResponse } from "../types";

/**
 * 新規登録リクエストを送る関数
 */

export const registerUser = (data: UserRegisterInput): Promise<UserResponse> => {
    return apiClient.post("auth/register", data);
};