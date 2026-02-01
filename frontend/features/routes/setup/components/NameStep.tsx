"use client";

import { useState } from "react";
import { Input } from "@/components/Input";
import { motion } from "framer-motion";
import { useSetup } from "../hooks";

export const NameStep = () => {
  const [name, setName] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { submitName } = useSetup();

  const handleNext = async () => {
    if (!name || isSubmitting) return;
    setIsSubmitting(true);
    try {
      await submitName(name);
    } catch (error) {
      console.error("名前送信エラー:", error);
      alert("名前の保存に失敗しました。");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    // 1. 画面全体の背景色と、スマホ時に端に張り付かないための p-4 を追加
    <main className="min-h-screen flex items-center justify-center bg-[#F8F9FF] p-4">
      {/* 2. カード部分: FaceStepと同じデザインにする */}
      <div className="max-w-md w-full bg-white rounded-4xl shadow-xl p-8">
        <motion.div
          initial={{ opacity: 0, x: 20 }}
          animate={{ opacity: 1, x: 0 }}
          exit={{ opacity: 0, x: -20 }}
          className="space-y-8"
        >
          {/* ヘッダー部分 */}
          <div className="text-center space-y-2">
            <h2 className="text-2xl font-bold text-gray-900">Step 1</h2>
            <p className="text-gray-600">あなたのニックネームを教えてください</p>
          </div>

          {/* 入力フォーム部分 */}
          <div className="space-y-6">
            <Input
              label="名前"
              placeholder="例: たろう"
              value={name}
              onChange={(e) => setName(e.target.value)}
              disabled={isSubmitting}
            />

            {/* 送信ボタン: FaceStepと丸みや高さを統一 */}
            <button
              onClick={handleNext}
              disabled={!name || isSubmitting}
              className={`w-full py-4 rounded-2xl font-bold text-lg transition-all shadow-md active:scale-[0.98] ${
                name && !isSubmitting
                  ? "bg-[#7C74F7] text-white hover:brightness-110"
                  : "bg-gray-200 text-gray-400 cursor-not-allowed"
              }`}
            >
              {isSubmitting ? (
                <div className="flex items-center justify-center gap-2">
                  <div className="h-5 w-5 border-2 border-white border-t-transparent rounded-full animate-spin" />
                  <span>保存中...</span>
                </div>
              ) : (
                "次へ進む"
              )}
            </button>
          </div>
        </motion.div>
      </div>
    </main>
  );
};