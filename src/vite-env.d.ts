/// <reference types="vite/client" />
interface ImportMetaEnv {
    readonly VITE_GCLIENTID: string
  }
  
  interface ImportMeta {
    readonly env: ImportMetaEnv
  }