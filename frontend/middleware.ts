export { default } from "next-auth/middleware"

export const config = {
  matcher: ["/","/users", "/profile", "/roles"],
}
