import { LoginForm } from "@/components/login-form"
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Login Page",
  description: "Login to your account",
};

export default function LoginPage() {
  return (
    <div className="bg-muted flex min-h-svh flex-col items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm md:max-w-3xl">
        <LoginForm />
      </div>
    </div>
  )
}
