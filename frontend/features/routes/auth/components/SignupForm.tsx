'use client';

import { useState } from "react";
import { Input } from "@/components/Input";
import Link from "next/link";

export const SignupForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const isEnabled = email !== "" && password !== "";

  return (
    <div className="w-full p-8 space-y-8 bg-white rounded-2xl shadow-sm border border-gray-100">
      <div className="text-center space-y-2">
        <h1 className="text-3xl font-bold tracking-tight text-gray-900">
          Destiny Face
        </h1>
        <h2 className="text-xl font-medium text-gray-600">新規登録</h2>
      </div>

      <form className="space-y-5" onSubmit={(e) => e.preventDefault()}>
        <div className="space-y-4">
          <Input
            label="Email"
            type="email"
            placeholder="example@mail.com"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <Input
            label="Password"
            type="password"
            placeholder="••••••••"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
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
          登録
        </button>
      </form>

      <div className="text-center">
        <Link
          href="/login"
          className="text-sm text-gray-500 hover:text-[#7C74F7] transition-colors"
        >
          すでにアカウントをお持ちの方はこちら
        </Link>
      </div>
    </div>
  );
};
