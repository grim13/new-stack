import  { DefaultSession, DefaultUser } from "next-auth"
import { DefaultJWT } from "next-auth/jwt"

declare module "next-auth" {
  /**
   * Tipe `User` adalah objek yang dikembalikan dari callback `authorize`.
   * Di sini kita menambahkan field kustom kita.
   */
  interface User extends DefaultUser {
    id: uuid.UUID;
    username: string;
    roleId: number;
    roleName: string;
    accessToken: string;
  }

  /**
   * Tipe `Session` adalah objek yang dikembalikan oleh `useSession`, `getSession`, etc.
   * Kita menambahkan field kustom ke objek `user` di dalam session.
   */
  interface Session {
    accessToken?: string;
    user: {
      id: uuid.UUID;
      username: string;
      roleId: number;
      roleName: string;
    } & DefaultSession["user"]; // Menyimpan properti default seperti name, email, image
  }
}

declare module "next-auth/jwt" {
  /**
   * Tipe `JWT` adalah untuk token yang diteruskan ke callback `jwt` dan `session`.
   * Kita menambahkan field kustom kita ke dalamnya.
   */
  interface JWT extends DefaultJWT {
    id: uuid.UUID;
    username: string;
    roleId: number;
    roleName: string;
    accessToken: string;
  }
}