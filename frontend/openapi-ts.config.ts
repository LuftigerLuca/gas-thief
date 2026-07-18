import { defineConfig } from "@hey-api/openapi-ts"

export default defineConfig({
  input: "../backend/app/docs/swagger.json",
  output: {
    path: "src/client",
    postProcess: ["prettier"],
  },
  plugins: [
    "@hey-api/client-fetch",
    "@tanstack/react-query",
  ],
})
