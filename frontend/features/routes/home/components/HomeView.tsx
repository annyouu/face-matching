'use client';

import React from 'react';
import Cookies from 'js-cookie';
import { useRouter } from 'next/navigation';

export const HomeView = () => {
  const router = useRouter();

  const handleLogout = () => {
    // Cookieを消してログイン画面へ戻る
    Cookies.remove('token');
    router.push('/login');
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col items-center justify-center p-4">
      <div className="bg-white p-8 rounded-2xl shadow-xl max-w-md w-full text-center space-y-6">
        <div className="inline-flex items-center justify-center w-16 h-16 bg-green-100 text-green-600 rounded-full">
          <svg className="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        
        <h1 className="text-3xl font-bold text-gray-900">ログイン成功！</h1>
        <p className="text-gray-600">
          無事に認証を通過し、セッションが確立されました。
        </p>

        <div className="pt-4">
          <button
            onClick={handleLogout}
            className="text-sm text-gray-500 hover:text-red-500 transition-colors"
          >
            ログアウトする
          </button>
        </div>
      </div>
    </div>
  );
};