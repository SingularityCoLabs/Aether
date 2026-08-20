import type { ComponentProps } from "react";

import { cn } from "@/lib/utils";

export function Card({ className, ...props }: ComponentProps<"section">) {
  return (
    <section
      className={cn(
        "rounded-2xl border border-white/[0.08] bg-[#111214]/85 shadow-[0_24px_80px_rgba(0,0,0,0.2)] backdrop-blur",
        className,
      )}
      {...props}
    />
  );
}
