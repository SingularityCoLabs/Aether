import {
  Activity,
  Boxes,
  Command,
  Database,
  Gauge,
  Hexagon,
  LayoutDashboard,
  ListChecks,
  Network,
  Server,
  ShieldCheck,
} from "lucide-react";
import type { ReactNode } from "react";

import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";

const navigation = [
  { label: "Overview", icon: LayoutDashboard, active: true },
  { label: "Projects", icon: Boxes },
  { label: "Nodes", icon: Server },
  { label: "Applications", icon: Gauge },
  { label: "Databases", icon: Database },
  { label: "Networking", icon: Network },
  { label: "Operations", icon: ListChecks },
  { label: "Audit", icon: ShieldCheck },
];

export function AppShell({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen lg:grid lg:grid-cols-[248px_1fr]">
      <aside className="hidden border-r border-white/[0.07] bg-black/20 lg:flex lg:flex-col">
        <div className="flex h-16 items-center gap-3 border-b border-white/[0.07] px-5">
          <div className="relative grid size-8 place-items-center">
            <Hexagon className="size-8 text-cyan-300" strokeWidth={1.25} />
            <span className="absolute size-1.5 rounded-full bg-cyan-200 shadow-[0_0_12px_#67e8f9]" />
          </div>
          <div>
            <p className="text-sm font-semibold tracking-[0.18em] text-zinc-100">
              AETHER
            </p>
            <p className="text-[10px] tracking-wide text-zinc-500">
              CONTROL PLANE
            </p>
          </div>
        </div>

        <div className="px-3 py-5">
          <p className="mb-2 px-3 text-[10px] font-medium tracking-[0.16em] text-zinc-600 uppercase">
            Workspace
          </p>
          <nav aria-label="Primary navigation" className="space-y-1">
            {navigation.map(({ label, icon: Icon, active }) => (
              <button
                key={label}
                type="button"
                aria-current={active ? "page" : undefined}
                disabled={!active}
                className={cn(
                  "flex h-9 w-full items-center gap-3 rounded-lg px-3 text-sm transition-colors",
                  active
                    ? "bg-white/[0.07] text-white"
                    : "cursor-not-allowed text-zinc-600",
                )}
              >
                <Icon className="size-4" strokeWidth={1.7} />
                {label}
                {!active && (
                  <span className="ml-auto text-[9px] tracking-wider text-zinc-700 uppercase">
                    Soon
                  </span>
                )}
              </button>
            ))}
          </nav>
        </div>

        <div className="mt-auto border-t border-white/[0.07] p-4">
          <div className="rounded-xl border border-white/[0.07] bg-white/[0.025] p-3.5">
            <div className="mb-2 flex items-center gap-2">
              <Command className="size-3.5 text-zinc-500" />
              <span className="text-xs font-medium text-zinc-300">Phase 0</span>
              <Badge className="ml-auto border-cyan-400/15 bg-cyan-400/[0.06] text-cyan-200">
                Foundation
              </Badge>
            </div>
            <p className="text-[11px] leading-relaxed text-zinc-600">
              No node, Docker, or AI execution is enabled yet.
            </p>
          </div>
        </div>
      </aside>

      <div className="min-w-0">
        <header className="sticky top-0 z-20 flex h-16 items-center justify-between border-b border-white/[0.07] bg-[#090a0b]/85 px-5 backdrop-blur-xl sm:px-8">
          <div className="flex items-center gap-3 lg:hidden">
            <Hexagon className="size-6 text-cyan-300" strokeWidth={1.25} />
            <span className="text-sm font-semibold tracking-[0.16em]">
              AETHER
            </span>
          </div>
          <div className="hidden items-center gap-2 text-xs text-zinc-500 lg:flex">
            <Activity className="size-3.5" />
            Local development workspace
          </div>
          <Badge className="border-amber-400/15 bg-amber-400/[0.06] text-amber-200">
            Pre-release
          </Badge>
        </header>

        <main className="relative isolate">
          <div className="aether-grid pointer-events-none absolute inset-x-0 top-0 -z-10 h-[540px]" />
          {children}
        </main>
      </div>
    </div>
  );
}
