import { Input } from "@/components/Input";
import { motion } from "framer-motion";

interface NameStepProps {
  name: string;
  setName: (name: string) => void;
  onNext: () => void;
}

export const NameStep = ({ name, setName, onNext }: NameStepProps) => {
  return (
    <motion.div
      initial={{ opacity: 0, x: 20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: -20 }}
      className="space-y-6"
    >
      <div className="text-center space-y-2">
        <h2 className="text-2xl font-bold text-gray-900">Step 1</h2>
        <p className="text-gray-600">あなたのニックネームを教えてください</p>
      </div>

      <div className="space-y-4">
        <Input
          label="名前"
          placeholder="例: たろう"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        
        <button
          onClick={onNext}
          disabled={!name}
          className={`w-full py-3 rounded-xl font-bold transition-all shadow-md active:scale-[0.98] ${
            name
              ? "bg-[#7C74F7] text-white hover:brightness-110"
              : "bg-gray-200 text-gray-400 cursor-not-allowed"
          }`}
        >
          次へ進む
        </button>
      </div>
    </motion.div>
  );
};