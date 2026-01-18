import { apiClient } from "@/lib/apiClient";

// UserLoginInput
export interface SignupRequest {
    email: string;
    password: string;
}

// AuthTokenOutput
export interface SignupResponse {
    message: string;
    token: string;
}

export const signupRequest = async (data: SignupRequest) => {
    return await apiClient.post("/api/v1/signup", data);
}