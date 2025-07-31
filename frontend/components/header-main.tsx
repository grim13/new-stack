import { ModeToggle } from "@/components/mode-tooggle";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import { SidebarTrigger } from "@/components/ui/sidebar";
import React from "react";
export function HeaderMain({
  breadcrumbs,
}: {
  breadcrumbs: {
    name: string;
    link: string;
  }[];
}) {
  return (
    <header className="bg-background sticky top-0 flex h-14 shrink-0 border-b px-4">
      <div className="flex flex-1 items-center gap-2">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb>
          <BreadcrumbList>
            {breadcrumbs.map((breadcrumb, index) => (
              <React.Fragment key={index}>
                <BreadcrumbItem
                  className={
                    index < breadcrumbs.length - 1 ? "hidden md:block" : ""
                  }
                >
                  {index < breadcrumbs.length - 1 ? (
                    <BreadcrumbLink href={breadcrumb.link}>
                      {breadcrumb.name}
                    </BreadcrumbLink>
                  ) : (
                    <BreadcrumbPage>{breadcrumb.name}</BreadcrumbPage>
                  )}
                </BreadcrumbItem>
                {index < breadcrumbs.length - 1 && (
                  <BreadcrumbSeparator className="hidden md:block" />
                )}
              </React.Fragment>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
      </div>
      <div className="ml-auto items-center flex gap-2">
        <ModeToggle />
      </div>
    </header>
  );
}
