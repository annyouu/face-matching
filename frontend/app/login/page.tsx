import { LoginForm } from "@/features/routes/auth/components/LoginForm";

export default function LoginPage() {
  return (
    <main className="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div className="absolute inset-0 bg-linear-to-br from-blue-50 to-indigo-100 -z-10" />
      
      <LoginForm />
    </main>
  );
}