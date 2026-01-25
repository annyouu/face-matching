'use client';

import { useState } from "react";
import { Input } from "@/components/Input";
import { useAuth } from "../hooks";

export const LoginForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { login, isLoading } = useAuth();

  const isEnabled = email !== "" && password !== "" && !isLoading;
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await login({
      email,
      password,
    });
  };

  return (
    <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-xl shadow-lg">
      <div className="text-center">
        <h2 className="text-2xl font-bold text-gray-900">ログイン</h2>
        <p className="mt-2 text-sm text-gray-600">
          おかえりなさい！情報を入力してください。
        </p>
      </div>

      <form className="space-y-4" onSubmit={handleSubmit}>
        <Input
          label="メールアドレス"
          type="email"
          placeholder="example@mail.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          disabled={isLoading}
        />

        <Input
          label="パスワード"
          type="password"
          placeholder="••••••••"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          disabled={isLoading}
        />

        <button
          type="submit"
          disabled={!isEnabled}
          className={`
            w-full py-2.5 px-4 font-semibold rounded-lg transition-all duration-200
            ${isEnabled
              ? "bg-[#7C74F7] text-white hover:brightness-110"
              : "bg-gray-300 text-gray-500 cursor-not-allowed"}
          `}
        >
          {isLoading ? "ログイン中..." : "ログイン"}
        </button>
      </form>
    </div>
  );
};
