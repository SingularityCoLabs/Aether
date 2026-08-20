import {
  ArrowRight,
  Braces,
  Boxes,
  Database,
  GitPullRequestArrow,
  RadioTower,
  ShieldCheck,
} from "lucide-react";

import { AppShell } from "@/components/layout/app-shell";
import { ControlPlaneStatus } from "@/components/system/control-plane-status";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";

const foundations = [
  {
    title: "Shared contract",
    detail: "Protobuf · Buf · ConnectRPC",
    icon: Braces,
  },
  {
    title: "Control-plane state",
    detail: "PostgreSQL · pgx · Goose",
    icon: Database,
  },
  {
    title: "Provider boundary",
    detail: "Plan · Apply · Observe · Delete",
    icon: Boxes,
  },
  {
    title: "Execution safety",
    detail: "Risk · Approval · Operation · Audit",
    icon: ShieldCheck,
  },
];

const sequence = [
  { label: "Repository foundation", state: "complete" },
  { label: "Identity and projects", state: "next" },
  { label: "Node enrollment and heartbeat", state: "later" },
  { label: "Resources and operations", state: "later" },
  { label: "Docker application provider", state: "later" },
];

export default function Home() {
  return (
    <AppShell>
      <div className="mx-auto max-w-7xl px-5 py-10 sm:px-8 sm:py-14">
        <section className="max-w-3xl">
          <Badge className="border-cyan-400/15 bg-cyan-400/[0.06] text-cyan-200">
            <RadioTower className="size-3" />
            Architecture online
          </Badge>
          <h1 className="mt-5 text-3xl font-medium tracking-[-0.035em] text-zinc-50 sm:text-5xl">
            One control plane.
            <span className="block text-zinc-500">
              Every interface aligned.
            </span>
          </h1>
          <p className="mt-5 max-w-2xl text-sm leading-7 text-zinc-500 sm:text-base">
            Aether begins with a typed resource and operation boundary. The
            dashboard, CLI, and future agent all enter through the same API;
            infrastructure providers remain behind it.
          </p>
        </section>

        <div className="mt-10 grid gap-5 xl:grid-cols-[1.15fr_0.85fr]">
          <ControlPlaneStatus />

          <Card className="p-5 sm:p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-xs font-medium tracking-[0.12em] text-zinc-500 uppercase">
                  Release target
                </p>
                <h2 className="mt-2 text-lg font-medium text-zinc-100">
                  Aether v0.1 · Node
                </h2>
              </div>
              <GitPullRequestArrow className="size-5 text-zinc-600" />
            </div>
            <p className="mt-4 text-sm leading-6 text-zinc-500">
              Connect one Linux server and deploy, observe, restart, and delete
              one Docker application with audit history.
            </p>
          </Card>
        </div>

        <section className="mt-5 grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
          {foundations.map(({ title, detail, icon: Icon }) => (
            <Card key={title} className="p-5">
              <Icon className="size-4 text-cyan-200/75" strokeWidth={1.6} />
              <h3 className="mt-6 text-sm font-medium text-zinc-200">
                {title}
              </h3>
              <p className="mt-1.5 text-xs leading-5 text-zinc-600">{detail}</p>
            </Card>
          ))}
        </section>

        <section className="mt-5 grid gap-5 xl:grid-cols-[0.92fr_1.08fr]">
          <Card className="p-5 sm:p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-xs font-medium tracking-[0.12em] text-zinc-500 uppercase">
                  Capability order
                </p>
                <h2 className="mt-2 text-lg font-medium text-zinc-100">
                  Build from the boundary inward
                </h2>
              </div>
              <ArrowRight className="size-4 text-zinc-700" />
            </div>
            <ol className="mt-6 space-y-1">
              {sequence.map(({ label, state }, index) => (
                <li
                  key={label}
                  className="flex items-center gap-3 rounded-xl px-3 py-2.5"
                >
                  <span
                    className={
                      state === "complete"
                        ? "grid size-6 place-items-center rounded-full bg-emerald-400/10 text-[10px] font-semibold text-emerald-300"
                        : state === "next"
                          ? "grid size-6 place-items-center rounded-full border border-cyan-300/25 bg-cyan-300/[0.05] text-[10px] font-semibold text-cyan-200"
                          : "grid size-6 place-items-center rounded-full border border-white/[0.07] text-[10px] text-zinc-700"
                    }
                  >
                    {String(index + 1).padStart(2, "0")}
                  </span>
                  <span
                    className={
                      state === "later"
                        ? "text-sm text-zinc-600"
                        : "text-sm text-zinc-300"
                    }
                  >
                    {label}
                  </span>
                  {state === "next" && (
                    <Badge className="ml-auto border-cyan-400/15 bg-cyan-400/[0.05] text-cyan-300">
                      Next
                    </Badge>
                  )}
                </li>
              ))}
            </ol>
          </Card>

          <Card className="overflow-hidden">
            <div className="border-b border-white/[0.07] px-5 py-4 sm:px-6">
              <p className="text-xs font-medium tracking-[0.12em] text-zinc-500 uppercase">
                Non-negotiable path
              </p>
            </div>
            <div className="grid min-h-64 place-items-center p-7">
              <div className="flex w-full max-w-lg flex-col items-center">
                <div className="grid w-full grid-cols-[1fr_auto_1fr] items-center gap-3">
                  <FlowNode label="Web · CLI · AI" muted />
                  <ArrowRight className="size-4 text-zinc-700" />
                  <FlowNode label="Aether API" accent />
                </div>
                <div className="my-3 h-8 w-px bg-gradient-to-b from-cyan-300/40 to-white/10" />
                <FlowNode label="Resources · Plans · Operations" />
                <div className="my-3 h-8 w-px bg-white/10" />
                <FlowNode label="Typed providers" muted />
              </div>
            </div>
          </Card>
        </section>

        <footer className="mt-10 flex flex-col gap-2 border-t border-white/[0.06] pt-5 text-[11px] text-zinc-700 sm:flex-row sm:items-center sm:justify-between">
          <span>Aether Phase 0 · no infrastructure mutations enabled</span>
          <span>Resource API first · agent tools last</span>
        </footer>
      </div>
    </AppShell>
  );
}

function FlowNode({
  label,
  accent = false,
  muted = false,
}: {
  label: string;
  accent?: boolean;
  muted?: boolean;
}) {
  return (
    <div
      className={[
        "w-full rounded-xl border px-4 py-3 text-center text-xs font-medium",
        accent
          ? "border-cyan-300/20 bg-cyan-300/[0.06] text-cyan-100 shadow-[0_0_30px_rgba(34,211,238,0.06)]"
          : muted
            ? "border-white/[0.06] bg-black/10 text-zinc-600"
            : "border-white/[0.09] bg-white/[0.03] text-zinc-300",
      ].join(" ")}
    >
      {label}
    </div>
  );
}
