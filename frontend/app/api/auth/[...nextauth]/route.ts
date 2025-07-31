import NextAuth, { AuthOptions } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import { jwtVerify, importSPKI } from 'jose';
import fs from 'fs';
import path from 'path';
interface UserJwtPayload {
  sub: string;
  iat: number;
  exp: number;
  userdata: {
    name: string;
    email: string;
    username: string;
    roleId: number;
    roleName: string;
  }
  // tambahkan klaim lain yang Anda harapkan ada di token
}
export const authOptions: AuthOptions = {
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        email: { label: "Email", type: "email" },
        password: {  label: "Password", type: "password" }
      },
      async authorize(credentials) {
        if (!credentials) {
          return null;
        }
        
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/auth/login`, {
          method: 'POST',
          body: JSON.stringify({ identifier: credentials.email, password: credentials.password }),
          headers: { "Content-Type": "application/json" }
        })
        const resData = await res.json()

        if (res.ok && resData) {
          try {
            // 1. Tentukan path ke file public.pem di root proyek
            const publicKeyPath = path.join(process.cwd(), 'public.pem');

            // 2. Baca isi file sebagai string
            const publicKey = fs.readFileSync(publicKeyPath, 'utf-8');

            // 3. Impor kunci dan verifikasi token
            const spki = await importSPKI(publicKey, 'RS256');
            const { payload } = await jwtVerify<UserJwtPayload>(resData.token, spki);
            return {
              id: payload.sub,
              name: payload.userdata.name,
              email: payload.userdata.email,
              username: payload.userdata.username,
              roleId: payload.userdata.roleId,
              roleName: payload.userdata.roleName,
              accessToken: resData.token
            };
          } catch (error) {
            console.error("JWT Verification failed:", error);
            return null;
          }
        }
        return null
      }
    })
  ],
  pages: {
    signIn: '/login',
  },
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.sub = String(user.id);
        token.username = user.username;
        token.roleId = user.roleId;
        token.roleName = user.roleName;
        token.accessToken = user.accessToken;
      }
      return token;
    },
    async session({ session, token }) {
      if (token && session.user) {
        session.user.id = token.sub as string;
        session.user.username = token.username as string;
        session.user.roleId = token.roleId as number;
        session.user.roleName = token.roleName as string;
        // Anda juga bisa meletakkan token akses langsung di level atas sesi
        session.accessToken = token.accessToken as string;
      }
      return session
    }
  }
}

const handler = NextAuth(authOptions)

export { handler as GET, handler as POST }
