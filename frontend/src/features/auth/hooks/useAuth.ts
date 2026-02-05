import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../api/Register";
import type { UserRegisterInput, UserStatus } from "../types";

export const useAuth = () => {
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();

    // バックエンドのStatusの応じて、次の画面を決める
    const navigateByStatus = (status: UserStatus) => {
        switch (status) {
            case 'PENDING_NAME':
                navigate('/setup/name');
                break;
            case 'PENDING_IMAGE':
                navigate('/setup/image');
                break;
            case 'ACTIVE':
                navigate('/home');
                break;
            default:
                navigate('/login');
        }
    };

    const register = async (data: UserRegisterInput) => {
        setIsLoading(true);
        try {
            const res = await registerUser(data);

            if (res.token) {
                document.cookie = `token=${res.token}; path=/; max-age=86400; SameSite=Lax`;
            }

            navigateByStatus(res.status);
        } catch (error) {
            console.error("登録失敗:", error);
            alert("新規登録に失敗しました。すでにアカウントがある可能性があります。");
        } finally {
            setIsLoading(false);
        }
    };

    return {
        register,
        isLoading
    };
};