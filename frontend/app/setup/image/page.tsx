import { FaceStep } from "@/features/routes/setup/components/FaceStep";

export default function NameSetupPage() {
    return (
        <main className="min-h-screen flex items-center justify-center bg-[#F8F9FF] p-4">
            <div className="max-w-md w-full bg-white rounded-3xl shadow-xl p-8">
                <FaceStep />
            </div>
        </main>
    );
}