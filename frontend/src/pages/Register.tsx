import { RegisterForm } from "@/features/auth/components/RegisterForm";

const RegisterPage = () => {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gra-50 py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full">
                <RegisterForm />
            </div>
        </div>
    );
};

export default RegisterPage;