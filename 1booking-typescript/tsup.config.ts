export default defineConfig({
  entry: ["src/index.ts"],
  format: ["cjs"],
  target: "node18",
  splitting: false,
  sourcemap: true,
  clean: true,
  dts: true,
  outDir: "dist",
  minify: true,
  rollup: {
    external: [],
    output: {
      entryFileNames: "index.js",
    },
  },
});
