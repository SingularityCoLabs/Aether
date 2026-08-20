"use client";

import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { SystemService } from "@aether/api/system_pb";
import { Check, LoaderCircle, RefreshCw, Unplug } from "lucide-react";
import { useCallback, useEffect, useState } from "react";

import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";

type SystemState =
  | { status: "checking" }
  | {
      status: "online";
      version: string;
      environment: string;
      commit: string;
    }
  | { status: "offline"; message: string };

const baseURL =
  process.env.NEXT_PUBLIC_AETHER_API_URL ?? "http://localhost:8080";

export function ControlPlaneStatus() {
  const [state, setState] = useState<SystemState>({ status: "checking" });

  useEffect(() => {
    let active = true;
    void readSystemInfo().then((nextState) => {
      if (active) {
        setState(nextState);
      }
    });
    return () => {
      active = false;
    };
  }, []);

  const refresh = useCallback(() => {
    setState({ status: "checking" });
    void readSystemInfo().then(setState);
  }, []);

  return (
    <Card className="relative overflow-hidden p-5 sm:p-6">
      <div
        className="pointer-events-none absolute inset-x-12 -top-20 h-36 rounded-full bg-cyan-400/[0.07] blur-3xl"
        aria-hidden="true"
      />
      <div className="relative flex items-start justify-between gap-4">
        <div>
          <p className="text-xs font-medium tracking-[0.12em] text-zinc-500 uppercase">
            Control plane
          </p>
          <div className="mt-3 flex items-center gap-2.5">
            {state.status === "checking" && (
              <LoaderCircle className="size-4 animate-spin text-zinc-400" />
            )}
            {state.status === "online" && (
              <span className="grid size-5 place-items-center rounded-full bg-emerald-400/10">
                <Check className="size-3 text-emerald-300" strokeWidth={2.4} />
              </span>
            )}
            {state.status === "offline" && (
              <Unplug className="size-4 text-amber-300" />
            )}
            <h2 className="text-lg font-medium text-zinc-100">
              {state.status === "checking" && "Checking connection"}
              {state.status === "online" && "aetherd is online"}
              {state.status === "offline" && "aetherd is not connected"}
            </h2>
          </div>
          <p className="mt-2 max-w-md text-sm leading-6 text-zinc-500">
            {state.status === "checking" &&
              "Reading the typed SystemService at " + baseURL + "."}
            {state.status === "online" &&
              state.environment +
                " · " +
                state.version +
                " · commit " +
                state.commit}
            {state.status === "offline" && state.message}
          </p>
        </div>
        <Button
          variant="ghost"
          type="button"
          aria-label="Refresh control-plane connection"
          onClick={refresh}
          disabled={state.status === "checking"}
          className="shrink-0 px-2.5"
        >
          <RefreshCw
            className={cnRefresh(state.status === "checking")}
            aria-hidden="true"
          />
        </Button>
      </div>
    </Card>
  );
}

async function readSystemInfo(): Promise<SystemState> {
  try {
    const transport = createConnectTransport({ baseUrl: baseURL });
    const client = createClient(SystemService, transport);
    const response = await client.getSystemInfo({});
    return {
      status: "online",
      version: response.version,
      environment: response.environment,
      commit: response.commit,
    };
  } catch {
    return {
      status: "offline",
      message: "Start aetherd to connect this dashboard.",
    };
  }
}

function cnRefresh(spinning: boolean) {
  return "size-4 " + (spinning ? "animate-spin" : "");
}
