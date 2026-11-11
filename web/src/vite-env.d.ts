interface ImportMetaEnv {
  readonly VITE_API_URL: string
  readonly VITE_AS_OF: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}