import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./main.css";
import { QueryClientProvider } from "@tanstack/react-query";
import RootPage from "@/app/Root.page.tsx";
import { reactQueryClient } from "@/lib/reactQueryClient";
import { Toaster } from "sonner";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={reactQueryClient}>
      <RootPage />
      <Toaster />
    </QueryClientProvider>
  </StrictMode>,
);
