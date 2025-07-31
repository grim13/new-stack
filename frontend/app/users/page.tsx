import { AppSidebar } from "@/components/app-sidebar"
import { HeaderMain } from "@/components/header-main"
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar"

export default function Page() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <HeaderMain
          breadcrumbs={[
            { name: "Dashboard", link: "/" },
            { name: "Manage Users", link: "#" },
          ]}
        />
        <div className="flex flex-1 flex-col gap-4 p-4">
          
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
