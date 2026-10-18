# Gas Thief – Frontend

Web frontend for Gas Thief, built with [Next.js](https://nextjs.org/) 16, React 19, Tailwind CSS 4 and [shadcn/ui](https://ui.shadcn.com/).

## Getting Started

```bash
pnpm install
pnpm dev
```

The dev server runs on [http://localhost:3000](http://localhost:3000).

## Scripts

```bash
pnpm dev         # start dev server
pnpm build       # production build
pnpm start       # serve the production build
pnpm lint        # run ESLint
pnpm typecheck   # run tsc
pnpm format      # format with Prettier
```

## Adding Components

```bash
pnpm dlx shadcn@latest add button
```

Components are placed in `components/ui` and imported via the `@/` alias:

```tsx
import { Button } from "@/components/ui/button";
```
