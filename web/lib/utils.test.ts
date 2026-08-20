import { describe, expect, it } from "vitest";

import { cn } from "./utils";

describe("cn", () => {
  it("resolves conflicting Tailwind classes", () => {
    expect(cn("px-2 text-white", false && "hidden", "px-4")).toBe(
      "text-white px-4",
    );
  });
});
