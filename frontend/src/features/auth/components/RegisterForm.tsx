import { useState } from "react";
import { Link } from "react-router-dom";
import { Input } from "@/components/common/Input";
import { useAuth } from "../hooks/useAuth";

export const RegisterForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { register, isLoading } = useAuth();

  const isEnabled = email !== "" && password !== "" && !isLoading;

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!isEnabled) return;

    await register({ email, password });
  };

  return (
    <div className="w-full p-8 space-y-8 bg-white rounded-2xl shadow-xl border border-gray-100">
      <div className="text-center space-y-2">
        <h1 className="text-3xl font-bold tracking-tight text-gray-900">
          だぶりゼロ
        </h1>
        <h2 className="text-xl font-medium text-gray-600">新規登録</h2>
      </div>

      <form className="space-y-6" onSubmit={handleSubmit}>
        <div className="space-y-4">
          <Input
            label="Email"
            type="email"
            placeholder="example@gmail.com"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            disabled={isLoading}
          />

          <Input
            label="Password"
            type="password"
            placeholder="••••••••"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            disabled={isLoading}
          />
        </div>

        <button
          type="submit"
          disabled={!isEnabled}
          className={`
            w-full py-3 px-4 font-bold rounded-xl transition-all shadow-md
            active:scale-[0.98]
            ${
              isEnabled
                ? "bg-[#7C74F7] text-white hover:bg-[#6A62E5]"
                : "bg-gray-300 text-gray-500 cursor-not-allowed"
            }
          `}
        >
          {isLoading ? "登録中..." : "登録して次へ"}
        </button>
      </form>

      <div className="text-center">
        <Link
          to="/login"
          className="text-sm text-gray-500 hover:text-[#7C74F7] transition-colors"
        >
          すでにアカウントをお持ちの方はこちら
        </Link>
      </div>
    </div>
  );
};