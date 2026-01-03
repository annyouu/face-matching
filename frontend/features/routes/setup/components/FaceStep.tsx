import { motion } from "framer-motion";
import { useRef, useState } from "react";

interface FaceStepProps {
  image: File | null;
  setImage: (file: File | null) => void;
  onComplete: () => void;
  onBack: () => void;
  isLoading: boolean;
}

export const FaceStep = ({ image, setImage, onComplete, onBack, isLoading }: FaceStepProps) => {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);

  // ファイルが選択された時の処理
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setImage(file);
      // ブラウザ上で表示するためのURLを生成
      const url = URL.createObjectURL(file);
      setPreviewUrl(url);
    }
  };

  return (
    <motion.div
      initial={{ opacity: 0, x: 20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: -20 }}
      className="space-y-6"
    >
      <div className="text-center space-y-2">
        <h2 className="text-2xl font-bold text-gray-900">Step 2</h2>
        <p className="text-gray-600">顔写真を登録しましょう！</p>
      </div>

      <div className="space-y-4">
        {/* アップロードエリア: クリックすると隠されたinputを叩く */}
        <div
          onClick={() => fileInputRef.current?.click()}
          className="w-full aspect-square border-2 border-dashed border-gray-300 rounded-2xl flex flex-col items-center justify-center bg-gray-50 cursor-pointer hover:bg-gray-100 transition-colors overflow-hidden"
        >
          {previewUrl ? (
            <img src={previewUrl} alt="Preview" className="w-full h-full object-cover" />
          ) : (
            <div className="text-center p-4">
              <div className="mb-2 text-[#7C74F7]">
                {/* シンプルなプラスアイコン */}
                <svg className="w-10 h-10 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="12 4v16m8-8H4" />
                </svg>
              </div>
              <p className="text-gray-500 font-medium">画像をアップロード</p>
              <p className="text-gray-400 text-xs mt-1">タップしてカメラロールから選択</p>
            </div>
          )}
          <input
            type="file"
            ref={fileInputRef}
            onChange={handleFileChange}
            className="hidden"
            accept="image/*"
          />
        </div>

        <div className="space-y-3">
          <button
            onClick={onComplete}
            disabled={!image || isLoading}
            className={`w-full py-3 rounded-xl font-bold transition-all shadow-md active:scale-[0.98] ${
              image && !isLoading
                ? "bg-[#7C74F7] text-white hover:brightness-110"
                : "bg-gray-200 text-gray-400 cursor-not-allowed"
            }`}
          >
            {isLoading ? "登録中..." : "登録を完了する"}
          </button>
          
          <button
            onClick={onBack}
            className="w-full py-2 text-sm text-gray-500 hover:text-gray-700 font-medium transition-colors"
          >
            戻る
          </button>
        </div>
      </div>
    </motion.div>
  );
};