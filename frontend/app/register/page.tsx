import { RegisterForm } from "@/features/routes/auth/components/RegisterForm";

export default function RegisterPage() {
    return (
        <main className="min-h-screen flex items-center justify-center bg-slate-50 relative overflow-hidden">
            <div className="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] rounded-full bg-blue-100/50 blur-3xl" />
            <div className="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] rounded-full bg-purple-100/50 blur-3xl" />

            <section className="relative z-10 w-full max-w-md px-4">
                <RegisterForm />
            </section>
        </main>
    );
}