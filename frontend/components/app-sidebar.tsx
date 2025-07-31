"use client"

import * as React from "react"
import {
  LucideIcon,
  Users2Icon
} from "lucide-react"

import { NavMain } from "@/components/nav-main"
import { NavUser } from "@/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar"

import { useSession } from "next-auth/react"
import { RoleSwitcher } from "./role-switcher"


// This is sample data.
const data = {
  app: {
    name: process.env.APP_NAME || "New Stack",
  },
  user: {} as {
    id: string,
    roleId: number,
    roleName: string,
    username: string,
    name: string,
    email: string,
    avatar: string,
  },
  navMain: [] as {
    title: string,
    url: string,
    icon: LucideIcon | undefined,
    items: {
      title: string,
      url: string
    }[]
  }[],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { data: session } = useSession()
  data.navMain = []
  if (session?.user.roleId === 1) {
    data.navMain.push({
      title: "Authentication",
      url: "#",
      icon: Users2Icon,
      items: [
        {
          title: "Manage Users",
          url: "/users",
        },
        {
          title: "Manage Roles",
          url: "/roles",
        },
      ],
    })
  }
  data.user = {
    id: session?.user?.id || "",
    roleId: session?.user?.roleId || 0,
    roleName: session?.user?.roleName || "Guest",
    username: session?.user?.username || "Guest",
    name: session?.user?.name || "Guest",
    email: session?.user?.email || "guest@example.com",
    avatar: session?.user?.image || "/avatars/shadcn.jpg",
  }
  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <RoleSwitcher user={data.user} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
