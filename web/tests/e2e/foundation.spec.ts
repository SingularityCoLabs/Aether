import { expect, test } from "@playwright/test";

test("renders the honest Phase 0 foundation", async ({ page }) => {
  await page.goto("/");

  await expect(
    page.getByRole("heading", {
      name: "One control plane. Every interface aligned.",
    }),
  ).toBeVisible();
  await expect(page.getByText("Aether v0.1 · Node")).toBeVisible();
  await expect(
    page.getByText("no infrastructure mutations enabled", { exact: false }),
  ).toBeVisible();
});
